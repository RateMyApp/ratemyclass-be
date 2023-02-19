package dao

import (
	"context"

	"github.com/ratemyapp/config"
	"github.com/ratemyapp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	db               *gorm.DB
	connectionString string
}

func (pc *PostgresClient) Init() error {
	db, err := gorm.Open(postgres.Open(pc.connectionString), &gorm.Config{})

	if err != nil {
		return err
	}

	pc.db = db

	// migrations

	tables := []interface{}{
		&models.School{},
		&models.Program{},
	}

	for _, table := range tables {

		err := pc.db.AutoMigrate(table)
		if err != nil {
			return err
		}
	}

	return nil

}

func (pc *PostgresClient) Close(ctx context.Context) error {
	postgresDb, err := pc.db.DB()

	if err != nil {
		return err
	}

	postgresDb.Close()
	return nil
}

func NewPostgresClient(appConfig config.AppConfig) DbClient {

	return &PostgresClient{connectionString: appConfig.POSTGRES_URI}
}
