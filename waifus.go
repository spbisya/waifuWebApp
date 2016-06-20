package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetWaifus(c *gin.Context) {
	var waifus []Waifu
	_, err := dbmap.Select(&waifus, "SELECT * FROM waifu")

	if err == nil {
		c.JSON(200, waifus)
	} else {
		c.JSON(404, gin.H{"error": "no waifu(s) into the table"})
	}

	// curl -i http://localhost:8080/api/v1/waifus
}

func GetWaifu(c *gin.Context) {
	id := c.Params.ByName("id")
	var waifu Waifu
	err := dbmap.SelectOne(&waifu, "SELECT * FROM waifu WHERE id=? LIMIT 1", id)

	if err == nil {
		waifu_id, _ := strconv.ParseInt(id, 0, 64)

		content := &Waifu{
			Id:        waifu_id,
			Firstname: waifu.Firstname,
			Lastname:  waifu.Lastname,
      Characters: waifu.Characters,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "waifu not found"})
	}

	// curl -i http://localhost:8080/api/v1/waifus/1
}

func GetWaifuByCharacter(c *gin.Context) {
	character := c.Params.ByName("characters")
	var waifus []Waifu
	_, err := dbmap.Select(&waifus, "SELECT * FROM waifu WHERE characters=? ", character)

	if err == nil {
		c.JSON(200, waifus)
	} else {
		c.JSON(404, gin.H{"error": "no waifu(s) into the table"})
	}
	// curl -i http://localhost:8080/api/v1/waifus/1
}

func PostWaifu(c *gin.Context) {
	if _, err := CheckAndDecodeToken(c.Query("token")); err != nil {
 	 c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
	var waifu Waifu
	c.Bind(&waifu)

	log.Println(waifu)

	if waifu.Firstname != "" && waifu.Lastname != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO waifu (firstname, lastname, characters) VALUES (?, ?, ?)`, waifu.Firstname, waifu.Lastname, waifu.Characters); insert != nil {
      waifu_id, err := insert.LastInsertId()
			if err == nil {
				content := &Waifu{
          Id:        waifu_id,
          Firstname: waifu.Firstname,
          Lastname:  waifu.Lastname,
          Characters: waifu.Characters,
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
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\", \"characters\": \"yandere\" }" http://localhost:8080/api/waifus/?token=gD8DgA265rCFsOmPLSM0tk3TCKENFYeUV8IFd4Un14.xuyInMlpXCmpvpfAjRVzaMhxMBHrBmzv67LxMqJhtu5gqEKfpTQqtbV.Aq3t7Nl5Q5d7b5E5VZe7qDOVeuZR8ud9mPXBjJ7wSL61azkuuCiscVpjL6r5xbaoSy7H2ZyAfg1k54AkqYLj@5T@oqC9xbcT55R1ZyeQFqNAh.FENztmnN0huv99OBC2pX2H@wA1U3VefuUoSg0YCtUIb@hDXsnM.aRS0PaAaJeh8nJ1b7Snxxmihf4JWVoW44OEli3muA2x7YVhRl8nbCpJWhwioZnKUDpXpaeS5aO4FeyOigRBkvptn5GpcxQEa6@j2W4ezZSVjlT2ny7u0@5han1ANVOVGEXSiW3nwklJq3ZEr@qogDHjkUSSnFVYn6E
}

func UpdateWaifu(c *gin.Context) {
	if _, err := CheckAndDecodeToken(c.Query("token")); err != nil {
 	  c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
	id := c.Params.ByName("id")
	var waifu Waifu
	err := dbmap.SelectOne(&waifu, "SELECT * FROM waifu WHERE id=?", id)

	if err == nil {
		var json Waifu
		c.Bind(&json)

		waifu_id, _ := strconv.ParseInt(id, 0, 64)

		waifu := Waifu{
			Id:       waifu_id,
			Firstname: json.Firstname,
			Lastname:  json.Lastname,
      Characters: json.Characters,
		}

		if waifu.Firstname != "" && waifu.Lastname != "" && waifu.Characters != "" {
			_, err = dbmap.Update(&waifu)

			if err == nil {
				c.JSON(200, waifu)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "waifu not found"})
	}
}
	 // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/waifus/1
}

func DeleteWaifu(c *gin.Context) {
	if token, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
		if token.Claims["admin"] != "true"{
			c.JSON(403, gin.H{"error": "You're not admin!"})
		} else {
	id := c.Params.ByName("id")

	var waifu Waifu
	err := dbmap.SelectOne(&waifu, "SELECT * FROM waifu WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&waifu)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "waifu not found"})
	}
}
}
	// curl -i -X DELETE http://localhost:8080/api/v1/waifus/1
}
