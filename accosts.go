package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAccostsChar(c *gin.Context) {
	id := c.Params.ByName("characters")
	var accosts []Accost
	_, err := dbmap.Select(&accosts, "SELECT * FROM accost WHERE characters=?", id)

	if err == nil {

		c.JSON(200, accosts)
	} else {
		c.JSON(404, gin.H{"error": "accosts for this query not found"})
	}

	// curl -i http://localhost:8080/api/v1/waifus/1
}

func PostAccost(c *gin.Context) {
	if _, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
	var accost Accost
	c.Bind(&accost)

	log.Println(accost)

	if accost.Characters != "" && accost.Texts != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO accost (characters, texts) VALUES (?, ?)`, accost.Characters, accost.Texts); insert != nil {
      accost_id, err := insert.LastInsertId()
			if err == nil {
				content := &Accost{
          Id:        accost_id,
          Characters: accost.Characters,
          Texts: accost.Texts,
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

func UpdateAccost(c *gin.Context) {
	if _, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
	id := c.Params.ByName("id")
	var accost Accost
	err := dbmap.SelectOne(&accost, "SELECT * FROM accost WHERE id=?", id)

	if err == nil {
		var json Accost
		c.Bind(&json)

		accost_id, _ := strconv.ParseInt(id, 0, 64)

		accost := Accost{
			Id:       accost_id,
			Characters: json.Characters,
			Texts:  json.Texts,
		}

		if accost.Texts != "" && accost.Characters != "" {
			_, err = dbmap.Update(&accost)

			if err == nil {
				c.JSON(200, accost)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "accost not found"})
	}
}

	 // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/waifus/1
}

func DeleteAccost(c *gin.Context) {
	if token, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
		if token.Claims["admin"] != "true"{
			c.JSON(403, gin.H{"error": "You're not admin!"})
		} else {
	id := c.Params.ByName("id")

	var accost Accost
	err := dbmap.SelectOne(&accost, "SELECT * FROM accost WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&accost)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "accost not found"})
	}
}
}
	// curl -i -X DELETE http://localhost:8080/api/v1/waifus/1
}
