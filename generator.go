package main

import (
  "math/rand"
   "time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetRandomPhrase(c *gin.Context) {
query := c.DefaultQuery("v", "yandere")
var greetings []Greeting
_, err := dbmap.Select(&greetings, "SELECT * FROM greeting WHERE characters=?", query)

var accosts []Accost
_, err = dbmap.Select(&accosts, "SELECT * FROM accost WHERE characters=?", query)

var questions []Question
_, err = dbmap.Select(&questions, "SELECT * FROM question WHERE characters=?", query)

rand.Seed(time.Now().UTC().UnixNano())
retstr := greetings[rand.Intn(len(greetings)-1)].Texts
retstr += " "

rand.Seed(time.Now().UTC().UnixNano())
retstr += accosts[rand.Intn(len(accosts)-1)].Texts
retstr += " "

rand.Seed(time.Now().UTC().UnixNano())
retstr += questions[rand.Intn(len(questions)-1)].Texts

if err == nil {
  c.JSON(200, retstr)
} else {
checkErr(err, "Random Failed! ")
}
}
