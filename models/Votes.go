package models

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"../config"
	"net/http"
	"github.com/jmoiron/sqlx/types"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"fmt"
	"time"
)

type Vote struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:255"`
	Category string `json:"category" gorm:"size:255"`
}
type VoteQuestion struct {
	gorm.Model
	Question types.JSONText
	Vote     Vote
	VoteID   uint
}
type VoteAnswerToQuestion struct {
	gorm.Model
	Answer types.JSONText
	Vote   Vote
	VoteID uint
}
type VoteStats struct {
	gorm.Model
	VoteID    uint
	Vote      Vote
	VoteStats types.JSONText
}
type UserAnswer struct {
	gorm.Model
	UserID     uint
	User       User
	UserAnswer types.JSONText
	Vote       Vote
	VoteID     int
	Time       string
}

type saveUserVoteParse struct {
	Vote   []struct{ Answer []string  `json:"answer"` }`json:"vote"`
	VoteID int `json:"voteid"`
}

func SaveUserVote(c echo.Context) error {

	time.Now()
	//t := new(saveUserVoteParse)
	//
	//if err := c.Bind(t); err != nil {
	//	return err
	//
	//}
	t := saveUserVoteParse{}

	json.Unmarshal([]byte(c.FormValue("data")), &t)
	fmt.Println(c.FormValue("data"))
	updateVoteStats(t)
	userAnswer := UserAnswer{}
	byteJson, err := json.Marshal(t)
	if err != nil {
		log.Printf("\n%s", err)
	}
	userAnswer.UserAnswer = byteJson
	userAnswer.VoteID = t.VoteID

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	userAnswer.UserID = claims.ID

	time := time.Now()
	//fmt.Println(time.String())
	//fmt.Println(time.Format("2006-01-02 15:04:05"))
	//
	userAnswer.Time = time.Format("2006-01-02 15:04:05")

	config.DB.Create(&userAnswer)
	return c.JSON(http.StatusOK, t)
}

type voteParse struct {
	Name             string //`json:"name" gorm:"size:255"`
	Category         string //`json:"category" gorm:"size:255"`
	Questions        struct {
				 Question []string //`json:"questions" gorm:"size:255"`
				 Type     []string
			 }
	AnswerToQuestion [][]string
}

func SaveVote(c echo.Context) error {

	t := voteParse{}

	json.Unmarshal([]byte(c.FormValue("data")), &t)


	//if err := c.Bind(t); err != nil {
	//	return err
	//}

	vote := Vote{}
	vote.Name = t.Name;
	vote.Category = t.Category

	byteJson, err := json.Marshal(t.Questions)
	if err != nil {
		log.Printf("\n%s", err)
	}

	voteQuestion := VoteQuestion{}
	voteQuestion.Question = byteJson

	byteJson, err = json.Marshal(t.AnswerToQuestion)
	if err != nil {
		log.Printf("\n%s", err)
	}
	//fmt.Print(t.AnswerToQuestion)
	voteAnswToQuest := VoteAnswerToQuestion{}

	voteAnswToQuest.Answer = byteJson

	config.DB.Create(&vote)

	voteQuestion.VoteID = vote.ID
	voteAnswToQuest.VoteID = vote.ID

	config.DB.Create(&voteAnswToQuest)
	config.DB.Create(&voteQuestion)
 	createVoteStatsRow(vote.ID, voteAnswToQuest)

	return c.String(http.StatusOK, "ok")
}

func GetVote(c echo.Context) error {

	type VoteRequest struct {
		IdVote int
	}
	//json.Unmarshal([]byte(c.FormValue("idvote")), &request)
	//request := VoteRequest{}
	//if err := c.Bind(request); err != nil {
	//	return err
	//
	//}

	request := c.FormValue("idvote")
	vote := Vote{}
	voteQuest := VoteQuestion{}
	voteAnswer := VoteAnswerToQuestion{}

	config.DB.Where("id = ?", request).First(&vote)
	config.DB.Where("vote_id = ?", request).Find(&voteQuest)
	config.DB.Where("vote_id = ?", request).Find(&voteAnswer)

	jsonToSend := new(voteParse)

	jsonToSend.Name = vote.Name
	jsonToSend.Category = vote.Category

	voteQuest.Question.Unmarshal(&jsonToSend.Questions)
	voteAnswer.Answer.Unmarshal(&jsonToSend.AnswerToQuestion)

	return c.JSON(http.StatusOK, jsonToSend)
}
func GetVotes(c echo.Context) error {
	vote := []Vote{}
	type Votes struct {
		Vote []Vote
	}
	config.DB.Find(&vote)
	votes := Votes{vote}

	return c.JSON(http.StatusOK, votes)
}

func GetVotesStats(c echo.Context) error {
	voteStat := VoteStats{}
	type VoteRequest struct {
		VoteId int
	}
	//request := new(VoteRequest)
	//if err := c.Bind(request); err != nil {
	//	return err
	//
	//}
	request := c.FormValue("voteid")
	config.DB.Where("vote_id = ?", request).First(&voteStat)

	return c.JSON(http.StatusOK, voteStat.VoteStats)
}

type voteStatsParse struct {
	Question []struct {
		CountAnswers int
		Stats        []int
	}`json:"question"`
}

func createVoteStatsRow(voteId uint, answers VoteAnswerToQuestion) {

	voteStats := voteStatsParse{}
	var voteansw [][]string

	answers.Answer.Unmarshal(&voteansw)

	for key, answer := range voteansw {

		var temp struct{ CountAnswers int; Stats []int }

		voteStats.Question = append(voteStats.Question, temp)
		voteStats.Question[key].CountAnswers = 0

		for i, _ := range answer {
			if i == 0 {
				voteStats.Question[key].Stats = make([]int, len(answer))
			}
			voteStats.Question[key ].Stats[i] = 0
		}
	}
	byteJson, err := json.Marshal(voteStats)
	if err != nil {
		log.Printf("\n%s", err)
	}
	v := VoteStats{}
	v.VoteStats = byteJson
	v.VoteID = voteId
	config.DB.Create(&v)

}
func updateVoteStats(userVote saveUserVoteParse) {

	voteStats := VoteStats{}
	voteStatsParse := voteStatsParse{}

	config.DB.Where("vote_id = ?", userVote.VoteID).First(&voteStats)

	voteStats.VoteStats.Unmarshal(&voteStatsParse)

	for key, objAnswers := range voteStatsParse.Question {
		voteStatsParse.Question[key].CountAnswers++
		for i, _ := range objAnswers.Stats {
			if userVote.Vote[key].Answer[i] == "1" {

				voteStatsParse.Question[key].Stats[i]++
			}
		}
	}

	byteJson, err := json.Marshal(voteStatsParse)
	if err != nil {
		log.Printf("\n%s", err)
	}
	voteStats.VoteStats = byteJson
	config.DB.Save(&voteStats)
}