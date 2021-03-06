package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	jwt "github.com/dgrijalva/jwt-go"
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

	// curl -i http://volhgroup.tk/api/v1/waifus/1
}

func GetAllGreetingsHTML(c *gin.Context) {
	GetAllPartsHTML("greeting", "Greeting", 3, c)
	// curl -i http://localhost:8080/api/v1/waifus/1
}

func GetGreetingNew(c *gin.Context) {
GetPartsNew("greeting", "Greeting", 3, c)
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
// curl -i -X POST -H "Content-Type: application/json" -d "{ \"characters\": \"genke\", \"texts\": \"друзья опять звали моего альфу гулять?\" }" http://volhgroup.tk/api/question/?token=QGWi3yN4RZjZ7TowJs6067FcWyJfgzzbou9As05Bo1@FjxtfPzj5kwLbPJJXRifqb@23m4uTpxp7D@JSEK1MZwMbgRritUFBR4MWkaULhm2iNYHiUt1egxq13OW056GtCdo3ZN9qNTr7gyW4PMGKsSgmFGIf.k6xqtVjhTtMkRnW8dDm6zyzRsHzEQVzJkK07n8O6q3LBPQgH8zC9GXQh7IWG2s91YCce.SPWenSdUqrGIuEUyiRS85KM4R1qWUPf1Y2Iclrsh9ro8W@pf4TWs28lFKVXdsAA9pp4ZtJ3dX8gbkJgF6.6fZ.GnxAlHFLRhj0fiSukUq4yWJR26pCyjvMDOGXXwS3f8OiaDrOXi5eqU1hJXywcfQ4UfdVh4vSZSdsPAE0CgazZ@R4dRlsSjv9h.Xs.u7m.2d75S
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
	 // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://volhgroup.tk/api/v1/waifus/1
}

func DeleteGreeting(c *gin.Context) {
	if token, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
		if claims := token.Claims.(jwt.MapClaims);claims["admin"] != "true"{
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
	// curl -i -X DELETE http://volhgroup.tk/api/v1/waifus/1
}
