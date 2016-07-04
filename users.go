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

      claims := MyCustomClaims{
        login,
        pass,
        email,
        waifu,
        admin,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        },
      }
      // Create the token
      token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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

        // ?token=QGWi3yN4RZjZ7TowJs6067FcWyJfgzzbou9As05Bo1@FjxtfPzj5kwLbPJJXRifqb@23m4uTpxp7D@JSEK1MZwMbgRritUFBR4MWkaULhm2iNYHiUt1egxq13OW056GtCdo3ZN9qNTr7gyW4PMGKsSgmFGIf.k6xqtVjhTtMkRnW8dDm6zyzRsHzEQVzJkK07n8O6q3LBPQgH8zC9GXQh7IWG2s91YCce.SPWenSdUqrGIuEUyiRS85KM4R1qWUPf1Y2Iclrsh9ro8W@pf4TWs28lFKVXdsAA9pp4ZtJ3dX8gbkJgF6.6fZ.GnxAlHFLRhj0fiSukUq4yWJR26pCyjvMDOGXXwS3f8OiaDrOXi5eqU1hJXywcfQ4UfdVh4vSZSdsPAE0CgazZ@R4dRlsSjv9h.Xs.u7m.2d75S
    // curl -i -X POST -H "Content-Type: application/json" -d "{ \"login\": \"Thea\", \"password\": \"Queen\", \"email\": \"sexy@gmail.com\", \"waifu\": \"asuna\" }" http://volhgroup.tk/register
  } else {
    c.JSON(400, gin.H{"error": "Some fields are empty"})
  }
}
