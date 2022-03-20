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
var res = float32(50)

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
} //

var example = ` 
{
   "fullname":"Jane Doe",
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
         "fullname":"Joanna Doe",
         "job":"Senior Go Developer",
         "company":"Google",
         "mobile":"+905055055252",
         "email":"joannadoe@google.com"
      },
      {
         "fullname":"John Doe",
         "job":"Senior Java Developer",
         "company":"Google",
         "mobile":"+905055055252",
         "email":"johndoe@smthng.com"
      }
   ],
   "contact":{
      "mobile":"050520602545",
      "email":"jane.doe@hotmail.com",
      "location":"Istanbul",
      "websites":[
         {
            "website":"Github",
            "url":"https://github.com/janesgithub",
            "alternative":"@janesgithub"
         },
         {
            "website":"Portfolio",
            "url":"https://www.iwishihad.com",
            "alternative":null
         }
      ]
   },
   "career":[
      {
         "header":"Intern",
         "company":"Google",
         "from-to":"2021-2022",
         "tasks":[
            "observing",
            "coding",
            "chilling"
         ]
      }
   ],
   "education":[
      {
         "header":"Some High School",
         "from-to":"2012-2016",
         "detail":"High school stuffs"
      },
      {
         "header":"Some University",
         "from-to":"2016-2020",
         "detail":"Something spectacular, something magnificent"
      },
      {
         "header":"Hackerrank",
         "from-to":"2019-2022",
         "detail":"Problem solving achievements with Go, Python and Javascript. (@janedoe)"
      }
   ]
}
 `