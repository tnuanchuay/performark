package model

import (
	"io/ioutil"
	"fmt"
)

var DefaultScript = ""

func loadDefaultScrip() (string, error){
	buf, err := ioutil.ReadFile("./lua/default.lua")
	if err != nil {
		fmt.Println(err)
		return "", error("Error\tDefault lua script cannot be loaded")
	}else{
		return buf, nil
	}
}

func Init() error{
	DefaultScript,err := loadDefaultScrip()
	if err != nil {
		return error("Initialize Failed")
	}
}
