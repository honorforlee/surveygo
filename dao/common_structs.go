package dao

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Alias    string `json:"alias"`
}

type Session struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
type Survey struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
	Questions Ques
}
type Ques struct {
	Question string `json:"question"`
	Options []string `json:"Options"`
}

func AuthenticateUser(cred UserCredentials) string {
	session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("user")
	query := clctn.Find(bson.M{"username": cred.Username, "password": cred.Password})
	err := query.One(&response)

	uuidStr := uuid.Must(uuid.NewV4()).String()
	sessionStruct := Session{cred.Username, uuidStr}
	if err != nil {
		return ""
	}

	sessionClctn := session.DB("simplesurveys").C("session")
	sessionClctn.Insert(sessionStruct)
	return uuidStr
}
func RegisterUser(cred UserCredentials) {
	session := MgoSession.Clone()
	defer session.Close()

//	var response interface{}
	clctn := session.DB("simplesurveys").C("user")
	if err := clctn.Insert(cred); err != nil {
		panic(err)
    }
  // fmt.Fprintln("sucess")  
}

func GetSessionDetails(token string) UserCredentials {
	session := MgoSession.Clone()
	defer session.Close()

	var response Session
	sessionClctn := session.DB("simplesurveys").C("session")
	query := sessionClctn.Find(bson.M{"token": token})
	err := query.One(&response)
	if err != nil {
		return UserCredentials{}
	}

	var cred UserCredentials
	clctn := session.DB("simplesurveys").C("user")
	query = clctn.Find(bson.M{"username": response.Username})
	err = query.One(&cred)
	return cred
}
func AddServey(ser Survey) {
	session := MgoSession.Clone()
	defer session.Close()

//	var response interface{}
	clctn := session.DB("simplesurveys").C("survey")
	if err := clctn.Insert(ser); err != nil {
		panic(err) 
	}
}
func GetActiveServey() ([]Survey, error){
	var survey []Survey
	session := MgoSession.Clone()
	defer session.Close()
	clctn := session.DB("simplesurveys").C("survey")
	err := clctn.Find(bson.M {}).All(&survey)
	if err != nil {
		panic(err)
	}
}