package model

import (
	"regexp"
	"unit/mtime"
	"fmt"
	"strings"
	"unit/si"
	"strconv"
	"gopkg.in/mgo.v2"
)

type WrkResult struct {
	Unique 		string
	IsError		bool
	WhatsError	[]string
	Url		string
	Duration	float64
	Thread		int
	Connection	int
	Latency		Latency
	ReqPerSec	ReqPerSec
	Requests	int
	RequestPerSec	float64
	TransferPerSec	float64
	RawOutput	string
}

type Latency struct {
	Avg	float64
	Stdev	float64
	Max 	float64
}

type ReqPerSec struct{
	Avg	float64
	Stdev	float64
	Max	float64
}

func (wrkResult *WrkResult) SetData(url, out, time string){
	wrkResult.Url = url
	wrkResult.Unique = time
	wrkResult.SetDuration(out)
	wrkResult.SetThread(out)
	wrkResult.SetConnection(out)
	wrkResult.SetRequestPerSec(out)
	wrkResult.SetRequests(out)
	wrkResult.SetTransferPerSec(out)
	wrkResult.SetRawOutput(out)
	wrkResult.SetLatency(out)
	wrkResult.SetReqPerSec(out)
}

func (t *WrkResult) SetReqPerSec(s string){
	reqexpReqPerSec := regexp.MustCompile("Req/Sec[ ]*[0-9A-Za-z.]*[ ]*[0-9A-Za-z.]*[ ]*[0-9A-Za-z.]*")
	result := reqexpReqPerSec.FindAllStringSubmatch(s, -1)
	if len(result) != 1{
		t.SetError("Req/Sec")
	}else{
		textReqPerSec := result[0][0]
		sqlitedTextReqPerSec := strings.Fields(textReqPerSec)
		reqPerSec := ReqPerSec{}
		reqPerSec.Avg, _ = si.SIToFloat(sqlitedTextReqPerSec[1])
		reqPerSec.Stdev, _ = si.SIToFloat(sqlitedTextReqPerSec[2])
		reqPerSec.Max, _ = si.SIToFloat(sqlitedTextReqPerSec[3])
		t.ReqPerSec = reqPerSec
		fmt.Println("t.Req/Sec", t.ReqPerSec)
	}
}

func (t *WrkResult) SetLatency(s string){
	regexpLatency := regexp.MustCompile("Latency[ ]*[0-9A-Za-z.]*[ ]*[0-9A-Za-z.]*[ ]*[0-9A-Za-z.]*")
	result := regexpLatency.FindAllStringSubmatch(s, -1)
	if len(result) != 1{
		t.SetError("Latency")
	}else{
		textLatency := result[0][0]
		splitedTextLatency := strings.Fields(textLatency)
		latency := Latency{}
		latency.Avg, _ = mtime.StringToFloat(splitedTextLatency[1])
		latency.Stdev, _ = mtime.StringToFloat(splitedTextLatency[2])
		latency.Max, _ = mtime.StringToFloat(splitedTextLatency[3])
		t.Latency = latency
		fmt.Println("t.Latency", t.Latency)
	}
}

func (t *WrkResult) SetTransferPerSec(s string){
	regexpTps := regexp.MustCompile("Transfer/sec:[ ]*[0-9.]*[kMG]B")
	result := regexpTps.FindAllStringSubmatch(s, -1)
	if len(result) != 1{
		t.SetError("TransferPerSec")
	}else{
		textTps := result[0][0]
		splitedTextTps := strings.Fields(textTps)
		t.TransferPerSec, _ = si.SIToFloat(splitedTextTps[len(splitedTextTps) - 1])
		fmt.Println("t.TransferPerSec", t.TransferPerSec)
	}
}

func (t *WrkResult) SetRawOutput(s string){
	t.RawOutput = s
}

func (t *WrkResult) SetError(s string){
	t.IsError = true
	t.WhatsError = append(t.WhatsError, s)
}

func (t *WrkResult) SetRequestPerSec(s string){
	regexpRps := regexp.MustCompile("Requests/sec:[ ]*[0-9.]*")
	result := regexpRps.FindAllStringSubmatch(s, -1)
	if len(result) != 1{
		t.SetError("RequestPerSec")
	}else{
		textRps := result[0][0]
		splitedTextRps := strings.Fields(textRps)
		t.RequestPerSec, _ = strconv.ParseFloat(splitedTextRps[len(splitedTextRps) - 1], 64)
		fmt.Println("t.RequestPerSec", t.RequestPerSec)
	}
}

func (t *WrkResult) SetRequests(s string){
	regexpRps := regexp.MustCompile("[0-9]* requests")
	result := regexpRps.FindAllStringSubmatch(s, -1)

	if len(result) != 1{
		t.SetError("Requests")
	}else{
		textReq := result[0][0]
		splitedTestReq := strings.Fields(textReq)[0]
		t.Requests, _ = strconv.Atoi(splitedTestReq)
		fmt.Println("t.Requests", t.Requests)
	}
}

func (t *WrkResult) SetDuration(s string){
	regexpDuration := regexp.MustCompile("requests in [0-9A-Za-z.]*,")
	result := regexpDuration.FindAllStringSubmatch(s, -1)

	if len(result) != 1{
		t.SetError("Duration")
	}else{
		textTime := result[0][0]
		textTime = strings.Replace(textTime, ",", "", -1)
		splitedTextTime := strings.Fields(textTime)[2]
		t.Duration, _ = mtime.StringToFloat(splitedTextTime)
		fmt.Println("t.duration", t.Duration)
	}
}

func (t *WrkResult) SetThread(s string){
	regexpThread := regexp.MustCompile("[0-9]* threads")
	result := regexpThread.FindAllStringSubmatch(string(s), -1)

	if len(result) != 1{
		t.SetError("Thread")
	}else{
		textThread := result[0][0]
		splitedTextThread := strings.Fields(textThread)[0]
		threadNum, _ := si.SIToFloat(splitedTextThread)
		t.Thread = int(threadNum)
		fmt.Println("t.Thread", t.Thread)
	}
}

func (t *WrkResult) SetConnection(s string){
	regexpConnection := regexp.MustCompile("[0-9]* connections")
	result := regexpConnection.FindAllStringSubmatch(s, -1)

	if len(result) != 1{
		t.SetError("Connection")
	}else{
		textConnection := result[0][0]
		splitedTextConnection := strings.Fields(textConnection)[0]
		threadNum, _ := si.SIToFloat(splitedTextConnection)
		t.Connection = int(threadNum)
		fmt.Println("t.Connection", t.Connection)
	}
}

func (t *WrkResult) Save(session *mgo.Session){
	c := session.DB("performark").C("mark")
	c.Insert(t)
}