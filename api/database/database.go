package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "host=localhost user=postgres password=123456 dbname=goauth port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the databas! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database !!")
	db.Logger = logger.Default.LogMode(logger.Info)

	// db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}

}
