package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var exist bool
	var user string
	var pwd string
	var host string
	user, exist = os.LookupEnv("DB_USER")
	pwd, exist = os.LookupEnv("DB_PASSWORD")
	host, exist = os.LookupEnv("DB_HOST")
	if !exist {
		panic("Missing environment settings")
	}
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/LetsGo?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connected to database")
}
