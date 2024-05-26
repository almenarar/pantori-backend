package infra

import (
	core "pantori/internal/domains/goods/core"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

func NewMySQL(conn string) *sql {
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect to database")
	}

	db.AutoMigrate(&core.Good{})

	return &sql{
		db: db,
	}
}

func (s *sql) CreateItem(good core.Good) error {
	good.ID = uuid.New().String()
	result := s.db.Create(good)
	if result.Error != nil {
		return errors.Wrap(result.Error, "")
	}
	return nil
}

func (s *sql) GetItemByID(good core.Good) (core.Good, error) {
	var item core.Good
	result := s.db.First(&item, good.ID)
	if result.Error != nil {
		return core.Good{}, errors.Wrap(result.Error, "")
	}
	return item, nil
}

func (s *sql) GetAllItems(string) ([]core.Good, error) {
	var goods []core.Good
	result := s.db.Find(&goods)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "")
	}
	return goods, nil
}

func (s *sql) EditItem(good core.Good) error {
	return errors.New("not implemented")
}

func (s *sql) DeleteItem(good core.Good) error {
	result := s.db.Delete(good)
	if result.Error != nil {
		return errors.Wrap(result.Error, "")
	}
	return nil
}
