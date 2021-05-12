package channel_demo

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var exitChan = make(chan bool)

func f1() {
	defer wait.Done() // -1
F1:
	for {
		fmt.Println("hello world")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break F1
		default:
			// nothing to do
		}
	}
	fmt.Println("f1 over")
	time.Sleep(time.Second * 2)
}

func ChanDemo() {
	wait.Add(1)
	go f1()
	time.Sleep(time.Second * 5)
	exitChan <- true // 发送通知
	wait.Wait()
	fmt.Println("main over")
}
