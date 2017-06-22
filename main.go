package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"./config"
	"./models"
	"github.com/labstack/echo/middleware"

	"fmt"
)


func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Print("in func middleware ")
		return next(c)
	}

}

func ServerHeaderPre(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Print("in func middleware Pre ")
		return next(c)
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

	//e.Use(	ServerHeader)
	//e.Pre(	ServerHeaderPre)
	//e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
	//	StackSize:  1 << 10, // 1 KB
	//}))
	//e.GET("/testsession", models.UserAuthorization, Test)
	//e.POST("/testsession", models.UserAuthorization, Test)
	//e.POST("/login", models.UserAuthorization)
	e.Static("/", "assets")

	//e.File("/", "login.html")
	e.File("/", "public/dashboard.html")
	//e.File("/", "public/index.html")
	e.File("/signin", "public/signin.html")
	e.File("/registration", "public/registration.html")
	e.File("/navbar", "public/navbar.html")

	e.File("/createtest","public/createTest.html")

	e.POST("/getjwt", models.GetJWT)

	e.POST("/savevote", models.SaveVote)
	e.POST("/getvote", models.GetVote)
	e.GET("/getvotes", models.GetVotes)
	e.POST("/getuserstovote", models.GetUsersToVote)


	e.POST("/getvotestat",models.GetVotesStats)

	e.POST("/adduser", models.AddUser)
	e.POST("/isUserCompleteTest", models.IsUserCompleteTest)

	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	jwtConfig := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(jwtConfig))
	//r.POST("/getjwt", models.GetJWT,ServerHeader)
	r.POST("/saveuservote", models.SaveUserVote)
	r.POST("/test",models.TestJWT)
	e.Logger.Fatal(e.Start(":1323"))
}