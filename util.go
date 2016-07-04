package main

import (
	"database/sql"
	"gopkg.in/gorp.v1"
	"log"
	"fmt"
		"io"
		"io/ioutil"
		"os"
		"errors"
		"crypto/aes"
		"crypto/rand"
		"crypto/cipher"
		"encoding/base64"
		"strings"
		"math"
		"strconv"
"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	  jwt "github.com/dgrijalva/jwt-go"
)

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:pass@/api")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Waifu{}, "waifu").SetKeys(true, "Id")
  dbmap.AddTableWithName(Greeting{}, "greeting").SetKeys(true, "Id")
  dbmap.AddTableWithName(Accost{}, "accost").SetKeys(true, "Id")
  dbmap.AddTableWithName(Question{}, "question").SetKeys(true, "Id")
	dbmap.AddTableWithName(Post{}, "post").SetKeys(true, "Id").ColMap("text").SetMaxSize(21845)
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	// _, err = dbmap.Exec("ALTER TABLE post DROP text;")
	// _, err = dbmap.Exec("ALTER TABLE post ADD text VARCHAR(65535) AFTER title;")
	// _, err = dbmap.Exec("ALTER TABLE post DROP tags;")
	// _, err = dbmap.Exec("ALTER TABLE post ADD tags VARCHAR(65535) AFTER truncated;")
	// _, err = dbmap.Exec("ALTER TABLE post DROP summary;")
	// _, err = dbmap.Exec("ALTER TABLE post ADD summary VARCHAR(65535) AFTER tags;")
	// checkErr(err, "Setting mode failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func loadData(p string) ([]byte, error) {
	if p == "" {
		return nil, fmt.Errorf("No path specified")
	}

	var rdr io.Reader
	if p == "-" {
		rdr = os.Stdin
	} else {
		if f, err := os.Open(p); err == nil {
			rdr = f
			defer f.Close()
		} else {
			return nil, err
		}
	}
	return ioutil.ReadAll(rdr)
}

//token as encoded string, return token, error
func CheckAndDecodeToken(tokenStr string) (*jwt.Token, error){
	if tokenStr != "" {
	tokenStr, err :=	decrypt([]byte(super_secret), tokenStr)
		if err != nil {
			return nil, err
		} else {
		var key interface{}
		key, _ = loadData(*flagKey)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
    	return key, nil
		})
		if err == nil {
			if token.Valid {
				return token, nil
			} else {
				return nil, errors.New("Invalid Token")
			}
		} else {
			return nil, err
		}
	}
	} else {
		return nil, errors.New("Empty string")
	}
}

func encrypt(key, text []byte) (string, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    b := base64.StdEncoding.EncodeToString(text)
    ciphertext := make([]byte, aes.BlockSize+len(b))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
		newText := base64.StdEncoding.EncodeToString(ciphertext)
		if strings.ContainsAny(newText, "/") {
		newText = 	strings.Replace(newText, "/", ".", -1)
		}

		if strings.ContainsAny(newText, "+") {
			newText = strings.Replace(newText, "+", "@", -1)
		}
		 newText =	strings.TrimRight(newText, "=")
		 newText = Reverse(newText)
    return newText, nil
}

