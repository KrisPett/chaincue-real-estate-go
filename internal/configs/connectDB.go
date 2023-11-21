package configs

import (
	"chaincue-real-estate-go/internal/domains"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=admin password=admin dbname=chaincue-real-estate-db port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&domains.Broker{},
		&domains.Country{},
	)

	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Database migration completed successfully.")
}

func GetDB() *gorm.DB {
	return db
}
