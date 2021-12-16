package main

import (
	"bufio"
	"github.com/BusyPeopleAPZ/architecture-lab-4/engine"
	"os"
	"strings"
)

func main() {
	inputFile := "file.txt"

	eventLoop := new(EventLoop)
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

func parse(str string) Command {
	s := strings.Split(str, " ")
	if s[0] == "print" {
		return &PrintCommand{arg: s[1]}
	}
	if s[0] == "sha1" {
		return &Sha1Command{arg: s[1]}
	}
	return &PrintCommand{arg: "Error parsing expression"}
}
