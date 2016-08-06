package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Chart struct{
	Unique				string
	dataSet				[]WrkResult
	RequestPerSec			[]float64
	TransferPerSec			[]float64
	LatencyMax			[]float64
	LatencyAvg			[]float64
	LatencyStd			[]float64
	ThreadMax			[]float64
	ThreadAvg			[]float64
	ThreadStd			[]float64
	Requests			[]int
	TotalTransfer			[]float64
	SocketErrorsTotal		[]int
	SocketErrorsConnect		[]int
	SocketErrorsRead		[]int
	SocketErrorsWrite		[]int
	SocketErrorsTimeOut		[]int
	SocketErrorsNon2xx3xx		[]int
}


func (Chart) NewInstance(unique string) *Chart{
	return &Chart{Unique:unique}
}

func (c *Chart) RetrieveSocketError(session *mgo.Session)(*Chart){
	unique := c.Unique
	data := []WrkResult{}
	if c.dataSet == nil{
		col := session.DB("performark").C("mark")
		col.Find(bson.M{"unique":unique}).All(&data)
		c.dataSet = data
	}else{
		data = c.dataSet
	}

	for _, one := range data{
		c.SocketErrorsConnect = append(c.SocketErrorsConnect, one.SocketErrors.Connect)
		c.SocketErrorsRead = append(c.SocketErrorsRead, one.SocketErrors.Read)
		c.SocketErrorsWrite = append(c.SocketErrorsWrite, one.SocketErrors.Write)
		c.SocketErrorsTimeOut = append(c.SocketErrorsTimeOut, one.SocketErrors.Timeout)
		c.SocketErrorsNon2xx3xx = append(c.SocketErrorsNon2xx3xx, one.SocketErrors.Non2xx3xx)
		errr := one.SocketErrors
		sumerr := (errr.Connect + errr.Read + errr.Timeout + errr.Write + errr.Non2xx3xx)
		c.SocketErrorsTotal = append(c.SocketErrorsTotal, sumerr)
	}

	return c
}

func (c *Chart) RetrieveTransfer(session *mgo.Session)(*Chart){
	unique := c.Unique
	data := []WrkResult{}
	if c.dataSet == nil{
		col := session.DB("performark").C("mark")
		col.Find(bson.M{"unique":unique}).All(&data)
		c.dataSet = data
	}else{
		data = c.dataSet
	}

	for _, one := range data{
		c.TotalTransfer = append(c.TotalTransfer, one.TotalTransfer)
	}

	return c
}

func (c *Chart) RetrieveRequestPerSec(session *mgo.Session)(*Chart){
	unique := c.Unique
	data := []WrkResult{}
	if c.dataSet == nil{
		col := session.DB("performark").C("mark")
		col.Find(bson.M{"unique":unique}).All(&data)
		c.dataSet = data
	}else{
		data = c.dataSet
	}

	for _, one := range data{
		c.RequestPerSec = append(c.RequestPerSec, one.RequestPerSec)
	}

	return c
}

func (c *Chart) RetrieveTransferPerSec(session *mgo.Session)(*Chart){
	unique := c.Unique
	data := []WrkResult{}
	if c.dataSet == nil{
		col := session.DB("performark").C("mark")
		col.Find(bson.M{"unique":unique}).All(&data)
		c.dataSet = data
	}else{
		data = c.dataSet
	}

	for _, one := range data{
		c.TransferPerSec = append(c.TransferPerSec, one.TransferPerSec)
	}

	return c
}

func (c *Chart) RetrieveLatency(session *mgo.Session)(*Chart){
	unique := c.Unique
	data := []WrkResult{}
	if c.dataSet == nil{
		col := session.DB("performark").C("mark")
		col.Find(bson.M{"unique":unique}).All(&data)
		c.dataSet = data
	}else{
		data = c.dataSet
	}

	for _, one := range data{
		c.LatencyMax = append(c.LatencyMax, one.Latency.Max)
		c.LatencyStd = append(c.LatencyStd, one.Latency.Stdev)
		c.LatencyAvg = append(c.LatencyAvg, one.Latency.Avg)
	}

	return c
}

func (c *Chart) RetrieveThread(session *mgo.Session)(*Chart){
	unique := c.Unique
	data := []WrkResult{}
	if c.dataSet == nil{
		col := session.DB("performark").C("mark")
		col.Find(bson.M{"unique":unique}).All(&data)
		c.dataSet = data
	}else{
		data = c.dataSet
	}

	for _, one := range data{
		c.ThreadMax = append(c.ThreadMax, one.ReqPerSec.Max)
		c.ThreadAvg = append(c.ThreadAvg, one.ReqPerSec.Avg)
		c.ThreadStd = append(c.ThreadStd, one.ReqPerSec.Stdev)
	}

	return c
}

func (c *Chart) RetrieveRequest(session *mgo.Session)(*Chart){
	unique := c.Unique
	data := []WrkResult{}
	if c.dataSet == nil{
		col := session.DB("performark").C("mark")
		col.Find(bson.M{"unique":unique}).All(&data)
		c.dataSet = data
	}else{
		data = c.dataSet
	}

	for _, one := range data{
		c.Requests = append(c.Requests, one.Requests)
	}

	return c
}

