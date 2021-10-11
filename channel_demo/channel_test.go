package channel_demo

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var wait sync.WaitGroup
var pipe chan bool = make(chan bool)

func f1() {
	defer wait.Done() // -1
F1:
	for {
		fmt.Println("hello world")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-pipe:
			break F1
		default:
			// nothing to do
		}
	}
	fmt.Println("f1 over")
	time.Sleep(time.Second * 2)
}

func TestChan1(t *testing.T) {
	wait.Add(1)
	go f1()
	time.Sleep(time.Second * 5)
	pipe <- true // 发送通知
	wait.Wait()
	fmt.Println("main over")
}

/*
使用 context.Context 发送通知
*/
func f2(ctx context.Context) {
	defer wait.Done() // -1
F1:
	for {
		fmt.Println("hello world")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break F1
		default:
			// nothing to do
		}
	}
	fmt.Println("f1 over")
	time.Sleep(time.Second * 2)
}

func TestChan2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background()) // context
	wait.Add(1)
	go f2(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 发出通知 -> 通知子 goroutine 结束
	wait.Wait()
	fmt.Println("main over")
}
