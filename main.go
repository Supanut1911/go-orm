package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Gender struct {
	ID uint
	Code uint `gorm:"primaryKey"`
	Name string `gorm:"colume:myname;unique;default:YOYO;not null"`
	Desc string
}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error ) {
	sql, _ := fc()
	fmt.Printf("%v \n======================================\n", sql)
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
	db, err := gorm.Open(dial, &gorm.Config{Logger: &SqlLogger{}, DryRun: true})
	if err != nil {
		panic(err)
	}

	// errCreatetable := db.Migrator().CreateTable(Gender {})
	// if errCreatetable != nil {
	// 	panic("errCreatetable")
	// }

	// db.AutoMigrate(Gender{})
	db.Migrator().CreateTable(Gender{})
}