package core

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Service struct {
	numWorkers int
	svcGoods   GoodsPort
	svcUsers   UsersPort
	email      EmailPort
}

func NewService(
	goods GoodsPort,
	users UsersPort,
	email EmailPort,
	numWorkers int,
) *Service {
	return &Service{
		svcGoods:   goods,
		svcUsers:   users,
		email:      email,
		numWorkers: numWorkers,
	}
}

func (svc *Service) NotifyExpiredGoods(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			svc.runJobs()
			time.Sleep(24 * time.Hour)
		}
	}
}

func (svc *Service) runJobs() {
	users, err := svc.svcUsers.ListAllUsersByWorkspace()
	if err != nil {
		log.Println("Error running domain2 job:", err)
	}

	jobs := make(chan []User, len(users))

	var wg sync.WaitGroup

	for i := 0; i < svc.numWorkers; i++ {
		wg.Add(1)
		go svc.worker(jobs, &wg)
	}

	for k := range users {
		jobs <- users[k]
	}
	close(jobs)

	wg.Wait()

	log.Println("Scheduled job completed.")
}

func (svc *Service) worker(jobs <-chan []User, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		for _, user := range job {
			svc.processEntry(user)
		}
	}
}

func (svc *Service) processEntry(user User) {
	dateFormat := "02/01/2006"
	currentDate := time.Now()

	var expired []Good
	var expiresToday []Good
	var expiresSoon []Good

	goods, err := svc.svcGoods.GetGoodsFromWorkspace(user.Workspace)
	if err != nil {
		log.Println("Error running domain1 job:", err)
	}

	for _, good := range goods {
		parsedDate, err := time.Parse(dateFormat, good.Expire)
		if err != nil {
			fmt.Printf("oops at %s/n", user.Name)
		}

		diff := parsedDate.Sub(currentDate).Hours() / 24
		switch {
		case diff == 0:
			expiresToday = append(expiresToday, good)
		case 0 < diff && diff < 3:
			expiresSoon = append(expiresSoon, good)
		case diff < 0:
			expired = append(expired, good)
		}
	}

	err = svc.email.SendEmail(user, expiresToday, expiresSoon, expired)
	if err != nil {
		panic(err)
	}
}
