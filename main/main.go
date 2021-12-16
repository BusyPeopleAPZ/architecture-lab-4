package main

import (
	"bufio"
	"github.com/BusyPeopleAPZ/architecture-lab-4/engine"
	"os"
	"strings"
)

func main() {
	inputFile := "file.txt"

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}

func parse(str string) engine.Command {
	s := strings.Split(str, " ")
	if s[0] == "print" {
		return &engine.PrintCommand{Arg: s[1]}
	}
	if s[0] == "sha1" {
		return &engine.Sha1Command{Arg: s[1]}
	}
	return &engine.PrintCommand{Arg: "Error parsing expression"}
}
