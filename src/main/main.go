package main

import (
	"gopkg.in/mgo.v2"
	"github.com/kataras/iris"
	"log"
	"time"
	"model"
	"strings"
	"github.com/googollee/go-socket.io"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"sync"
	"github.com/PuerkitoBio/goquery"
)

var minimalTestSuite	model.Testsuite
var c1kToC10kTestSuit	model.Testsuite

func initMinimalTestSuite(session *mgo.Session){
	minimalTestSuite.SetName("minimal").
		AddTestcase(model.Testcase{Thread:"1", Connection:"1", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"10", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"100", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"1k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"10k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"100k", Duration:"10s"}).
		Save(session)
}

func initC1kToC10kTestSuit(session *mgo.Session){
	c1kToC10kTestSuit.SetName("c1k-c10k").
		AddTestcase(model.Testcase{Thread:"4", Connection:"1k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"2k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"3k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"4k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"5k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"6k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"7k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"8k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"9k", Duration:"10s"}).
		AddTestcase(model.Testcase{Thread:"4", Connection:"10k", Duration:"10s"}).
		Save(session)
}

func initMongoSession() *mgo.Session{
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	return session
}

func initBasicTestcase(session *mgo.Session){
	initMinimalTestSuite(session)
	initC1kToC10kTestSuit(session)
}

func createBasicChannel()(chan *model.Job, chan model.WrkResult){
	return make(chan *model.Job, 100), make(chan model.WrkResult, 100)
}

