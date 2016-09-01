package model

import (
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
	"os/exec"
	"fmt"
	"bufio"
	"strings"
)

type (
	Job struct{
		Name		string
		Unique		string
		IsComplete	bool
		Request		Request
		Label		[]string
		TestcaseName	string
		Error		bool
		Script		string
		Grade		rune
	}
)

func (j *Job) Complete(session *mgo.Session){
	j.IsComplete = true
	c := session.DB("performark").C("job")
	c.Upsert(bson.M{"unique":j.Unique}, j)
}

func (j *Job)ReRunWrk(session *mgo.Session)(*Job){
	return Job{}.NewInstance(session, j.TestcaseName, j.Name, j.Request)
}

func (j *Job) RunWrk(ts Testcase, label string, time string, mongoChan chan WrkResult){
	t := ts.Thread
	c := ts.Connection
	d := ts.Duration

	url := j.Request.Url
	var command *exec.Cmd

	command = exec.Command("wrk", "-t"+t, "-c"+c, "-d"+d, "-s", fmt.Sprintf("lua/%s.lua", j.Unique),url)

	fmt.Println("label", label)
	if label == "time" {
		j.Label = append(j.Label, d)
	}else{
		j.Label = append(j.Label, c)
	}

	fmt.Println(command.Args)
	cmdReader, _ := command.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	var out string
	go func() {
		for scanner.Scan() {
			out = fmt.Sprintf("%s\n%s", out, scanner.Text())
			if strings.Contains(out, "Transfer"){
				break;
			}
		}
	}()
	command.Start()
	command.Wait()
	fmt.Println(out)
	wrkResult := WrkResult{}
	wrkResult.SetData(url, out, time)

	mongoChan <- wrkResult
}

func (Job) NewInstance(session *mgo.Session, testcase, name string, request Request) *Job{
	t := time.Now().Format("20060102150405")
	j := Job{Unique:t, IsComplete:false}
	j.TestcaseName = testcase
	j.Name = name
	request.GenerateScript(t)
	j.Request = request
	j.Save(session)
	return &j
}

func (Job) GetAllJob(session *mgo.Session) []Job{
	result := []Job{}
	c := session.DB("performark").C("job")
	c.Find(bson.M{}).Sort("-unique").All(&result)
	return result
}

func (j *Job) Save(session *mgo.Session){
	c := session.DB("performark").C("job")
	c.Upsert(bson.M{"unique":j.Unique}, j)
}

func (Job) Find(session *mgo.Session, unique string)(*Job){
	c := session.DB("performark").C("job")
	j := Job{}
	c.Find(bson.M{"unique":unique}).One(&j)
	return &j
}

func (Job) Delete(session *mgo.Session, unique string){
	c := session.DB("performark").C("job")
	c.Remove(bson.M{"unique":unique})
}

func (Job) SetError(session *mgo.Session){
	c := session.DB("performark").C("job")
	j := []Job{}
	c.Find(bson.M{"iscomplete":false}).All(&j)
	for _, jj := range j {
		jj.Error = true
		c.Update(bson.M{"unique":jj.Unique}, jj)
	}
}

func (Job) FindNotLikeThis(session *mgo.Session, unique string) []Job{
	c := session.DB("performark").C("job")
	result  := []Job{}
	err := c.Find(bson.M{"unique":bson.M{"$ne":fmt.Sprintf("%s", unique)}}).All(&result)
	fmt.Println(err)
	return result
}

func (j *Job) Grading(session *mgo.Session){
	if j.Error {
		j.Grade = 'F'
		return;
	}

	var successRates []float64
	var result []WrkResult
	var testcaseThoughputs        []float64
	result = WrkResult{}.FindByUnique(session, j.Unique)

	//calc success rate
	for _, tResult := range result{
		successRates = append(successRates, float64(tResult.Requests - tResult.Non2xx3xx)/tResult.Duration)
	}

	for i, successRate := range successRates{
		testcaseThoughputs = append(testcaseThoughputs, ((result[i].RequestPerSec * 45) + (successRate * 55))/100)
	}

	for i, tp := range testcaseThoughputs {
		fmt.Println(result[i].Connection, tp)
	}

	var thoughputMultiplyWithConnection	float64
	var sumConnection	float64
	var systemThoughtput	float64

	for i, tp := range testcaseThoughputs {
		thoughputMultiplyWithConnection += (tp * float64(result[i].Connection))
		sumConnection += float64(result[i].Connection)
	}

	systemThoughtput = thoughputMultiplyWithConnection / sumConnection
	fmt.Println("SYSTEM THOUGHTPUT", systemThoughtput)
}