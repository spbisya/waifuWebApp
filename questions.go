package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetQuestionsChar(c *gin.Context) {
	id := c.Params.ByName("characters")
	var questions []Question
	_, err := dbmap.Select(&questions, "SELECT * FROM question WHERE characters=?", id)

	if err == nil {

		c.JSON(200, questions)
	} else {
		c.JSON(404, gin.H{"error": "questions for this query not found"})
	}

	// curl -i http://localhost:8080/api/v1/waifus/1
}

func PostQuestion(c *gin.Context) {
	if _, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
	var question Question
	c.Bind(&question)

	log.Println(question)

	if question.Characters != "" && question.Texts != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO question (characters, texts) VALUES (?, ?)`, question.Characters, question.Texts); insert != nil {
      question_id, err := insert.LastInsertId()
			if err == nil {
				content := &Question{
          Id:        question_id,
          Characters: question.Characters,
          Texts: question.Texts,
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

func UpdateQuestion(c *gin.Context) {
	if _, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
	id := c.Params.ByName("id")
	var question Question
	err := dbmap.SelectOne(&question, "SELECT * FROM question WHERE id=?", id)

	if err == nil {
		var json Question
		c.Bind(&json)

		question_id, _ := strconv.ParseInt(id, 0, 64)

		question := Question{
			Id:       question_id,
			Characters: json.Characters,
			Texts:  json.Texts,
		}

		if question.Texts != "" && question.Characters != "" {
			_, err = dbmap.Update(&question)

			if err == nil {
				c.JSON(200, question)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "question not found"})
	}
}
	 // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/waifus/1
}

func DeleteQuestion(c *gin.Context) {
	if token, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
		if token.Claims["admin"] != "true"{
			c.JSON(403, gin.H{"error": "You're not admin!"})
		} else {
	id := c.Params.ByName("id")

	var question Question
	err := dbmap.SelectOne(&question, "SELECT * FROM question WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&question)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "question not found"})
	}
}
}
	// curl -i -X DELETE http://localhost:8080/api/v1/waifus/1
}
