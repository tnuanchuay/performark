package model

import (
	"regexp"
	"unit/time"
	"fmt"
	"strings"
	"unit/si"
	"strconv"
)

type WrkResult struct {
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
	RawOutput	[]byte
}

type Latency struct {
	Avg	float64
	Stdev	float64
	Max 	float64
}

type ReqPerSec struct{
	Avg	int
	Stdev	int
	Max	int
}

func (wrkResult *WrkResult) SetData(url string, out []byte){
	wrkResult.Url = url
	wrkResult.SetDuration(out)
	wrkResult.SetThread(out)
	wrkResult.SetConnection(out)
	wrkResult.SetRequestPerSec(out)
	wrkResult.SetRequests(out)
	wrkResult.SetTransferPerSec(out)
	wrkResult.SetRawOutput(out)
}

func (t *WrkResult) SetTransferPerSec(s []byte){
	regexpTps := regexp.MustCompile("Transfer/sec:[ ]*[0-9.]*[kMG]B")
	result := regexpTps.FindAllStringSubmatch(string(s), -1)
	if len(result) != 1{
		t.SetError("TransferPerSec")
	}else{
		textTps := string(result[0][0])
		splitedTextTps := strings.Split(textTps, " ")
		t.TransferPerSec, _ = si.SIToFloat(splitedTextTps[len(splitedTextTps) - 1])
		fmt.Println("t.TransferPerSec", t.TransferPerSec)
	}
}

func (t *WrkResult) SetRawOutput(s []byte){
	t.RawOutput = s
}

func (t *WrkResult) SetError(s string){
	t.IsError = true
	t.WhatsError = append(t.WhatsError, s)
}

func (t *WrkResult) SetRequestPerSec(s []byte){
	regexpRps := regexp.MustCompile("Requests/sec:[ ]*[0-9.]*")
	result := regexpRps.FindAllStringSubmatch(string(s), -1)
	if len(result) != 1{
		t.SetError("RequestPerSec")
	}else{
		textRps := string(result[0][0])
		splitedTextRps := strings.Split(textRps, " ")
		t.RequestPerSec, _ = strconv.ParseFloat(splitedTextRps[len(splitedTextRps) - 1], 64)
		fmt.Println("t.RequestPerSec", t.RequestPerSec)
	}
}

func (t *WrkResult) SetRequests(s []byte){
	regexpRps := regexp.MustCompile("[0-9]* requests")
	result := regexpRps.FindAllStringSubmatch(string(s), -1)

	if len(result) != 1{
		t.SetError("Requests")
	}else{
		textReq := string(result[0][0])
		splitedTestReq := strings.Split(textReq, " ")[0]
		t.Requests, _ = strconv.Atoi(splitedTestReq)
		fmt.Println("t.Requests", t.Requests)
	}
}

func (t *WrkResult) SetDuration(s []byte){
	regexpDuration := regexp.MustCompile("requests in [0-9A-Za-z.]*,")
	result := regexpDuration.FindAllStringSubmatch(string(s), -1)

	if len(result) != 1{
		t.SetError("Duration")
	}else{
		textTime := string(result[0][0])
		splitedTextTime := strings.Split(strings.Split(textTime, "requests in ")[1], ",")[0]
		t.Duration, _ = time.StringToFloat(splitedTextTime)
		fmt.Println("t.duration", t.Duration)
	}
}

func (t *WrkResult) SetThread(s []byte){
	regexpThread := regexp.MustCompile("[0-9]* threads")
	result := regexpThread.FindAllStringSubmatch(string(s), -1)

	if len(result) != 1{
		t.SetError("Thread")
	}else{
		textThread := string(result[0][0])
		splitedTextThread := strings.Split(textThread, " ")[0]
		threadNum, _ := si.SIToFloat(splitedTextThread)
		t.Thread = int(threadNum)
		fmt.Println("t.Thread", t.Thread)
	}
}

func (t *WrkResult) SetConnection(s []byte){
	regexpConnection := regexp.MustCompile("[0-9]* connections")
	result := regexpConnection.FindAllStringSubmatch(string(s), -1)

	if len(result) != 1{
		t.SetError("Connection")
	}else{
		textConnection := string(result[0][0])
		splitedTextConnection := strings.Split(textConnection, " ")[0]
		threadNum, _ := si.SIToFloat(splitedTextConnection)
		t.Connection = int(threadNum)
		fmt.Println("t.Connection", t.Connection)
	}
}