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

	"github.com/gorilla/sessions"
	"time"
	"github.com/dgrijalva/jwt-go"
	"log"
)

// Create a GORM-backend model




func GetJWT(c echo.Context) error {
	type User struct {
		Login string  `json:"login"`
		Pass  string `json:"pass"`
	}
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	fmt.Println(user.Pass)

	userFromDB, err := models.UserAuthentication(user.Login, user.Pass)
	if err != nil {
		log.Println(err)

	} else {

		claims := &config.JwtCustomClaims{
			userFromDB.Name,
			userFromDB.Role.UserRole,
			userFromDB.ID,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1000).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})

	}

	return echo.ErrUnauthorized
}
func restricted(c echo.Context) error {
	//body, errJson := ioutil.ReadAll(c.Request().Body)
	//var test map[string]interface{}
	//errJson = json.Unmarshal([]byte(body), &test)
	//if errJson != nil {
	//	log.Println(errJson)
	//}
	//log.Println(test["key"])
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
	name := claims.Name
	log.Println(claims.Name)
	if claims.Role != "admin" {
		return c.String(http.StatusOK, "Welcome " + name + "! access denied" + " your status is " + claims.Role)
	}
	return c.String(http.StatusOK, "Welcome " + name + "your status is " + claims.Role)
}

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
		return c.String(http.StatusOK, s)
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
	config.DB.AutoMigrate(&models.VoteQuestion{}, &models.Vote{}, &models.VoteAnswerToQuestion{},
		&models.User{}, &models.Role{}, &models.UserAnswer{}, &models.VoteStats{})
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
	//e.GET("/testsession", models.UserAuthorization, Test)
	//e.POST("/testsession", models.UserAuthorization, Test)
	//e.POST("/login", models.UserAuthorization)
	e.POST("/getjwt", GetJWT)
	e.POST("/savevote", models.SaveVote)
	e.POST("/getvote", models.GetVote)

	e.POST("/adduser", models.AddUser)


	r := e.Group("/restricted")


	// Configure middleware with the custom claims type
	jwtConfig := middleware.JWTConfig{
		Claims:     &config.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(jwtConfig))
	r.POST("", restricted)
	r.POST("/adduservote", models.SaveUserVote)
	e.Logger.Fatal(e.Start(":1323"))
}