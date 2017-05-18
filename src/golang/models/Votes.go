package models

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"../config"
	"net/http"
	"fmt"

	"github.com/jmoiron/sqlx/types"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"github.com/labstack/gommon/log"
)

type requestJson struct {
	Name             string //`json:"name" gorm:"size:255"`
	Category         string //`json:"category" gorm:"size:255"`
	Questions        struct {
				 Question []string //`json:"questions" gorm:"size:255"`
			 }
	AnswerToQuestion map[int][]string
}

type Vote struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:255"`
	Category string `json:"category" gorm:"size:255"`
}
type VoteQuestion struct {
	gorm.Model
	NumberQuestion int
	Question       string
	Vote           Vote
	VoteID         uint
}
type VoteAnswerToQuestion struct {
	gorm.Model
	Answer         string
	NumberQuestion int
	Vote           Vote
	VoteID         uint
}
type VoteStats struct {
	gorm.Model
	VoteID uint
	Vote Vote
	VoteStats types.JSONText

}
type UserAnswer struct {
	gorm.Model
	UserID     uint
	User       User
	UserAnswer types.JSONText
	Vote	Vote
	VoteID	int
}


type request struct {
	Vote []struct{ Answer []string  `json:"answer"` }`json:"vote"`
	VoteID int
}
func SaveUserVote(c echo.Context) error {


	t := new(request)

	if err := c.Bind(t); err != nil {
		return err

	}
	userAnswer := UserAnswer{}
	name, err := json.Marshal(t)
	if err != nil {
		log.Printf("\n%s", err)
	}
	userAnswer.UserAnswer = name
	userAnswer.VoteID = t.VoteID


	log.Print(userAnswer.UserAnswer)
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
	userAnswer.UserID = claims.ID

	config.DB.Create(&userAnswer)
	return c.JSON(http.StatusOK, t)
}
func SaveVote(c echo.Context) error {
	// TODO записувати шаблон в таблицю votestats
	//TODO змінить спосіб запису в БД на json формат
	t := new(requestJson)

	if err := c.Bind(t); err != nil {
		return err

	}

	vote := Vote{}

	vote.Name = t.Name;
	vote.Category = t.Category

	config.DB.Create(&vote)

	for i, item := range t.Questions.Question {
		voteQuestion := VoteQuestion{}

		voteQuestion.Question = item
		voteQuestion.NumberQuestion = i
		voteQuestion.VoteID = vote.ID

		config.DB.Create(&voteQuestion)

	}
	for i, arr := range t.AnswerToQuestion {
		for _, item := range arr {
			voteAnswToQuest := VoteAnswerToQuestion{}

			voteAnswToQuest.VoteID = vote.ID
			voteAnswToQuest.Answer = item
			fmt.Println(i)
			voteAnswToQuest.NumberQuestion = i

			config.DB.Create(&voteAnswToQuest)
		}
	}

	return c.String(http.StatusOK, "ok")
}

func GetVote(c echo.Context) error {
	//TODO змінить спосіб отримання з БД на json формат
	type VoteRequest struct {
		IdVote int
	}
	t := new(VoteRequest)
	if err := c.Bind(t); err != nil {
		return err

	}
	vote := Vote{}
	voteQuest := []VoteQuestion{}
	voteAnswer := []VoteAnswerToQuestion{}

	config.DB.Where("vote_id = ?", t.IdVote).Find(&voteQuest)
	config.DB.Where("vote_id = ?", t.IdVote).Find(&voteAnswer)
	config.DB.Where("id = ?", t.IdVote).First(&vote)

	jsonToSend := new(requestJson)

	jsonToSend.Name = vote.Name
	jsonToSend.Category = vote.Category
	jsonToSend.Questions.Question = make([]string, len(voteQuest))

	for i, item := range voteQuest {
		jsonToSend.Questions.Question[i] = item.Question

	}

	test := make(map[int][]string)
	for _, q := range voteQuest {
		var tempArray[] string
		for _, a := range voteAnswer {
			if a.NumberQuestion == q.NumberQuestion {
				tempArray = append(tempArray, a.Answer)
			}

		}
		test[q.NumberQuestion] = tempArray
	}
	jsonToSend.AnswerToQuestion = test

	return c.JSON(http.StatusOK, jsonToSend)

}

func updateVoteStats(request request) {
//TODO  оновлювати інфу після отримання нового голосу
	voteStats :=VoteStats{}
	config.DB.Where("id = ?", request.VoteID).First(&voteStats)

}