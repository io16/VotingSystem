# VotingSystem
Курсовий проект
request example  /savevote
 {  "name":"test",
    "category" : "test",
    "questions" : {
    	"question" :["Кількість студентів в чдту","Кількість жителів в Черкасах"],
    	"type" : ["radio","checkbox"]
  	},
    "answertoquestion":[[" 100","1000","1000","не знаю"], ["100 000","200 000","300 000"]]
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

/restricted/saveuservote
{
"vote":[{"answer":["dasdas"] } ],
"voteid":1
}

/getuserstovote
{
"idvote" : 10
}
8///////////////
{
"vote":[{"answer":[1,2,3,4 ], countAnswers :0, stats:[] } ],
"voteid":1
}

{
  "vote":[{"answer":["1","1","0","1"] },{"answer":["1","0","0"]} ],
"voteid":21
}

/