package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := "host=3yt.h.filess.io user=pointingpoker123_winterpale password=b008c4730d68fa04059a0d536849cf9e1cd38de1 dbname=pointingpoker123_winterpale port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}
}
