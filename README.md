# Api for Waifu Sim 0.1

##Setup:
  Install Go, mySql, then
  ```
  go get github.com/gin-gonic/gin
  go get github.com/go-sql-driver/mysql
  go get gopkg.in/gorp.v1
  ```
  ```
  git clone https://github.com/spbisya/waifuWebApp
  ```
  Copy the project to $GOPATH/src/github.com/spbisya/waifuWebApp directory
  ```
  go install github.com/spbisya/waifuWebApp && waifuWebApp
  ```
  
##UPDATE 06.16.2016
  **All POST requests should be made with adding ?token=<your token> to path**
  
  **All DELETE requests can be only made with administrator token**
##Features

####New features will be added below each update

  Since the last update if user want to POST some data to the DB he should POST model 
  
  **User {login, password, email, waifu}**
  
  to "/register" path. Then he should remember the token and use it in future.

  **Waifu Api allows you to control waifus on the server**
  
  Waifu model:
  
    - Id
  
    - FirstName
  
    - LastName
  
    - Characters
  ```
  curl -i http://localhost:8080/api/v1/waifus
  curl -i http://localhost:8080/api/v1/waifus/1
  curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Malcolm\", \"lastname\": \"Merlin\", \"characters\": \"genka\" }" http://localhost:8080/api/v1/waifus
  curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\", \"characters\": \"yandere\" }" http://localhost:8080/api/v1/waifus/1
  curl -i -X DELETE http://localhost:8080/api/v1/waifus/1
  ```
#### Get parts of the sentences
  ```
  curl -i http://localhost:8080/words?v=greetings
  curl -i http://localhost:8080/words?v=questions
  curl -i http://localhost:8080/words?v=accosts
  ```
**Greetings, Accosts and Questions Apis allow you to control three same models:**

  Model
  
    - Id
    
    - Characters
    
    - Texts
  ```
  curl -i http://localhost:8080/api/greeting
  curl -i http://localhost:8080/api/greeting/yandere
  curl -i -X POST -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"Shalom \" }" http://localhost:8080/api/greeting/
  curl -i -X PUT -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"hello\" }" http://localhost:8080/api/greeting/1
  curl -i -X DELETE http://localhost:8080/api/v1/greeting/1

  curl -i http://localhost:8080/api/accost
  curl -i http://localhost:8080/api/accost/yandere
  curl -i -X POST -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"baka\" }" http://localhost:8080/api/accost/
  curl -i -X PUT -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"mother\" }" http://localhost:8080/api/accost/1
  curl -i -X DELETE http://localhost:8080/api/v1/accost/1

  curl -i http://localhost:8080/api/question
  curl -i http://localhost:8080/api/question/yandere
  curl -i -X POST -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"who R u?\" }" http://localhost:8080/api/question/
  curl -i -X PUT -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"how R U?\" }" http://localhost:8080/api/question/1
  curl -i -X DELETE http://localhost:8080/api/v1/question/1
  ```
  **Finally, you can create the sentence using 1 part from each model.**
  ```
  curl -i http://localhost:8080/random
  ```
