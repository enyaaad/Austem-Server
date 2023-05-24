package db

import (
	"fmt"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file found")
	}
}

func StartDB() {
	host, exists := os.LookupEnv("POSTGRES_HOST")
	user, exists := os.LookupEnv("POSTGRES_USER")
	pass, exists := os.LookupEnv("POSTGRES_PASSWORD")
	name, exists := os.LookupEnv("POSTGRES_DB")
	port, exists := os.LookupEnv("POSTGRES_PORT")

	if !exists {
		fmt.Println("err")
	}

	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user,
		pass, name, port,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed connect to database")
	}
	fmt.Println("success while connecting to db")
}
