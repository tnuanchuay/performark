#!/bin/sh
export GOPATH=$(pwd)
echo 'Getting iris...'
go get -u -v github.com/kataras/iris/iris
echo 'Getting go-socket.io'
go get -v github.com/googollee/go-socket.io
echo 'Getting mgo.v2'
go get -v gopkg.in/mgo.v2
echo 'Building performark...'
go build -o performark src/main/main.go
echo 'done'
