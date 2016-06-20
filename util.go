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

"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	  jwt "github.com/dgrijalva/jwt-go"
)

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:98muzunu@/api")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Waifu{}, "Waifu").SetKeys(true, "Id")
  dbmap.AddTableWithName(Greeting{}, "Greeting").SetKeys(true, "Id")
  dbmap.AddTableWithName(Accost{}, "Accost").SetKeys(true, "Id")
  dbmap.AddTableWithName(Question{}, "Question").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

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