func decrypt(key []byte, textS string) (string, error) {
	 textS = Reverse(textS)
	 textS += "=="
	if strings.ContainsAny(textS, ".") {
		textS = strings.Replace(textS, ".", "/", -1)
	}

	if strings.ContainsAny(textS, "@") {
		textS = strings.Replace(textS, "@", "+", -1)
	}
	text, err := base64.StdEncoding.DecodeString(textS)
	if err != nil {
			return "", err
	}
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    if len(text) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }
    iv := text[:aes.BlockSize]
    text = text[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(text, text)
    data, err := base64.StdEncoding.DecodeString(string(text))
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ReverseArray(s []Post) []Post{
	for i,j:=0, len(s)-1;i<j;i,j=i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseArrayW(s []Waifu) []Waifu{
	for i,j:=0, len(s)-1;i<j;i,j=i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseArrayQ(s []Question) []Question{
	for i,j:=0, len(s)-1;i<j;i,j=i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseArrayA(s []Accost) []Accost{
	for i,j:=0, len(s)-1;i<j;i,j=i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseArrayG(s []Greeting) []Greeting{
	for i,j:=0, len(s)-1;i<j;i,j=i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func GetType1(part string, caps string, typ int, c *gin.Context, posts []Question, page int) {
	_, err := dbmap.Select(&posts, "SELECT * FROM "+part)

		posts = ReverseArrayQ(posts)


	pagesCount:=int(math.Ceil(float64(len(posts))/30))

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
			if len(posts)<30 {
				posts1=posts[0:len(posts)]
			}	else{
				posts1 = posts[0:30]
			}
			c.HTML(200, "allParts.html", gin.H{
				"Name": caps,
				"Type": part,
									"Parts": posts1,
									"Page": 1,
									"Next": 2,
									"Previous": 0,
									"IsNext": isNext,
									"IsPrevious": false,
									"Count": array,
									"IsFirst": true,
									"IsLast": isLast,
									"Last": int(math.Ceil(float64(len(posts))/30)),
									"More": more,
									"Less": false,
							})
		} else {
			checkErr(err, "Couldn't get "+part+"s")
			c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
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
			if len(posts)<60 {
				posts1=posts[30:len(posts)]
			}	else{
				posts1 = posts[30:60]
			}
			c.HTML(200, "allParts.html", gin.H{
				"Name": caps,
				"Type": part,
									"Parts": posts1,
									"Page": 2,
									"Next": 3,
									"Previous": 1,
									"IsNext": isNext,
									"IsPrevious": true,
									"Count": array,
									"IsFirst": true,
									"IsLast": isLast,
									"Last": int(math.Ceil(float64(len(posts))/30)),
									"More": more,
									"Less": false,
							})
		} else {
			checkErr(err, "Couldn't get "+part+"s")
			c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
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
			if len(posts)<90 {
				posts1=posts[60:len(posts)]
			}	else{
				posts1 = posts[60:90]
			}
			c.HTML(200, "allParts.html", gin.H{
				"Name": caps,
				"Type": part,
									"Parts": posts1,
									"Page": 3,
									"Next": 4,
									"Previous": 3,
									"IsNext": isNext,
									"IsPrevious": true,
									"Count": array,
									"IsFirst": true,
									"IsLast": isLast,
									"Last": int(math.Ceil(float64(len(posts))/30)),
									"More": more,
									"Less": false,
							})
		} else {
			checkErr(err, "Couldn't get "+part+"s")
			c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
		}
	default:
		pagesCount=3;
	isNext:=CheckIfHasNextW(page, len(posts))
	isLast:=!CheckIfHasNextW(page, len(posts))
	more:=CheckIfHasNextW(page+1, len(posts))
	less:=CheckIfHasNextW(page-1, len(posts))
			if err == nil {
				posts1:=posts
				if len(posts)<page*30 {
					posts1=posts[(page-1)*30:len(posts)]
				}	else{
					posts1 = posts[(page-1)*30:page*30]
				}
				c.HTML(200, "allParts.html", gin.H{
					"Name": caps,
					"Type": part,
										"Parts": posts1,
										"Page": page,
										"Next": page+1,
										"Previous": page-1,
										"IsNext": isNext,
										"IsPrevious": true,
										"Count": GetArray(len(posts), page),
										"IsFirst": false,
										"IsLast": isLast,
										"Last": int(math.Ceil(float64(len(posts))/30)),
										"More": more,
										"Less": less,
								})
			} else {
				checkErr(err, "Couldn't get "+part+"s")
				c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
			}
		}
}

func GetType2(part string, caps string, typ int, c *gin.Context, posts []Accost , page int) {
	_, err := dbmap.Select(&posts, "SELECT * FROM "+part)

		posts = ReverseArrayA(posts)


	pagesCount:=int(math.Ceil(float64(len(posts))/30))

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
			if len(posts)<30 {
				posts1=posts[0:len(posts)]
			}	else{
				posts1 = posts[0:30]
			}
			c.HTML(200, "allParts.html", gin.H{
				"Name": caps,
				"Type": part,
									"Parts": posts1,
									"Page": 1,
									"Next": 2,
									"Previous": 0,
									"IsNext": isNext,
									"IsPrevious": false,
									"Count": array,
									"IsFirst": true,
									"IsLast": isLast,
									"Last": int(math.Ceil(float64(len(posts))/30)),
									"More": more,
									"Less": false,
							})
		} else {
			checkErr(err, "Couldn't get "+part+"s")
			c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
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
			if len(posts)<60 {
				posts1=posts[30:len(posts)]
			}	else{
				posts1 = posts[30:60]
			}
			c.HTML(200, "allParts.html", gin.H{
				"Name": caps,
				"Type": part,
									"Parts": posts1,
									"Page": 2,
									"Next": 3,
									"Previous": 1,
									"IsNext": isNext,
									"IsPrevious": true,
									"Count": array,
									"IsFirst": true,
									"IsLast": isLast,
									"Last": int(math.Ceil(float64(len(posts))/30)),
									"More": more,
									"Less": false,
							})
		} else {
			checkErr(err, "Couldn't get "+part+"s")
			c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
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
			if len(posts)<90 {
				posts1=posts[60:len(posts)]
			}	else{
				posts1 = posts[60:90]
			}
			c.HTML(200, "allParts.html", gin.H{
				"Name": caps,
				"Type": part,
									"Parts": posts1,
									"Page": 3,
									"Next": 4,
									"Previous": 3,
									"IsNext": isNext,
									"IsPrevious": true,
									"Count": array,
									"IsFirst": true,
									"IsLast": isLast,
									"Last": int(math.Ceil(float64(len(posts))/30)),
									"More": more,
									"Less": false,
							})
		} else {
			checkErr(err, "Couldn't get "+part+"s")
			c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
		}
	default:
		pagesCount=3;
	isNext:=CheckIfHasNextW(page, len(posts))
	isLast:=!CheckIfHasNextW(page, len(posts))
	more:=CheckIfHasNextW(page+1, len(posts))
	less:=CheckIfHasNextW(page-1, len(posts))
			if err == nil {
				posts1:=posts
				if len(posts)<page*30 {
					posts1=posts[(page-1)*30:len(posts)]
				}	else{
					posts1 = posts[(page-1)*30:page*30]
				}
				c.HTML(200, "allParts.html", gin.H{
					"Name": caps,
					"Type": part,
										"Parts": posts1,
										"Page": page,
										"Next": page+1,
										"Previous": page-1,
										"IsNext": isNext,
										"IsPrevious": true,
										"Count": GetArray(len(posts), page),
										"IsFirst": false,
										"IsLast": isLast,
										"Last": int(math.Ceil(float64(len(posts))/30)),
										"More": more,
										"Less": less,
								})
			} else {
				checkErr(err, "Couldn't get "+part+"s")
				c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
			}
		}
}

func GetType3(part string, caps string, typ int, c *gin.Context, posts []Greeting, page int) {
	_, err := dbmap.Select(&posts, "SELECT * FROM "+part)

		posts = ReverseArrayG(posts)


	pagesCount:=int(math.Ceil(float64(len(posts))/30))

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
			if len(posts)<30 {
				posts1=posts[0:len(posts)]
			}	else{
				posts1 = posts[0:30]
			}
			c.HTML(200, "allParts.html", gin.H{
				"Name": caps,
				"Type": part,
									"Parts": posts1,
									"Page": 1,
									"Next": 2,
									"Previous": 0,
									"IsNext": isNext,
									"IsPrevious": false,
									"Count": array,
									"IsFirst": true,
									"IsLast": isLast,
									"Last": int(math.Ceil(float64(len(posts))/30)),
									"More": more,
									"Less": false,
							})
		} else {
			checkErr(err, "Couldn't get "+part+"s")
			c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
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
			if len(posts)<60 {
				posts1=posts[30:len(posts)]
			}	else{
				posts1 = posts[30:60]
			}
			c.HTML(200, "allParts.html", gin.H{
				"Name": caps,
				"Type": part,
									"Parts": posts1,
									"Page": 2,
									"Next": 3,
									"Previous": 1,
									"IsNext": isNext,
									"IsPrevious": true,
									"Count": array,
									"IsFirst": true,
									"IsLast": isLast,
									"Last": int(math.Ceil(float64(len(posts))/30)),
									"More": more,
									"Less": false,
							})
		} else {
			checkErr(err, "Couldn't get "+part+"s")
			c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
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
			if len(posts)<90 {
				posts1=posts[60:len(posts)]
			}	else{
				posts1 = posts[60:90]
			}
			c.HTML(200, "allParts.html", gin.H{
				"Name": caps,
				"Type": part,
									"Parts": posts1,
									"Page": 3,
									"Next": 4,
									"Previous": 3,
									"IsNext": isNext,
									"IsPrevious": true,
									"Count": array,
									"IsFirst": true,
									"IsLast": isLast,
									"Last": int(math.Ceil(float64(len(posts))/30)),
									"More": more,
									"Less": false,
							})
		} else {
			checkErr(err, "Couldn't get "+part+"s")
			c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
		}
	default:
		pagesCount=3;
	isNext:=CheckIfHasNextW(page, len(posts))
	isLast:=!CheckIfHasNextW(page, len(posts))
	more:=CheckIfHasNextW(page+1, len(posts))
	less:=CheckIfHasNextW(page-1, len(posts))
			if err == nil {
				posts1:=posts
				if len(posts)<page*30 {
					posts1=posts[(page-1)*30:len(posts)]
				}	else{
					posts1 = posts[(page-1)*30:page*30]
				}
				c.HTML(200, "allParts.html", gin.H{
					"Name": caps,
					"Type": part,
										"Parts": posts1,
										"Page": page,
										"Next": page+1,
										"Previous": page-1,
										"IsNext": isNext,
										"IsPrevious": true,
										"Count": GetArray(len(posts), page),
										"IsFirst": false,
										"IsLast": isLast,
										"Last": int(math.Ceil(float64(len(posts))/30)),
										"More": more,
										"Less": less,
								})
			} else {
				checkErr(err, "Couldn't get "+part+"s")
				c.JSON(404, gin.H{"error": "Couldn't get "+part+"s"})
			}
		}
}

