# VotingSystem
Курсовий проект
request example  /savevote
 {  "name":"test",
    "category" : "test",
    "questions" : {
    	"question" :["Кількість студентів в чдту","Кількість жителів в Черкасах"]
  	},
    "answertoquestion":{"1" :[" 100","1000","1000","не знаю"], "2":["100 000","200 000","300 000"]
     }


 }

 request example /getvote
 { "idvote" : 1 }

request example  /adduser
{
 "name" :"Alex",
 "login":  "string" ,
 "email":"lalla@ukr.net" ,
 "pass":"123456"


 }