func main(){
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	initBasicTestcase(session)
	modelchan, mongochan := createBasicChannel()
	model.Job{}.SetError(session)

	dat, _ := ioutil.ReadFile("templates/script.js")
	CHART_SCRIPT := string(dat)

	iris.Config.IsDevelopment = true

	iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	})

	iris.Static("/assets", "./static/assets", 1)
	iris.Static("/images", "./static/images", 1)

	iris.Get("/", func(ctx *iris.Context){
		testCase := model.Testsuite{}.GetAll(session)
		jsonTestcase, _ := json.Marshal(testCase)
		ctx.Render("index.html", map[string]interface{}{"testcase":string(jsonTestcase)})
	})

	iris.Get("/job/:unique", func(ctx *iris.Context){
		unique := ctx.Param("unique")
		j := model.Job{}.Find(session, unique)
		ctx.Render("job.html", map[string]interface{}{
			"Unique":unique,
			"Name":j.Name,
			"Url":j.Request.Url,
			"Load":j.Request.Load,
			"TestCaseName":j.TestcaseName,
		})
	})

	iris.Get("/delete/:unique", func(ctx *iris.Context){
		unique := ctx.Param("unique")
		model.Job{}.Delete(session, unique)
		model.WrkResult{}.Delete(session, unique)
		ctx.Redirect("/", iris.StatusOK)
	})

	iris.Get("/script/wrk-stats/:unique", func(ctx *iris.Context) {
		unique := ctx.Param("unique")
		j := model.Job{}.Find(session, unique)
		chart := model.Chart{}.NewInstance(unique)

		chart.RetrieveRequestPerSec(session).
			RetrieveTransferPerSec(session).
			RetrieveLatency(session).
			RetrieveThread(session).
			RetrieveRequest(session).
			RetrieveTransfer(session).
			RetrieveSocketError(session).
			RetrieveNon2xx3xx(session)

		jsonrps, err := json.Marshal(chart.RequestPerSec)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsontps, err := json.Marshal(chart.TransferPerSec)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsonlm, err := json.Marshal(chart.LatencyMax)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsonla, err := json.Marshal(chart.LatencyAvg)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsonls, err := json.Marshal(chart.LatencyStd)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsontm, err := json.Marshal(chart.ThreadMax)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsonta, err := json.Marshal(chart.ThreadAvg)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsonts, err := json.Marshal(chart.ThreadStd)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsonr, err := json.Marshal(chart.Requests)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsontt, err := json.Marshal(chart.TotalTransfer)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsonec, err := json.Marshal(chart.SocketErrorsConnect)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsoner, err := json.Marshal(chart.SocketErrorsRead)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsonew, err := json.Marshal(chart.SocketErrorsWrite)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsonet, err := json.Marshal(chart.SocketErrorsTimeOut)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		jsone, err := json.Marshal(chart.SocketErrorsTotal)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		rp2xx, err := json.Marshal(chart.Non2xx3xx)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		label, err := json.Marshal(j.Label)
		if err != nil{
			ctx.JSON(iris.StatusOK, err)
		}

		s := CHART_SCRIPT
		s = strings.Replace(s, "{{.Unique}}", unique, -1)
		s = strings.Replace(s, "{{.rps}}", string(jsonrps), -1)
		s = strings.Replace(s, "{{.tps}}", string(jsontps), -1)
		s = strings.Replace(s, "{{.lm}}", string(jsonlm), -1)
		s = strings.Replace(s, "{{.la}}", string(jsonla), -1)
		s = strings.Replace(s, "{{.ls}}", string(jsonls), -1)
		s = strings.Replace(s, "{{.tm}}", string(jsontm), -1)
		s = strings.Replace(s, "{{.ta}}", string(jsonta), -1)
		s = strings.Replace(s, "{{.ts}}", string(jsonts), -1)
		s = strings.Replace(s, "{{.r}}", string(jsonr), -1)
		s = strings.Replace(s, "{{.tt}}", string(jsontt), -1)
		s = strings.Replace(s, "{{.ec}}", string(jsonec), -1)
		s = strings.Replace(s, "{{.er}}", string(jsoner), -1)
		s = strings.Replace(s, "{{.ew}}", string(jsonew), -1)
		s = strings.Replace(s, "{{.et}}", string(jsonet), -1)
		s = strings.Replace(s, "{{.e}}", string(jsone), -1)
		s = strings.Replace(s, "{{.label}}", string(label), -1)
		s = strings.Replace(s, "{{.rp2xx}}", string(rp2xx), -1)

		ctx.Text(iris.StatusOK, s)
	})

	iris.Get("/api/job", func(ctx *iris.Context){
		j := model.Job{}.GetAllJob(session)
		ctx.JSON(iris.StatusOK, j)
	})

	iris.Get("/test/new/custom", func(ctx *iris.Context){
		ctx.Render("test-new-custom.html", nil)
	})

	iris.Get("/test/new/timed", func(ctx *iris.Context){
		ctx.Render("test-new-time.html", nil)
	})

	iris.Get("/test/new/concd", func(ctx *iris.Context){
		ctx.Render("test-new-concurrency.html", nil)
	})

	iris.Post("/test/new", func(ctx *iris.Context){
		typeForm := string(ctx.FormValue("type"))
		fmt.Println("typeForm", typeForm)
		switch typeForm {
		case "cc":
			name := string(ctx.FormValue("name"))
			t := ctx.FormValues("t")
			c := ctx.FormValues("c")
			d := ctx.FormValues("d")
			testsuite := model.Testsuite{}
			for i, tt := range t{
				fmt.Println(t)
				testcase := model.Testcase{}
				testcase.Duration = d[i]
				testcase.Connection = c[i]
				testcase.Thread = tt
				testsuite.AddTestcase(testcase)
			}
			testsuite.Name = name
			testsuite.Domain = "custom"
			testsuite.Save(session)
		case "td":
			fmt.Println("do time domain")
			name := string(ctx.FormValue("name"))
			t := ctx.FormValues("t")
			c := ctx.FormValue("c")
			d := ctx.FormValues("d")
			testsuite := model.Testsuite{}
			for i, dd := range d{
				fmt.Println(t)
				testcase := model.Testcase{}
				testcase.Duration = dd
				testcase.Connection = string(c)
				testcase.Thread = t[i]
				testsuite.AddTestcase(testcase)
			}
			testsuite.Name = name
			testsuite.Domain = "time"
			testsuite.Save(session)
		case "cd":
			name := string(ctx.FormValue("name"))
			t := ctx.FormValues("t")
			c := ctx.FormValues("c")
			d := ctx.FormValue("d")
			testsuite := model.Testsuite{}
			for i, cc := range c{
				fmt.Println(t)
				testcase := model.Testcase{}
				testcase.Duration = string(d)
				testcase.Connection = cc
				testcase.Thread = t[i]
				testsuite.AddTestcase(testcase)
			}
			testsuite.Name = name
			testsuite.Domain = "concurrency"
			testsuite.Save(session)
		}

		ctx.Redirect("/")
	})

	iris.Post("/wrk", func(ctx *iris.Context){
		bUrl := ctx.FormValue("url")
		testcase := string(ctx.FormValue("testcase"))
		name := string(ctx.FormValue("name"))
		method := string(ctx.FormValue("method"))
		keys := ctx.FormValues("key")
		values := ctx.FormValues("value")

		ctx.Redirect("/")

		if bUrl == nil{
			return;
		}
		url := string(bUrl)

		keyValue := map[string]string{}
		for i, key := range keys{
			keyValue[key] = values[i]
		}

		r := model.Request{}
		r.Url = url
		r.Method = method
		r.KeyValueToLoad(keyValue)
		j := model.Job{}.NewInstance(session, testcase, name, r)
		modelchan <- j
	})

	iris.Get("/rerun/:unique", func(ctx *iris.Context){
		unique := ctx.Param("unique")
		job := model.Job{}.Find(session, unique)
		modelchan <- job.ReRunWrk(session)
		ctx.Redirect("/")
	})

	iris.Post("/lfu", func(ctx *iris.Context){
		url := string(ctx.FormValue("url"))
		fmt.Println(url)

		doc, err := goquery.NewDocument(url)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, map[string]string{"status":"err", "err":err.Error()})
			return;
		}

		var forms []model.Form

		doc.Find("form").Each(func(i int, s *goquery.Selection){
			var form model.Form
			form.Action, _ = s.Attr("action")
			form.Method, _ = s.Attr("method")
			s.Find("input").Each(func(j int, t *goquery.Selection){
				for k := 0 ; k < len(t.Nodes) ; k++{
					var input model.Input
					input.Name,_ = t.Attr("name")
					form.Input = append(form.Input, input)
				}
			})
			forms = append(forms, form)
		})

		ctx.JSON(iris.StatusOK, map[string]interface{}{"status":"ok", "data":forms})
	})

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func (so socketio.Socket){
		so.Join("real-time")
		fmt.Println("connection in")
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	iris.Handle(iris.MethodGet, "/socket.io/", iris.ToHandler(server))
	iris.Handle(iris.MethodPost, "/socket.io/", iris.ToHandler(server))

	go func(){
		wg := sync.WaitGroup{}
		for;;{
			select {
			case j := <-modelchan:
				wg.Add(1)
				go func() {
					testsuite := model.Testsuite{}.Find(session, j.TestcaseName)
					t := j.Unique
					selectedTestSuite := testsuite
					for i, testcase := range selectedTestSuite.Testcase {
						time.Sleep(10 * time.Second)
						j.RunWrk(testcase, selectedTestSuite.Domain, t, mongochan)
						server.BroadcastTo("real-time",
							t,
							`{"Unique":"` + t + `", "IsComplete":false, "Progress":` + fmt.Sprintf("%.2f", float64((i+1))/float64(len(selectedTestSuite.Testcase))*100.0) + `}`)
					}
					j.Complete(session)
					server.BroadcastTo("real-time", t, `{"Unique":"` + t + `", "IsComplete":true, "Progress":100}`)
					wg.Done()
				}()
				wg.Wait()

			}
		}
	}()

	go func(){
		for{
			select {
			case wrkResult := <-mongochan:
				go wrkResult.Save(session)
			}
		}
	}()

	iris.Listen(":8080")

}
