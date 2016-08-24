package model

import (
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
	"os/exec"
	"fmt"
	"bufio"
	"strings"
	"io/ioutil"
)

type Job struct{
	Name		string
	Unique		string
	IsComplete	bool
	Url		string
	Load		string
	Label		[]string
	TestcaseName	string
	Error		bool
	Script		string
}

func (j *Job) Complete(session *mgo.Session){
	j.IsComplete = true
	c := session.DB("performark").C("job")
	c.Upsert(bson.M{"unique":j.Unique}, j)
}

func (j *Job)ReRunWrk(session *mgo.Session)(*Job){
	testsuiteName := j.TestcaseName
	return Job{}.NewInstance(j.Url, session, j.Load, testsuiteName)
}

func (j *Job) RunWrk(ts Testcase, label string, time string, mongoChan chan WrkResult){
	t := ts.Thread
	c := ts.Connection
	d := ts.Duration

	url := j.Url
	var command *exec.Cmd

	if len(j.Load) > 0{
		command = exec.Command("wrk", "-t"+t, "-c"+c, "-d"+d, "-s", fmt.Sprintf("lua/%s.lua", j.Unique),url)
	}else{
		command = exec.Command("wrk", "-t"+t, "-c"+c, "-d"+d, url)
	}
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

func (j *Job) SetBody(load string) *Job{
	j.Load = load
	ioutil.WriteFile(fmt.Sprintf("lua/%s.lua", j.Unique), []byte(load), 0644)
	return j
}

func (Job) NewInstance(url string, session *mgo.Session, load, testcase, name string) *Job{
	t := time.Now().Format("20060102150405")
	j := Job{Unique:t, IsComplete:false}
	j.Url = url
	j.TestcaseName = testcase
	j.Name = name
	j.Save(session)
	if len(load) != 0{
		j.SetBody(load)
	}
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