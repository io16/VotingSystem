package main

import (

	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
)

// Create a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Create another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
	Type        int
}

func Hello(c echo.Context) error {

	return c.String(http.StatusOK, "hello")
}

func main() {
	DB, err := gorm.Open("postgres", "host=localhost user=postgres dbname=AIPPZ sslmode=disable password=299792458")
	if err != nil {
		log.Panic(err)
	}
	DB.AutoMigrate(&User{}, &Product{})
	e := echo.New()
	e.GET("/login", Hello)

	e.Logger.Fatal(e.Start(":1323"))
}