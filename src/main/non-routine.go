package main

import (
	"os/exec"
	"fmt"
	"bufio"
	"strings"
)

func main(){

	command := exec.Command("wrk", "-t4", "-c5", "-d2", "http://127.0.0.1")
	fmt.Println(command.Args)
	cmdReader, _ := command.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	go func() string{
		var out string
		for scanner.Scan() {
			out = fmt.Sprintf("%s\n%s", out, scanner.Text())
			if strings.Contains(out, "Transfer"){
				break;
			}
		}

		return out
	}()


	command.Start()
	command.Wait()

}
