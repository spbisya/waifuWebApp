package main

import (
	"log"
   "math"
	"strconv"
	jwt "github.com/dgrijalva/jwt-go"
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

	// curl -i http://volhgroup.tk/api/v1/waifus
}

func GetWaifuNew(c *gin.Context) {
	firstname := c.PostForm("first")
	lastname := c.PostForm("last")
	character := c.PostForm("character")
	if firstname != "" && lastname != "" && character != ""{
		if insert, _ := dbmap.Exec(`INSERT INTO waifu (firstname, lastname, characters) VALUES (?, ?, ?)`, firstname, lastname, character); insert != nil {
      waifu_id, err := insert.LastInsertId()
			if err == nil {
				content := &Waifu{
          Id:        waifu_id,
          Firstname: firstname,
          Lastname:  lastname,
          Characters: character,
				}
				c.HTML(201, "waifu.html", content)
			} else {
				c.JSON(201, gin.H{"error": err.Error()})
			}
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
		// curl -i http://volhgroup.tk/api/v1/waifus
}


func GetAllWaifusHTML(c *gin.Context){
  query:=c.DefaultQuery("page", "1")
  page, _:= strconv.Atoi(query);
  var posts []Waifu
  _, err := dbmap.Select(&posts, "SELECT * FROM waifu")
  posts = ReverseArrayW(posts)
  pagesCount:=int(math.Ceil(float64(len(posts))/10))

  switch page {
  case 1:
  isNext:=CheckIfHasNextW(page, len(posts))
  isLast:=!isNext
	more:=false
  if pagesCount>5 {
    pagesCount = 5
		more =true
  }


  array:=make([]Page, pagesCount)
  for i:=1;i<=pagesCount;i++ {
    content:=false
    if i == 1{
      content=true
    }
  array[i-1] = Page{
    Number: i,
    Current: content,
    }
  }
  	if err == nil {
			posts1:=posts
			if len(posts)<10 {
				posts1=posts[0:len(posts)]
			}	else{
				posts1 = posts[0:10]
			}
      c.HTML(200, "allWaifus.html", gin.H{
                  "Waifus": posts1,
                  "Page": 1,
                  "Next": 2,
                  "Previous": 0,
                  "IsNext": isNext,
                  "IsPrevious": false,
                  "Count": array,
                  "IsFirst": true,
                  "IsLast": isLast,
                  "Last": int(math.Ceil(float64(len(posts))/10)),
                  "More": more,
                  "Less": false,
              })
    } else {
      checkErr(err, "Couldn't get Waifus")
      c.JSON(404, gin.H{"error": "Couldn't get waifus"})
    }

  case 2:
  isNext:=CheckIfHasNextW(page, len(posts))
  isLast:=!isNext
	more:=false
	if pagesCount>5 {
		pagesCount = 5
		more =true
	}
  array:=make([]Page, pagesCount)
  for i:=1;i<=pagesCount;i++ {
    content:=false
    if i == 2{
      content=true
    }
  array[i-1] = Page{
    Number: i,
    Current: content,
    }
  }
    if err == nil {
			posts1:=posts
			if len(posts)<20 {
				posts1=posts[10:len(posts)]
			}	else{
				posts1 = posts[0:20]
			}
      c.HTML(200, "allWaifus.html", gin.H{
                  "Waifus": posts1,
                  "Page": 2,
                  "Next": 3,
                  "Previous": 1,
                  "IsNext": isNext,
                  "IsPrevious": true,
                  "Count": array,
                  "IsFirst": true,
                  "IsLast": isLast,
                  "Last": int(math.Ceil(float64(len(posts))/10)),
                  "More": more,
                  "Less": false,
              })
    } else {
      checkErr(err, "Couldn't get waifus")
      c.JSON(404, gin.H{"error": "Couldn't get waifus"})
    }

  case 3:
  isNext:=CheckIfHasNextW(page, len(posts))
  isLast:=!isNext
	more:=false
	if pagesCount>5 {
		pagesCount = 5
		more =true
	}
  array:=make([]Page, pagesCount)
  for i:=1;i<=pagesCount;i++ {
    content:=false
    if i == 3{
      content=true
    }
  array[i-1] = Page{
    Number: i,
    Current: content,
    }
  }
    if err == nil {
			posts1:=posts
			if len(posts)<30 {
				posts1=posts[20:len(posts)]
			}	else{
				posts1 = posts[20:30]
			}
			c.HTML(200, "allWaifus.html", gin.H{
									"Waifus": posts1,
                  "Page": 3,
                  "Next": 4,
                  "Previous": 3,
                  "IsNext": isNext,
                  "IsPrevious": true,
                  "Count": array,
                  "IsFirst": true,
                  "IsLast": isLast,
                  "Last": int(math.Ceil(float64(len(posts))/10)),
                  "More": more,
                  "Less": false,
              })
    } else {
			checkErr(err, "Couldn't get waifus")
			c.JSON(404, gin.H{"error": "Couldn't get waifus"})
    }
  default:
    pagesCount=3;
  isNext:=CheckIfHasNextW(page, len(posts))
  isLast:=!CheckIfHasNextW(page, len(posts))
  more:=CheckIfHasNextW(page+1, len(posts))
  less:=CheckIfHasNextW(page-1, len(posts))
      if err == nil {
				posts1:=posts
				if len(posts)<page*10 {
					posts1=posts[(page-1)*10:len(posts)]
				}	else{
					posts1 = posts[(page-1)*10:page*10]
				}
	      c.HTML(200, "allWaifus.html", gin.H{
	                  "Waifus": posts1,
                    "Page": page,
                    "Next": page+1,
                    "Previous": page-1,
                    "IsNext": isNext,
                    "IsPrevious": true,
                    "Count": GetArray(len(posts), page),
                    "IsFirst": false,
                    "IsLast": isLast,
                    "Last": int(math.Ceil(float64(len(posts))/10)),
                    "More": more,
                    "Less": less,
                })
      } else {
				checkErr(err, "Couldn't get waifus")
				c.JSON(404, gin.H{"error": "Couldn't get waifus"})
      }
    }
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

	// curl -i http://volhgroup.tk/api/v1/waifus/1
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
	// curl -i http://volhgroup.tk/api/v1/waifus/1
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
	//  curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Misato\", \"lastname\": \"Sakurai\", \"characters\": \"genke\" }" http://volhgroup.tk/api/waifus/?token=QGWi3yN4RZjZ7TowJs6067FcWyJfgzzbou9As05Bo1@FjxtfPzj5kwLbPJJXRifqb@23m4uTpxp7D@JSEK1MZwMbgRritUFBR4MWkaULhm2iNYHiUt1egxq13OW056GtCdo3ZN9qNTr7gyW4PMGKsSgmFGIf.k6xqtVjhTtMkRnW8dDm6zyzRsHzEQVzJkK07n8O6q3LBPQgH8zC9GXQh7IWG2s91YCce.SPWenSdUqrGIuEUyiRS85KM4R1qWUPf1Y2Iclrsh9ro8W@pf4TWs28lFKVXdsAA9pp4ZtJ3dX8gbkJgF6.6fZ.GnxAlHFLRhj0fiSukUq4yWJR26pCyjvMDOGXXwS3f8OiaDrOXi5eqU1hJXywcfQ4UfdVh4vSZSdsPAE0CgazZ@R4dRlsSjv9h.Xs.u7m.2d75S
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
	 // curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://volhgroup.tk/api/v1/waifus/1
}

func DeleteWaifu(c *gin.Context) {
	if token, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
	if claims := token.Claims.(jwt.MapClaims);claims["admin"] != "true"{
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
	// curl -i -X DELETE http://volhgroup.tk/api/v1/waifus/1
}
