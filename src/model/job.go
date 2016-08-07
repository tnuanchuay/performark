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

type Job struct{
	Unique		string
	IsComplete	bool
	Url		string
}

func (j *Job) Complete(session *mgo.Session){
	j.IsComplete = true
	c := session.DB("performark").C("job")
	c.Upsert(bson.M{"unique":j.Unique}, j)
}

func (j *Job) RunWrk(t, c, d, time string, mongoChan chan WrkResult){
	url := j.Url

	command := exec.Command("wrk", "-t"+t, "-c"+c, "-d"+d, url)
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

func (Job) NewInstance(url string, session *mgo.Session) *Job{
	t := time.Now().Format("20060102150405")
	j := Job{Unique:t, IsComplete:false}
	j.Url = url
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
	c.Insert(j)
}

func (Job) Find(session *mgo.Session, unique string)(*Job){
	c := session.DB("performark").C("job")
	j := Job{}
	c.Find(bson.M{"unique":unique}).One(&j)
	return &j
}