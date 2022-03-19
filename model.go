package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

/*var Server_Host = os.Getenv("hostAdr")
var Server_Port = os.Getenv("PORT")*/

var Server_Host = "localhost"
var Server_Port = "8080"
var R *gin.Engine
var SecureTag = os.Getenv("IS_SECURE")
var secure bool

type Info struct {
	Fullname string `json:"fullname"`
	Title    string `json:"title"`
	Profile  string `json:"profile"`
	Skills   []struct {
		Header string   `json:"header"`
		Set    []string `json:"set"`
	} `json:"skills"`
	References []struct {
		Fullname string `json:"fullname"`
		Job      string `json:"job"`
		Company  string `json:"company,omitempty"`
		Mobile   string `json:"mobile,omitempty"`
		Email    string `json:"email,omitempty"`
	} `json:"references"`
	Contact struct {
		Mobile   string `json:"mobile"`
		Email    string `json:"email"`
		Location string `json:"location"`
		Websites []struct {
			Website     string `json:"website,omitempty"`
			URL         string `json:"url,omitempty"`
			Alternative string `json:"alternative,omitempty"`
		} `json:"websites,omitempty"`
	} `json:"contact"`
	Career []struct {
		Header  string   `json:"header"`
		Company string   `json:"company"`
		FromTo  string   `json:"from-to"`
		Tasks   []string `json:"tasks"`
	} `json:"career"`
	Education []struct {
		Header string `json:"header"`
		FromTo string `json:"from-to"`
		Detail string `json:"detail"`
	} `json:"education"`
}

/*
 JSON example
{
   "fullname":"Furkan KARACA",
   "title":"Go Developer",
   "profile":"he loves golang",
   "skills":[
      {
         "header":"Coding Languages",
         "set":[
            "Go",
            "Python",
            "Javascript"
         ]
      },
      {
         "header":"Frameworks/Systems/Databases",
         "set":[
            "Linux",
            "PostgreSQL",
            "RabbitMQ",
            "Docker",
            "REST",
            "MongoDB"
         ]
      }
   ],
   "references":[
      {
         "fullname":"Ali Veli",
         "job":"Senior Go Developer",
         "company":"Google",
         "mobile":"+905055055252",
         "email":"aliveli@google.com"
      },
      {
         "fullname":"John Doe",
         "job":"Senior Go Developer",
         "company":"Google",
         "mobile":"+905055055252",
         "email":"johndoe@smthng.com"
      }
   ],
   "contact":{
      "mobile":"050520602545",
      "email":"fukaraca@hotmail.com",
      "location":"Istanbul",
      "websites":[
         {
            "website":"Github",
            "url":"https://github.com/fukaraca",
            "alternative":"@fukaraca"
         },
         {
            "website":"Portfolio",
            "url":"https://www.portfolyomolsaydi.com",
            "alternative":null
         }
      ]
   },
   "career":[
      {
         "header":"Testing Expert",
         "company":"TSE",
         "from-to":"2015-2022",
         "tasks":[
            "Testing and inspection",
            "Reporting",
            "Maintaining quality requirements"
         ]
      }
   ],
   "education":[
      {
         "header":"Malatya Fen Lisesi",
         "from-to":"2004-2008",
         "detail":"High school"
      },
      {
         "header":"Kocaeli University",
         "from-to":"2008-2013",
         "detail":"C, C#, Vbasic, Matlab and microprocessor programming lessons in context with software development."
      },
      {
         "header":"Hackerrank",
         "from-to":"2019-2022",
         "detail":"Problem solving achievements with Go, Python and Javascript. (@fukaraca)"
      }
   ]
}

*/
