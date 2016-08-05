package main

import (
	"github.com/kataras/iris"
	"os/exec"
	"fmt"
	"io/ioutil"
	"model"
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

	iris.Post("/wrk", func(ctx *iris.Context){
		url := ctx.FormValue("url")
		ctx.Redirect("/")
		go func(url []byte){
			command := exec.Command("wrk", "-t4", "-c5", "-d1", string(url), ">>", "out.txt")
			fmt.Println(command.Args)
			out, err := command.Output()
			if err != nil {
				ioutil.WriteFile("out.txt", []byte(err.Error()), 0644)
			}else{
				ioutil.WriteFile("out.txt", out, 0644)
			}

			wrkResult := model.WrkResult{}

			wrkResult.Url = string(url)
			wrkResult.SetDuration(out)
			wrkResult.SetThread(out)
			wrkResult.SetConnection(out)
			fmt.Println("wrkResult.SetConnection", wrkResult.Connection)
		}(url)
	})

	iris.Listen(":8080")
}
