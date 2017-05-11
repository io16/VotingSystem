package main
//export PATH=$PATH:$GOPATH/bin
import (
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)
import (
	"./config"
	"./models"
)

// Create a GORM-backend model


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
	config.InitDB()
	config.DB.AutoMigrate(&models.VoteQuestion{}, &models.Vote{}, &models.VoteAnswerToQuestion{})
	defer config.CloseDB()

	e := echo.New()
	e.GET("/login", Hello)
	e.POST("/test", models.SaveVote)
	e.POST("/getid", models.GetVote)


	e.Logger.Fatal(e.Start(":1323"))
}