package engine

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type PrintCommand struct {
	Arg string
}

func (p *PrintCommand) Execute(loop Handler) {
	fmt.Println(p.Arg)
}

type Sha1Command struct {
	Arg string
}

func (p *Sha1Command) Execute(loop Handler) {
	h := sha1.New()
	h.Write([]byte(p.Arg))
	bs := h.Sum(nil)
	res := fmt.Sprintf("%x", bs)
	loop.Post(&PrintCommand{Arg: res})
}

type EventLoop struct {
	sync.Mutex
	commandsQueue []Command
	isProcessing  bool
	finish        bool
}

func (loop *EventLoop) Post(cmd Command) {
	loop.Lock()
	loop.commandsQueue = append(loop.commandsQueue, cmd)
	defer loop.Unlock()
}

func (loop *EventLoop) Start() {
	wg.Add(1)
	loop.isProcessing = false
	loop.commandsQueue = []Command{}
	loop.finish = false
	go func() {
		for {
			if !loop.isProcessing && len(loop.commandsQueue) != 0 {
				loop.commandsQueue[0].Execute(loop)
				loop.commandsQueue = loop.commandsQueue[1:]
			}
			if len(loop.commandsQueue) == 0 && loop.finish {
				break
			}
		}
		defer wg.Done()
	}()
}

func (loop *EventLoop) AwaitFinish() {
	loop.finish = true
	wg.Wait()
}
