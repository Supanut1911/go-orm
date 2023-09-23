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

var db *gorm.DB

func main() {
	host := goDotEnvVariable("HOST")
	port := goDotEnvVariable("PORT")
	user := goDotEnvVariable("USERNAME")
	password := goDotEnvVariable("PASSWORD")
	dbname := goDotEnvVariable("DBNAME") 

	dsn := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, dbname)
	dial := postgres.Open(dsn)

	var err error
	db, err = gorm.Open(dial, &gorm.Config{Logger: &SqlLogger{}, DryRun: false})
	if err != nil {
		panic(err)
	}

	_ = db

	// db.AutoMigrate(Gender{}, Test{})

	// CreateGender("LGBT+")
	
	// GetGenders()
	// GetGenderById(50)
	// GetGenderByName("Male")
	// UpdateGender(1, "malee")
	// UpdateGender2(1, "MMALE")
	DeleteGender(5)
}

func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenders() {
	genders := []Gender{}
	tx := db.Order("id desc").Find(&genders)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGenderById(id uint) {
	gender := Gender{}
	tx := db.Find(&gender, id)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenderByName(name string) {
	gender := Gender{}
	// tx := db.Find(&gender, "name=?", name)
	tx := db.Where("name=?", name).Find(&gender)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return
	}
	fmt.Println(gender)
}

//ท่า1 -> query by id มาก่อน, แก้ไขบาง field จากนั้น save
func UpdateGender(id uint, name string) {
	gender := Gender{}

	//get by id
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return
	}

	//update 
	gender.Name = name

	tx = db.Save(gender)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return
	}
	fmt.Print("update success", gender)
}

//ท่า2 -> ท่านั้นไม่ได้ไป query จาก db, ตอนทที่ instance ให้ใส่ค่าที่จะ update เข้าไปเลย
func UpdateGender2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id = ?", id).Updates(gender)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return
	}

	GetGenderById(id)

}

func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Print(tx.Error)
		return
	}
	fmt.Print("Deleted ")
	GetGenderById(id)
	
}
