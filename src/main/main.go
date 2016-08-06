package main

import (
	"github.com/kataras/iris"
	"model"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
)

func main(){

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	iris.Config.IsDevelopment = true

	iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	})

	iris.Static("/assets", "./static/assets", 1)
	iris.Static("/images", "./static/images", 1)

	iris.Get("/", func(ctx *iris.Context){
		ctx.Render("index.html", nil)
	})

	iris.Get("/api/job", func(ctx *iris.Context){
		j := model.Job{}.GetAllJob(session)
		ctx.JSON(iris.StatusOK, j)
	})

	modelChan := make(chan *model.Job, 100)
	mongochan := make(chan model.WrkResult, 100)

	iris.Post("/wrk", func(ctx *iris.Context){
		url := string(ctx.FormValue("url"))
		ctx.Redirect("/")
		j := model.Job{}.NewInstance(url, session)
		modelChan <- j
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
		for;;{
			select {
			case j := <- modelChan:
				go func() {
					t := j.Unique
					time.Sleep(2 * time.Second)
					j.RunWrk("1", "10s", t, mongochan)
					server.BroadcastTo("real-time", t, `{"Unique":"` + t + `", "IsComplete":false, "Progress":17}`)
					time.Sleep(2 * time.Second)
					j.RunWrk("10", "10s", t, mongochan)
					server.BroadcastTo("real-time", t, `{"Unique":"` + t + `", "IsComplete":false, "Progress":34}`)
					time.Sleep(2 * time.Second)
					//j.Function("100", "10s", t)
					server.BroadcastTo("real-time", t, `{"Unique":"` + t + `", "IsComplete":false, "Progress":51}`)
					time.Sleep(2 * time.Second)
					//j.Function("1k", "10s", t)
					server.BroadcastTo("real-time", t, `{"Unique":"` + t + `", "IsComplete":false, "Progress":68}`)
					time.Sleep(2 * time.Second)
					//j.Function("10k", "10s", t)
					server.BroadcastTo("real-time", t, `{"Unique":"` + t + `", "IsComplete":false, "Progress":83}`)

					j.Complete(session)
					server.BroadcastTo("real-time", t, `{"Unique":"` + t + `", "IsComplete":true, "Progress":100}`)
				}()
			case wrkResult := <- mongochan:
				go wrkResult.Save(session)
			}
		}
	}()

	iris.Listen(":8080")
}
