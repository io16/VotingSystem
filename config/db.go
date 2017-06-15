package config

import (
	"github.com/jinzhu/gorm"

	"log"
	"fmt"
)

var DB *gorm.DB

func InitDB() {
	var err error;
	println("DB Init")
	DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=AIPPZ sslmode=disable password=299792458")

	if err != nil {
		log.Panic(err)
	}
}

func CloseDB() {
	fmt.Println("DB closed")
	DB.Close()
}