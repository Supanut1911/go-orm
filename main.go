package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gender struct {
	ID uint
	Name string
}

func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}



func main() {

	host := goDotEnvVariable("HOST")
	port := goDotEnvVariable("PORT")
	user := goDotEnvVariable("USERNAME")
	password := goDotEnvVariable("PASSWORD")
	dbname := goDotEnvVariable("DBNAME")


	dsn := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	dial := postgres.Open(dsn)
	db, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}

	db.Migrator().CreateTable(Gender {})
}