package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

)

var dbmap = initDb()


func main() {
	r := gin.Default()
r.LoadHTMLGlob("templates/*.tmpl")
	r.Use(Cors())

	r.POST("/register", Register)

	v1 := r.Group("api/waifus")
	{
		v1.GET("/", GetWaifus)
		v1.GET("/:characters", GetWaifuByCharacter)
		v1.POST("/", PostWaifu)
		v1.PUT("/:id", UpdateWaifu)
		v1.DELETE("/:id", DeleteWaifu)
	}

  r.GET("/words", GetWords)

  v2 := r.Group("api/greeting")
	{
		v2.GET("/:characters", GetGreetingsChar)
		v2.POST("/", PostGreeting)
		v2.PUT("/:id", UpdateGreeting)
		v2.DELETE("/:id", DeleteGreeting)
	}

  v3 := r.Group("api/accost")
	{
		v3.GET("/:characters", GetAccostsChar)
		v3.POST("/", PostAccost)
		v3.PUT("/:id", UpdateAccost)
		v3.DELETE("/:id", DeleteAccost)
	}

  v4 := r.Group("api/question")
	{
		v4.GET("/:characters", GetQuestionsChar)
		v4.POST("/", PostQuestion)
		v4.PUT("/:id", UpdateQuestion)
		v4.DELETE("/:id", DeleteQuestion)
	}

	r.GET("/random", GetRandomPhrase)

	r.Run(":8080")
}





func GetWords(c *gin.Context) {
 query := c.DefaultQuery("v", "greetings")
	switch query {
	case "greetings":
		var greetings []Greeting
		_, err := dbmap.Select(&greetings, "SELECT * FROM greeting")

		if err == nil && len(greetings)>0 {
			c.JSON(200, greetings)
		} else {
			c.JSON(404, gin.H{"error": "no greeting(s) into the table"})
		}
	case "accosts":
		var accosts []Accost
		_, err := dbmap.Select(&accosts, "SELECT * FROM accost")

		if err == nil && len(accosts)>0 {
			c.JSON(200, accosts)
		} else {
			c.JSON(404, gin.H{"error": "no accost(s) into the table"})
		}
	case "questions":
		var questions []Question
		_, err := dbmap.Select(&questions, "SELECT * FROM question")

		if err == nil && len(questions)>0 {
			c.JSON(200, questions)
		} else {
			c.JSON(404, gin.H{"error": "no question(s) into the table"})
		}
	}
	// curl -X GET 'http://localhost:8080/words' -d v=greetings -d token=dfdf
}
