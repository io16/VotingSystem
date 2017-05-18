package models

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"regexp"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"log"
	"../config"
	"github.com/gorilla/sessions"
)

type User struct {
	gorm.Model
	Name   string `json:"name" gorm:"size:255"`
	Login  string `json:"login" gorm:"size:255"`
	Email  string `json:"email"gorm:"size:255"`
	Pass   string `json:"pass"gorm:"size:255"`
	Role   Role
	RoleID int
}

type Role struct {
	gorm.Model
	UserRole string
}

func AddUser(c echo.Context) error {

	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	userStatus := isUserValid(user)
	if userStatus {
		userStatus = saveUserToDB(user)

	}
	return c.String(http.StatusOK, strconv.FormatBool(userStatus)) // if user created -- return true
}

func isUserValid(user *User) bool {
	validationStatus := false
	r, _ := regexp.Compile("^[A-Za-z0-9_-]{3,50}$")

	if r.MatchString(user.Login) &&  r.MatchString(user.Name) && r.MatchString(user.Pass) && isEmailValid(user.Email) {
		validationStatus = true
	} else {
		log.Print("user invalid")
	}

	return validationStatus

}

func isEmailValid(email string) bool {
	r, _ := regexp.Compile("^[a-zA-z_]([A-Za-z0-9_.-]{0,100})@([a-z]{2,8}[.][a-z]{2,8})$")

	return r.MatchString(email)
}

func saveUserToDB(user *User) bool {

	temp := User{}

	config.DB.Where("login = ?", user.Login).First(&temp)

	if (temp.ID == 0) {

		if user.Role.UserRole == "" {
			user.RoleID = 1;

		}
		password := []byte (user.Pass)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		user.Pass = string(hashedPassword)

		config.DB.Create(&user)
		fmt.Println("user Created")
		return true

	} else {
		fmt.Println("user is alerady exist")

		return false

	}
	return true
}

func UserAuthentication(login, password string) (User, error) {
	user :=User{}



	config.DB.Where("login = ?", login).First(&user)
	config.DB.Model(user).Related(&user.Role)

	err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(password))
	fmt.Println(user)
	if err != nil {
		fmt.Print(err)
		return user, err
	}

	return user, nil
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func MyHandler(c echo.Context) error {

	name :=c.QueryParam("name")
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(c.Request(), "session-name")
	if err != nil {

		return c.String(http.StatusInternalServerError,"error")
	}


	// Set some session values.
	fmt.Println(name)
	fmt.Println(session.Values["name"])
	session.Values["name"] = name
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	//session.Save(c.Request(), c.Response())
	return c.String(http.StatusOK,"ok")

}
