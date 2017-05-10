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
)

// Create a GORM-backend model

type Test struct {
	gorm.Model
	Name string
	Category string


}
type TestQuestion struct {
	gorm.Model
	NumberQuestion int
	ListQuestions string
	Test Test
	TestID int
}
type TestAnswerToQuestion struct {
	gorm.Model
	ListAnswers string
	NumberQuestion int
	Test Test
	TestID int
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
	config.InitDB()
	config.DB.AutoMigrate(&Test{},&TestAnswerToQuestion{},&TestQuestion{})
	defer  config.CloseDB()

	e := echo.New()
	e.GET("/login", Hello)

	e.Logger.Fatal(e.Start(":1323"))
}