package main
//export PATH=$PATH:$GOPATH/bin
import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)
import (
	"./config"
	"./models"
	"github.com/labstack/echo/middleware"

	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/gorilla/sessions"
	"time"
)

// Create a GORM-backend model


// Create another GORM-backend model
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Test(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		body, errJson := ioutil.ReadAll(c.Request().Body)
		var test map[string]string
		errJson = json.Unmarshal([]byte(body), &test)
		if errJson != nil {
			log.Print(errJson)
		}
		if test != nil {
			log.Print(test)
		}

		cookies, errCookies := c.Request().Cookie("session-name")
		if errCookies != nil {
			log.Print(errCookies)
		}
		if errCookies != nil && errJson != nil {
			session, err := store.Get(c.Request(), "session-name")
			if err != nil {

				return c.String(http.StatusInternalServerError, "error")
			}
			session.Values["foo"] = "bar"
			session.Save(c.Request(), c.Response())
			fmt.Println("////")
			fmt.Println(session.Values["foo"])
		}
		var s string
		if errJson == nil {
			cookie := new(http.Cookie)
			cookie.Name = "session-name"
			cookie.Value = test["key"]
			cookie.Expires = time.Now().Add(24 * time.Hour)
			c.Request().AddCookie(cookie)

			session, _ := store.Get(c.Request(), "session-name")
			s = session.Values["foo"].(string)
			fmt.Println(session.Values["foo"])

			session.Values["foo"] = test["name"]
			session.Save(c.Request(), c.Response())

		}

		fmt.Printf(cookies.Value)
		//c.Request().AddCookie(cookies)

		fmt.Println(test["key"])
		return c.String(http.StatusOK,s )
	}
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		//session, err := store.Get(c.Request(), "session-name")
		//if err != nil {
		//
		//	return c.String(http.StatusInternalServerError, "error")
		//}
		//fmt.Println(session.Values["name"])
		//session.Values["name"] = "1"
		return c.String(http.StatusRequestedRangeNotSatisfiable, "ok")
	}

}

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.VoteQuestion{}, &models.Vote{}, &models.VoteAnswerToQuestion{}, models.User{}, models.Role{})
	defer config.CloseDB()

	e := echo.New()

	// Middleware
	//e.Use(ServerHeader)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: " time=${time_rfc3339} method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	//e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
	//	StackSize:  1 << 10, // 1 KB
	//}))
	e.GET("/testsession", models.UserAuthorization, Test)
	e.POST("/testsession", models.UserAuthorization, Test)
	e.POST("/login", models.UserAuthorization)
	e.POST("/savevote", models.SaveVote)
	e.POST("/getvote", models.GetVote)

	e.POST("/adduser", models.AddUser)

	e.Logger.Fatal(e.Start(":1323"))
}