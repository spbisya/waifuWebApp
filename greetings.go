package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetGreetingsChar(c *gin.Context) {
	id := c.Params.ByName("characters")
	var greetings []Greeting
	_, err := dbmap.Select(&greetings, "SELECT * FROM greeting WHERE characters=?", id)

	if err == nil {

		c.JSON(200, greetings)
	} else {
		c.JSON(404, gin.H{"error": "greetings for this query not found"})
	}

	// curl -i http://localhost:8080/api/v1/waifus/1
}

func PostGreeting(c *gin.Context) {
	if _, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
	var greeting Greeting
	c.Bind(&greeting)

	log.Println(greeting)

	if greeting.Characters != "" && greeting.Texts != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO greeting (characters, texts) VALUES (?, ?)`, greeting.Characters, greeting.Texts); insert != nil {
      greeting_id, err := insert.LastInsertId()
			if err == nil {
				content := &Greeting{
          Id:        greeting_id,
          Characters: greeting.Characters,
          Texts: greeting.Texts,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}

		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\", \"characters\": \"yandere\" }" http://localhost:8080/api/v1/waifus
}

func UpdateGreeting(c *gin.Context) {
	if _, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
	id := c.Params.ByName("id")
	var greeting Greeting
	err := dbmap.SelectOne(&greeting, "SELECT * FROM greeting WHERE id=?", id)

	if err == nil {
		var json Greeting
		c.Bind(&json)

		greeting_id, _ := strconv.ParseInt(id, 0, 64)

		greeting := Greeting{
			Id:       greeting_id,
			Characters: json.Characters,
			Texts:  json.Texts,
		}

		if greeting.Texts != "" && greeting.Characters != "" {
			_, err = dbmap.Update(&greeting)

			if err == nil {
				c.JSON(200, greeting)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "greeting not found"})
	}
}
	 // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/waifus/1
}

func DeleteGreeting(c *gin.Context) {
	if token, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
		if token.Claims["admin"] != "true"{
			c.JSON(403, gin.H{"error": "You're not admin!"})
		} else {
	id := c.Params.ByName("id")

	var greeting Greeting
	err := dbmap.SelectOne(&greeting, "SELECT * FROM greeting WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&greeting)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "greeting not found"})
	}
}
}
	// curl -i -X DELETE http://localhost:8080/api/v1/waifus/1
}
