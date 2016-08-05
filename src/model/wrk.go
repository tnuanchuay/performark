package model

import (
	"regexp"
	"unit/time"
	"fmt"
	"strings"
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

func (t *WrkResult) SetDuration(s []byte){
	regexpDuration := regexp.MustCompile("requests in [0-9A-Za-z.]*,")
	result := regexpDuration.FindAllStringSubmatch(string(s), -1)
	t.RawOutput = s

	if len(result) != 1{
		t.IsError = true
		t.WhatsError = append(t.WhatsError, "Duration")
	}else{
		textTime := string(result[0][0])
		splitedTextTime := strings.Split(strings.Split(textTime, "requests in ")[1], ",")[0]
		t.Duration, _ = time.StringToInt(splitedTextTime)
		fmt.Println("t.durration", t.Duration)
	}
}