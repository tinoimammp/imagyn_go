package postgres

import (
	"gitlab.com/imagyn_go/internal/infrastructure/config"
	"gorm.io/gorm"
)

type DatabaseProvider struct {
	db *gorm.DB
}

func NewDatabaseProvider(cfg *config.DatabaseConfig) (*DatabaseProvider, error) {
	database, err := config.NewDatabase(cfg)
	if err != nil {
		return nil, err
	}
	return &DatabaseProvider{db: database.GetDB()}, nil
}

func (p *DatabaseProvider) GetDB() *gorm.DB {
	return p.db
}
