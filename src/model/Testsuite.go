package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Testcase struct{
		Thread		string		`json:"t"`
		Connection	string		`json:"c"`
		Duration	string		`json:"d"`
	}

	Testsuite struct{
		Name		string
		Testcase	[]Testcase
	}
)

func (t * Testsuite) AddTestcase(test Testcase)(* Testsuite){
	t.Testcase = append(t.Testcase, test)
	return t
}

func (t * Testsuite) SetName(name string)(* Testsuite){
	t.Name = name
	return t
}

func (t * Testsuite) Save(session *mgo.Session)(* Testsuite){
	c := session.DB("performark").C("testsuite")
	c.Upsert(bson.M{"name":t.Name}, t)
	return t
}

func (Testsuite) GetAll(session *mgo.Session)([]Testsuite){
	c := session.DB("performark").C("testsuite")
	var instance []Testsuite
	c.Find(bson.M{}).All(&instance)
	return instance
}

func (Testsuite) Find(session *mgo.Session, name string)(Testsuite){
	c := session.DB("performark").C("testsuite")
	var instance Testsuite
	c.Find(bson.M{"name":name}).One(&instance)
	return instance
}