package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"time"
	"log"
)

type JwtCustomClaims struct {
	Name string
	Role string
	ID uint
	jwt.StandardClaims
}

func GetJWT(c echo.Context) error {
	//dsadas
	fmt.Print("in func jwt ")
	type User struct {
		Login string
		Pass  string
	}
	user := User{}
	//if err := c.Bind(user); err != nil {
	//	return c.String(http.StatusInternalServerError, "")
	//}
	//json.Unmarshal([]byte(c.FormValue("data")), &user)
	user.Login = c.FormValue("login")
	user.Pass = c.FormValue("pass")


	userFromDB, err := UserAuthentication(user.Login, user.Pass)
	if err != nil {
		log.Println(err)

	} else {

		claims := JwtCustomClaims{
			userFromDB.Name,
			userFromDB.Role.UserRole,
			userFromDB.ID,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
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
			"user" : userFromDB.Login,
		})

	}

	return echo.ErrUnauthorized
}
func TestJWT(c echo.Context) error {
	return c.String(http.StatusOK,"")
}