func GetAllPartsHTML(part string, caps string, typ int, c *gin.Context) {
	query:=c.DefaultQuery("page", "1")
	page, _:= strconv.Atoi(query);

	switch typ {
	case 1:
		var posts []Question
		GetType1(part, caps, typ, c, posts, page)
	case 2:
		var posts []Accost
			GetType2(part, caps, typ, c, posts, page)
	case 3:
		var posts []Greeting
			GetType3(part, caps, typ, c, posts, page)
	}

	// curl -i http://localhost:8080/api/v1/waifus/1
}

func GetPartsNew(part string, caps string, typ int, c *gin.Context) {
	question := c.PostForm(part)
	characters := c.PostForm("character")
	if question != "" && characters != ""{
		if insert, _ := dbmap.Exec(`INSERT INTO `+part+` (characters, texts) VALUES (?, ?)`, characters, question ); insert != nil {
      _, err := insert.LastInsertId()
			if err == nil {
				GetAllPartsHTML(part, caps, typ, c)
			} else {
				c.JSON(201, gin.H{"error": err.Error()})
			}
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

	// curl -i http://localhost:8080/api/v1/waifus/1
}


func GetArray(c int, curr int) []Page{
  if int(math.Ceil(float64(c)/4))>5 {
  array:=make([]Page, 5)
  count:=int(math.Ceil(float64(c)/4))
  log.Println(count, " ", curr)
  if curr>3 {
    if curr>count-2{
    if curr==count-2{
      k:=0
      for i:=curr-2;i<curr+3;i++ {
        current:=false
        if i == curr{
          current = true
        }
        array[k] = Page{
          Number: i,
          Current: current,
        }
        k++
      }
    } else if curr==count-1{
      k:=0
      for i:=curr-3;i<curr+2;i++ {
        current:=false
        if i == curr{
          current = true
        }
        array[k] = Page{
          Number: i,
          Current: current,
        }
        k++
      }
    } else if curr==count{
      k:=0
      for i:=curr-4;i<curr+1;i++ {
        current:=false
        if i == curr{
          current = true
        }
        array[k] = Page{
          Number: i,
          Current: current,
        }
        k++
      }
    }
    } else {
    k:=0
    for i:=curr-2;i<curr+3;i++ {
      current:=false
      if i == curr{
        current = true
      }
      array[k] = Page{
        Number: i,
        Current: current,
      }
      k++
    }
  }
  }
  return array
  } else {
  array:=make([]Page, int(math.Ceil(float64(c)/4)))
  for i:=1;i<int(math.Ceil(float64(c)/4))+1;i++ {
    current:=false
    if i == curr{
      current = true
    }
    array[i-1] = Page{
      Number: i,
      Current: current,
    }
  }
  return array
  }
}

func CheckIfHasNext(c int, l int) bool{
  if c<int(math.Ceil(float64(l)/4)) {
    return true
  }
  return false
}

func CheckIfHasNextW(c int, l int) bool{
  if c<int(math.Ceil(float64(l)/30)) {
    return true
  }
  return false
}
