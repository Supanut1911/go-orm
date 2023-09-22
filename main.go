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

type Test struct {
	gorm.Model
	Name string `gorm:"colume:myname;unique;default:YOYO;not null"`
	Desc string
}

type Gender struct {
	ID uint
	Name string `gorm:"unique"`	
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


	dsn := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, dbname)
	dial := postgres.Open(dsn)

	var err error
	db, err := gorm.Open(dial, &gorm.Config{Logger: &SqlLogger{}, DryRun: false})
	if err != nil {
		panic(err)
	}

	_ = db

	// db.AutoMigrate(Gender{}, Test{})
}