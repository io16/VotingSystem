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
	"encoding/json"
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

func GetUsersToVote(c echo.Context) error {
	type UsersToVote struct {
		Name string
		Time string
	}

	idVote := c.FormValue("idvote")
	userAnswers := []UserAnswer{}
	config.DB.Where("vote_id = ?", idVote).Find(&userAnswers)
	users := []UsersToVote{}
	user := User{}

	for i, item := range userAnswers {
		t := UsersToVote{}
		config.DB.Where("id = ?", item.UserID).First(&user)
		t.Name = user.Name
		t.Time = userAnswers[i].Time
		users = append(users, t)
	}
	return c.JSON(http.StatusOK, users)
}
func AddUser(c echo.Context) error {

	user := User{}
	//if err := c.Bind(user); err != nil {
	//	return err
	//}
	json.Unmarshal([]byte(c.FormValue("data")), &user)
	userStatus := isUserValid(user)
	if userStatus {
		userStatus = saveUserToDB(user)

	}
	return c.String(http.StatusOK, strconv.FormatBool(userStatus)) // if user created -- return true
}

func isUserValid(user User) bool {
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

func saveUserToDB(user User) bool {

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
	user := User{}

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

