package core

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
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

func (svc *Service) NotifyExpiredGoods() {
	users, err := svc.svcUsers.ListAllUsersByWorkspace()
	if err != nil {
		log.Error().Stack().Err(&ErrListUsers{Err: err}).Msg("")
		return
	}

	svc.startWorkers(users)
}

func (svc *Service) startWorkers(users map[string][]User) {
	jobs := make(chan []User, len(users))

	var wg sync.WaitGroup

	for i := 0; i < svc.numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for users := range jobs {
				svc.worker(users)
			}
		}()
	}

	for workspace := range users {
		jobs <- users[workspace]
	}
	close(jobs)

	wg.Wait()

}

func (svc *Service) worker(users []User) {
	goods, err := svc.svcGoods.GetGoodsFromWorkspace(users[0].Workspace)
	if err != nil {
		log.Error().Stack().Err(&ErrGetGoods{Err: err}).Msg("")
		return
	}

	report := svc.CreateReport(goods)
	if len(report.Expired)+len(report.ExpiresSoon)+len(report.ExpiresToday) > 0 {
		for _, user := range users {
			err = svc.email.SendEmail(user, report)
			if err != nil {
				log.Error().Stack().Err(&ErrSendEmail{Err: err}).Msg("")
				return
			}
		}
	}
}

func (svc *Service) CreateReport(goods []Good) Report {
	dateFormat := "02/01/2006"
	currentDate := time.Now()

	var report Report

	for _, good := range goods {
		if good.Quantity == "Empty" {
			continue
		}

		parsedDate, err := time.Parse(dateFormat, good.Expire)
		if err != nil {
			log.Error().Stack().Err(&ErrTimeParse{Date: good.Expire}).Msg("")
			return report
		}

		diff := parsedDate.Sub(currentDate).Hours() / 24
		switch {
		case currentDate.Format("02/01/2006") == good.Expire:
			report.ExpiresToday = append(report.ExpiresToday, good)
		case 0 < diff && diff < 4:
			report.ExpiresSoon = append(report.ExpiresSoon, good)
		case diff < 0:
			report.Expired = append(report.Expired, good)
		}
	}

	return report
}
