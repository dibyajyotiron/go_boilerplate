package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB interface {
	Connect(dsn string) (*gorm.DB, error)
}

type GormDB struct{}

func (g *GormDB) Connect(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
