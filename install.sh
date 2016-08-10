#!/bin/sh
export GOPATH=$(pwd)
go get -u github.com/kataras/iris/iris
go get github.com/googollee/go-socket.io
go get gopkg.in/mgo.v2
go build -o performark src/main/main.go
