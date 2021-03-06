package main

import (
  "math/rand"
   "time"
   "strings"
   "log"
   "strconv"
   "math"

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
// curl -i http://volhgroup.tk/random?v=genke
}

func GetRandomPost(c *gin.Context){
  rand.Seed(time.Now().UTC().UnixNano())
  var posts []Post
	_, err := dbmap.Select(&posts, "SELECT * FROM post")

	if err == nil {
    post := posts[rand.Intn(len(posts)-1)]
    c.HTML(200, "post.html", gin.H{
                "title": post.Title,
                "content": post.Text,
                "time": post.CreatedAt,
            })
  } else {
    c.JSON(404, gin.H{"error": "Couldn't generate post"})
  }
}

func GetAllPosts(c *gin.Context){
  query:=c.DefaultQuery("page", "1")
  page, _:= strconv.Atoi(query);
  var posts []Post
  _, err := dbmap.Select(&posts, "SELECT * FROM post")
  posts = ReverseArray(posts)
  var newPosts []PostForSite
  pagesCount:=int(math.Ceil(float64(len(posts))/4))

  switch page {
  case 1:
  if len(posts)<4 {
    newPosts = GetTags(posts[0:len(posts)])
  } else {
    newPosts = GetTags(posts[0:4])
  }
  isNext:=CheckIfHasNext(page, len(posts))
  isLast:=!isNext

  if pagesCount>5 {
    pagesCount = 5
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
      c.HTML(200, "posts.html", gin.H{
                  "Posts": newPosts,
                  "Page": 1,
                  "Next": 2,
                  "Previous": 0,
                  "IsNext": isNext,
                  "IsPrevious": false,
                  "Count": array,
                  "IsFirst": true,
                  "IsLast": isLast,
                  "Last": int(math.Ceil(float64(len(posts))/4)),
                  "More": true,
                  "Less": false,
              })
    } else {
      checkErr(err, "Couldn't get Posts")
      c.JSON(404, gin.H{"error": "Couldn't generate post"})
    }

  case 2:
  if len(posts)<4 {
    newPosts = GetTags(posts[0:len(posts)])
  } else {
    newPosts = GetTags(posts[4:8])
  }
  isNext:=CheckIfHasNext(page, len(posts))
  isLast:=!isNext
  if pagesCount>5 {
    pagesCount = 5
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
      c.HTML(200, "posts.html", gin.H{
                  "Posts": newPosts,
                  "Page": 2,
                  "Next": 3,
                  "Previous": 1,
                  "IsNext": isNext,
                  "IsPrevious": true,
                  "Count": array,
                  "IsFirst": true,
                  "IsLast": isLast,
                  "Last": int(math.Ceil(float64(len(posts))/4)),
                  "More": true,
                  "Less": false,
              })
    } else {
      checkErr(err, "Couldn't get Posts")
      c.JSON(404, gin.H{"error": "Couldn't generate post"})
    }

  case 3:

  isNext:=CheckIfHasNext(page, len(posts))
  isLast:=!isNext
  if pagesCount>5 {
    pagesCount = 5
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
      c.HTML(200, "posts.html", gin.H{
                  "Posts": newPosts,
                  "Page": 3,
                  "Next": 4,
                  "Previous": 3,
                  "IsNext": isNext,
                  "IsPrevious": true,
                  "Count": array,
                  "IsFirst": true,
                  "IsLast": isLast,
                  "Last": int(math.Ceil(float64(len(posts))/4)),
                  "More": true,
                  "Less": false,
              })
    } else {
      checkErr(err, "Couldn't get Posts")
      c.JSON(404, gin.H{"error": "Couldn't generate post"})
    }
  default:
    pagesCount=3;
      newPosts = GetTags(posts[(page-1)*4:page*4])
  isNext:=CheckIfHasNext(page, len(posts))
  isLast:=!CheckIfHasNext(page, len(posts))
  more:=CheckIfHasNext(page+1, len(posts))
  less:=CheckIfHasNext(page-1, len(posts))
      if err == nil {
        c.HTML(200, "posts.html", gin.H{
                    "Posts": newPosts,
                    "Page": page,
                    "Next": page+1,
                    "Previous": page-1,
                    "IsNext": isNext,
                    "IsPrevious": true,
                    "Count": GetArray(len(posts), page),
                    "IsFirst": false,
                    "IsLast": isLast,
                    "Last": int(math.Ceil(float64(len(posts))/4)),
                    "More": more,
                    "Less": less,
                })
      } else {
        checkErr(err, "Couldn't get Posts")
        c.JSON(404, gin.H{"error": "Couldn't generate post"})
      }
    }
}




func GetPostsForSiteById(c *gin.Context) {
	id := c.Params.ByName("id")
	var posts []Post
  var postss []Post
  _, err := dbmap.Select(&postss, "SELECT * FROM post")
  length:=int64(len(postss))
	_, err = dbmap.Select(&posts, "SELECT * FROM post WHERE id=?", id)
  right:=false
  left:=false
  rightId:=int64(0)
  leftId:=int64(0)

  switch posts[0].Id {
  case length:
    left=true
    leftId=posts[0].Id-1
  case 1:
    right=true
    rightId=posts[0].Id+1
  default:
    right=true
    rightId=posts[0].Id+1
    left=true
    leftId=posts[0].Id-1
  }



  newPost := PostForSingle{
    Post:        posts[0],
    Tags: strings.Split(posts[0].Tags,";"),
    Right: right,
    Left: left,
    RightId: rightId,
    LeftId: leftId,
  }
  log.Println(newPost)
	if err == nil {
		c.HTML(200, "post.html", newPost)
	} else {
		c.JSON(404, gin.H{"error": "questions for this query not found"})
	}

	// curl -i http://localhost:8080/api/v1/waifus/1
}

func GetTags(s []Post) []PostForSite{
   newPosts := make([]PostForSite, len(s))
    for i:=0;i< len(s);i++{
      newPosts[i] = PostForSite{
    Post:        s[i],
    Tags: strings.Split(s[i].Tags,";"),
  }
    }
    return newPosts
}
