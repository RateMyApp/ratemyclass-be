package dao

import (
	"context"

	"github.com/ratemyapp/config"
	"github.com/ratemyapp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresClient struct {
	Db               *gorm.DB
	connectionString string
}

func (pc *PostgresClient) Init() error {
	db, err := gorm.Open(postgres.Open(pc.connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return err
	}

	pc.Db = db

	// Initial setup
	db.Exec(`
        CREATE TYPE status AS ENUM ('PENDING', 'APPROVED', 'DECLINED');
    `)

	// migrations
	for _, model := range *models.GetModels() {

		err := pc.Db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pc *PostgresClient) Close(ctx context.Context) error {
	postgresDb, err := pc.Db.DB()
	if err != nil {
		return err
	}

	postgresDb.Close()
	return nil
}

func NewPostgresClient(appConfig config.AppConfig) (DbClient, *PostgresClient) {
	client := &PostgresClient{connectionString: appConfig.POSTGRES_URI}
	return client, client
}
