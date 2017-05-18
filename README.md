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
/getjwt
{
    "login":"string",
 	 "pass":"123456"
}

/restricted/adduservote
{
"vote":[{"answer":["dasdas" ] } ],
"voteid":1
}


8///////////////
{
"vote":[{"answer":[1,2,3,4 ], countAnswers :0, 1:0, 2:0,3:0,4:0 } ],
"voteid":1
}

