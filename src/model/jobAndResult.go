package model

import (
	"gopkg.in/mgo.v2"
	"github.com/iris-contrib/errors"
)

type (
	JobAndResult struct {
		Job	Job
		Result	[]WrkResult
	}

	Compare struct{
		Job1	JobAndResult
		Job2	JobAndResult
		RequestPerSec1	[]float64
		RequestPerSec2	[]float64
		SocketError1	[]int
		SocketError2	[]int
		Success1	[]int
		Success2	[]int
		Label		[]string
	}
)

func (JobAndResult) Find(session *mgo.Session, unique string) JobAndResult{
	job := Job{}.Find(session, unique)
	result := WrkResult{}.FindByUnique(session, unique)
	jobAndResult := JobAndResult{Job:*job, Result:result}
	return jobAndResult
}

func (Compare) New(session *mgo.Session, uniq1, uniq2 string)(*Compare, error){
	compare := Compare{}
	compare.Job1 = JobAndResult{}.Find(session, uniq1)
	compare.Job2 = JobAndResult{}.Find(session, uniq2)

	//if not same testcase, we can't compare
	if compare.Job2.Job.TestcaseName != compare.Job1.Job.TestcaseName{
		return nil, errors.New("both job not same testcase")
	}

	compare.generateGraph()
	return &compare, nil
}

func (c *Compare) generateGraph(){
	//generate label
	c.Label = c.Job1.Job.Label

	//generate req/sec
	for i, result := range c.Job1.Result{
		c.RequestPerSec1 = append(c.RequestPerSec1, result.RequestPerSec)
		c.RequestPerSec2 = append(c.RequestPerSec2, c.Job2.Result[i].RequestPerSec)
	}

	//generate socket-error
	for i, result := range c.Job1.Result{
		error1 := result.SocketErrors
		error2 := c.Job2.Result[i].SocketErrors
		sumError1 := error1.Connect + error1.Read + error1.Timeout + error1.Write
		sumError2 := error2.Connect + error2.Read + error2.Timeout + error2.Write
		c.SocketError1 = append(c.SocketError1, sumError1)
		c.SocketError2 = append(c.SocketError2, sumError2)
	}

	//generate success response
	for i, result := range c.Job1.Result{
		success1 := result.Requests - result.Non2xx3xx
		success2 := c.Job2.Result[i].Requests - c.Job2.Result[i].Non2xx3xx
		c.Success1 = append(c.Success1, success1)
		c.Success2 = append(c.Success2, success2)
	}
}

