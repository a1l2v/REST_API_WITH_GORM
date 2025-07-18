package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB()(*gorm.DB){
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal("Error loading")
		return nil
	}
	fmt.Println("Connected to DB")
	return DB
}

