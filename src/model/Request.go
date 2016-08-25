package model

import (
	"io/ioutil"
	"fmt"
)

type Request struct{
	Url		string
	Method		string
	Load		string
	CustomScript	string
}

func (r *Request) GenerateScript(filename string){
	script := ""
	script += fmt.Sprintf(LUA_method, r.Method)
	if len(r.Load) > 0{
		script += fmt.Sprintf(LUA_load, r.Load)
		script += fmt.Sprintf(LUA_contentType, "application/x-www-form-urlencoded")
	}
	ioutil.WriteFile(fmt.Sprintf("lua/%s.lua", filename), []byte(script), 0644)
}

func (r *Request) KeyValueToLoad(keyValue map[string]string){
	for key, value := range keyValue{
		if len(key) > 0 {
			r.Load += key + "=" + value + "&"
		}
	}
}