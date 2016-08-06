package main

import (
	"github.com/kataras/iris"
	"os/exec"
	"fmt"
	"bufio"
	"strings"
	"model"
	"gopkg.in/mgo.v2"
	"time"
	"strconv"
)

func main(){

	iris.Config.IsDevelopment = true

	iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	})

	iris.Static("/assets", "./static/assets", 1)
	iris.Static("/images", "./static/images", 1)

	iris.Get("/", func(ctx *iris.Context){
		ctx.Render("index.html", nil)
	})

	funcChannel := make(chan func(string, string, string), 100)
	mongochan := make(chan model.WrkResult, 100)

	iris.Post("/wrk", func(ctx *iris.Context){
		url := string(ctx.FormValue("url"))
		ctx.Redirect("/")

		funcChannel <- func(c, d, time string){
			var t int = 1
			cc, _ := strconv.Atoi(c)
			if cc >= 4 {
				t = 4
			}

			command := exec.Command("wrk", "-t"+strconv.Itoa(t), "-c"+c, "-d"+d, url)
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
			wrkResult := model.WrkResult{}
			wrkResult.SetData(url, out, time)

			mongochan <- wrkResult
		}
	})

	go func(){
		session, err := mgo.Dial("127.0.0.1")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("performark").C("mark")
		for;;{
			select {
			case f := <- funcChannel:
				go func() {
					t := time.Now().Format("20060102150405")
					f("1", "10s", t)
					//f("10", "10s", t)
					//f("100", "10s", t)
					//f("1k", "10s", t)
					//f("10k", "10s", t)
				}()
			case wrkResult := <- mongochan:
				go wrkResult.Save(c)
			}
		}
	}()

	iris.Listen(":8080")
}
