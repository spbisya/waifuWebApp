  package main

  import (
  "flag"

    "time"
    "log"

  	"github.com/gin-gonic/gin"
  	_ "github.com/go-sql-driver/mysql"
    jwt "github.com/dgrijalva/jwt-go"
  )
  
  var flagKey = flag.String("key", "VelesYang", "")
  var super_secret = "3ImfMpd86d3H0VNcV4Su1DE5jIKhIX94"

  //"3ImfMpd86d3H0VNcV4Su1DE5jIKhIX94"


  func Register(c *gin.Context) {
    admin := c.Query("admin")
    if admin == "" {
      admin = "false"
    }
    var user User
    c.Bind(&user)
    login := user.Login
    pass := user.Password
    email := user.Email
    waifu := user.Waifu
    if login != "" && pass != "" && email != "" && waifu != "" {
      var key interface{}
      key, _ = loadData(*flagKey)
      // Create the token
      token := jwt.New(jwt.SigningMethodHS256)
      // Set some claims
      token.Claims["user"] = login
      token.Claims["password"] = pass
      token.Claims["email"] = email
      token.Claims["waifu"] = waifu
      token.Claims["admin"] = admin
      token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
      // Sign and get the complete encoded token as a string
      tokenString, err := token.SignedString(key)

      //c.JSON(200, tokenString)
      if err == nil {
        secretMessage := []byte(tokenString)
        labelAES := []byte(super_secret)
        ciphertext, err := encrypt(labelAES, secretMessage)
        if err != nil {
          c.JSON(400, "Token encryption error")
        }
            log.Println(ciphertext)
    				content := &Token{
              User:        login,
              Token: ciphertext,
    				}
    				c.JSON(201, content)
        } else {
          c.JSON(400, "Token creation error")
        //  checkErr(err, "Token creation error: ")
        }
    // curl -i -X POST -H "Content-Type: application/json" -d "{ \"login\": \"Thea\", \"password\": \"Queen\", \"email\": \"sexy@gmail.com\", \"waifu\": \"asuna\" }" http://localhost:8080/register
  } else {
    c.JSON(400, gin.H{"error": "Some fields are empty"})
  }
}
