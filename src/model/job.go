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
	Unique		string
	IsComplete	bool
	Url		string
	Load		string
	Label		[]string
	TestcaseName	string
	Error		bool
}

func (j *Job) Complete(session *mgo.Session){
	j.IsComplete = true
	c := session.DB("performark").C("job")
	c.Upsert(bson.M{"unique":j.Unique}, j)
}

func (j *Job) RunWrk(t, c, d, time string, mongoChan chan WrkResult){
	url := j.Url
	var command *exec.Cmd

	if len(j.Load) > 0{
		command = exec.Command("wrk", "-t"+t, "-c"+c, "-d"+d, "-s", fmt.Sprintf("lua/%s.lua", j.Unique),url)
	}else{
		command = exec.Command("wrk", "-t"+t, "-c"+c, "-d"+d, url)
	}

	j.Label = append(j.Label, c)
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

func (j *Job) SetLoad(load string) *Job{
	j.Load = load
	dat, _ := ioutil.ReadFile("lua/default.lua")
	luaDefaultScript := string(dat)
	luaCustomScript := strings.Replace(luaDefaultScript, "{{load}}", fmt.Sprintf(`"%s"`, load), -1)
	ioutil.WriteFile(fmt.Sprintf("lua/%s.lua", j.Unique), []byte(luaCustomScript), 0644)
	return j
}

func (Job) NewInstance(url string, session *mgo.Session, load string, testcase string) *Job{
	t := time.Now().Format("20060102150405")
	j := Job{Unique:t, IsComplete:false}
	j.Url = url
	j.TestcaseName = testcase
	j.Save(session)
	if len(load) != 0{
		j.SetLoad(load)
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