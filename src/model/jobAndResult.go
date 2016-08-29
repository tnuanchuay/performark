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
}

