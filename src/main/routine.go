package main

import (
	"os/exec"
	"fmt"
)

func main(){
	go func() {
		command := exec.Command("wrk", "-t4", "-c5", "-d2", "http://127.0.0.1")
		fmt.Println(command.Args)
		out, _ := command.Output()
		fmt.Println(string(out))
	}()

	for;;{

	}
}
