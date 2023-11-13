package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-2/config"
	"github.com/hacktiv8-ks07-g04/final-project-2/domain/entity"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	Connect()
	Migration()
}

func Connect() {
	dbConfig := config.Get().Database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database, ", err)
	}
}

func Migration() {
	db.AutoMigrate(
		&entity.User{},
		&entity.Photo{},
		&entity.Comment{},
		&entity.SocialMedia{},
	)
}

func GetInstance() *gorm.DB {
	if db == nil {
		log.Fatal("Database instance is not initialized")
	}

	return db
}
