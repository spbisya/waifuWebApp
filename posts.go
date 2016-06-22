package main

import (
	"log"
"time"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetPosts(c *gin.Context) {
	var posts []Post
	_, err := dbmap.Select(&posts, "SELECT * FROM post")

	if err == nil {
		c.JSON(200, posts)
	} else {
		c.JSON(404, gin.H{"error": err.Error()})
	}

	// curl -i http://localhost:8080/api/post
}

func GetPostsById(c *gin.Context) {
	id := c.Params.ByName("id")
	var posts []Post
	_, err := dbmap.Select(&posts, "SELECT * FROM post WHERE id=?", id)

	if err == nil {
		c.HTML(200, "post.html", posts[0])
	} else {
		c.JSON(404, gin.H{"error": "questions for this query not found"})
	}

	// curl -i http://localhost:8080/api/v1/waifus/1
}


func PostPost(c *gin.Context) {

	var post Post
	c.Bind(&post)

	log.Println(post)

	summary := ""
	truncated := false
	if len(post.Text)>255 {
		truncated = true
		summary = post.Text[0:255]
	} else {
		summary = post.Text
	}

	if post.Title != "" && post.Text != "" && post.Tags != "" && summary != ""{

		if insert, err := dbmap.Exec(`INSERT INTO post (title, text, created, truncated, tags, summary) VALUES (?, ?, ?,?,?,?)`,
		 post.Title, post.Text, time.Now().Format("2006.01.02 15:04:05"), truncated, post.Tags, summary); insert != nil {
      post_id, err := insert.LastInsertId()
			if err == nil {
				content := &Post{
          Id:        post_id,
          Title: post.Title,
          Text: post.Text,
          CreatedAt: time.Now().Format("2006.01.02 15:04"),
					Truncated: truncated,
					Tags: post.Tags,
					Summary: summary,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}

		}else {
			checkErr(err, "Insert Failed!")
			c.JSON(400, gin.H{"error": "Some Error"})
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"title\": \"Second Post\",\"tags\": \"Bootstrap;mac;PHP;Waifu Sim\", \"text\": \"\" }" http://localhost:8080/api/post/
}


func DeletePost(c *gin.Context) {
	if token, err := CheckAndDecodeToken(c.Query("token")); err != nil {
		c.JSON(403, gin.H{"error": "Invalid token!"})
  } else {
		if token.Claims["admin"] != "true"{
			c.JSON(403, gin.H{"error": "You're not admin!"})
		} else {
	id := c.Params.ByName("id")

	var post Post
	err := dbmap.SelectOne(&post, "SELECT * FROM post WHERE id=?", id)

	if err == nil {
		_, err = dbmap.Delete(&post)

		if err == nil {
			c.JSON(200, gin.H{"id #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "post not found"})
	}
}
}
	// curl -i -X DELETE http://localhost:8080/api/v1/waifus/1
}
