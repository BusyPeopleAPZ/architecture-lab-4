package main

import (
	"bufio"
	_ "fmt"
	"os"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type printCommand struct {
	arg string
}
//func (p *printCommand) Execute(loop engine.Handler) {
//	fmt.Println(p.arg)
//}

func main() {
	println("Hello")
	inputFile := "file.txt"
	//eventLoop := new(engine.EventLoop)
	///eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			println(cmd)
			//eventLoop.Post(cmd)
		}
	}
	//eventLoop.AwaitFinish()
}

func parse(str string) Command {
	return nil
}
