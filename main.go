package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

)

var dbmap = initDb()


func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*.html")
//	r.StaticFile("/index.html", "index.html")
	r.StaticFS("/css", http.Dir("css"))
	r.StaticFS("/img", http.Dir("img"))
	r.StaticFS("/js", http.Dir("js"))
	r.StaticFS("/less", http.Dir("less"))
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
	r.GET("/", func(c *gin.Context) {
			    c.HTML(http.StatusOK, "index.html", gin.H{})
			})
	r.GET("/about", func(c *gin.Context) {
			    c.HTML(http.StatusOK, "about.html", gin.H{})
			})
	r.GET("/tutorial", func(c *gin.Context) {
			    c.HTML(http.StatusOK, "tutorial.html", gin.H{})
	    })
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

	v5 := r.Group("api/post")
	{
		v5.GET("/", GetPosts)
		v5.GET("/:id", GetPostsById)
		v5.POST("/", PostPost)
		v5.DELETE("/:id", DeletePost)
	}

	v6 := r.Group("post")
	{
		v6.GET("/:id", GetPostsForSiteById)
	}

	v7 := r.Group("waifus")
	{
		v7.GET("/new", func(c *gin.Context) {
						c.HTML(http.StatusOK, "waifu_add.html", gin.H{})
				})
		v7.POST("/waifuAdd", GetWaifuNew)
		v7.GET("", GetAllWaifusHTML)
	}

	v8 := r.Group("questions")
	{
		v8.GET("/new", func(c *gin.Context) {
						c.HTML(http.StatusOK, "part_add.html", gin.H{
							"Name": "Question",
							"Type": "question",
						})
				})
		v8.POST("/questionAdd", GetQuestionNew)
		v8.GET("", GetAllQuestionsHTML)
	}

	v9 := r.Group("greetings")
	{
		v9.GET("/new", func(c *gin.Context) {
						c.HTML(http.StatusOK, "part_add.html", gin.H{
							"Name": "Greeting",
							"Type": "greeting",
						})
				})
		v9.POST("/greetingAdd", GetGreetingNew)
		v9.GET("", GetAllGreetingsHTML)
	}

	v10 := r.Group("accosts")
	{
		v10.GET("/new", func(c *gin.Context) {
						c.HTML(http.StatusOK, "part_add.html", gin.H{
							"Name": "Accost",
							"Type": "accost",
						})
				})
		v10.POST("/accostAdd", GetAccostNew)
		v10.GET("", GetAllAccostsHTML)
	}

	r.GET("/randompost", GetRandomPost)

	r.GET("/allposts", GetAllPosts)

	r.GET("/random", GetRandomPhrase)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "caps.png.html", gin.H{
				"title": "Main website",
		})
	})
	r.Run(":80")
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
