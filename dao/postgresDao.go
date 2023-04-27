package dao

import (
	"context"

	"github.com/ratemyapp/config"
	"github.com/ratemyapp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	Db               *gorm.DB
	connectionString string
}

func (pc *PostgresClient) Init() error {
	db, err := gorm.Open(postgres.Open(pc.connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	pc.Db = db

	// migrations

	tables := []interface{}{
		&models.School{},
		&models.Program{},
		&models.User{},
		&models.Course{},
		&models.CourseRating{},
		&models.Professor{},
	}

	for _, table := range tables {

		err := pc.Db.AutoMigrate(table)
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
