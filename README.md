# Api for Waifu Sim 0.1

##UPDATE 07.04.2016 Independenсe day
  **Created screens for creation and viewing waifus, questions, accosts and greetings. Visit volhgroup.tk to see changes**

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
  curl -i http://volhgroup.tk/api/v1/waifus
  curl -i http://volhgroup.tk/api/v1/waifus/1
  curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Malcolm\", \"lastname\": \"Merlin\", \"characters\": \"genka\" }" http://volhgroup.tk/api/v1/waifus
  curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\", \"characters\": \"yandere\" }" http://volhgroup.tk/api/v1/waifus/1
  curl -i -X DELETE http://volhgroup.tk/api/v1/waifus/1
  ```
#### Get parts of the sentences
  ```
  curl -i http://volhgroup.tk/words?v=greetings
  curl -i http://volhgroup.tk/words?v=questions
  curl -i http://volhgroup.tk/words?v=accosts
  ```
**Greetings, Accosts and Questions Apis allow you to control three same models:**

  Model
  
    - Id
    
    - Characters
    
    - Texts
  ```
  curl -i http://volhgroup.tk/api/greeting
  curl -i http://volhgroup.tk/api/greeting/yandere
  curl -i -X POST -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"Shalom \" }" http://volhgroup.tk/api/greeting/
  curl -i -X PUT -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"hello\" }" http://volhgroup.tk/api/greeting/1
  curl -i -X DELETE http://volhgroup.tk/api/v1/greeting/1

  curl -i http://volhgroup.tk/api/accost
  curl -i http://volhgroup.tk/api/accost/yandere
  curl -i -X POST -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"baka\" }" http://volhgroup.tk/api/accost/
  curl -i -X PUT -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"mother\" }" http://volhgroup.tk/api/accost/1
  curl -i -X DELETE http://volhgroup.tk/api/v1/accost/1

  curl -i http://volhgroup.tk/api/question
  curl -i http://volhgroup.tk/api/question/yandere
  curl -i -X POST -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"who R u?\" }" http://volhgroup.tk/api/question/
  curl -i -X PUT -H "Content-Type: application/json" -d "{ \"characters\": \"yandere\", \"texts\": \"how R U?\" }" http://volhgroup.tk/api/question/1
  curl -i -X DELETE http://volhgroup.tk/api/v1/question/1
  ```
  **Finally, you can create the sentence using 1 part from each model.**
  ```
  curl -i http://volhgroup.tk/random
  ```
