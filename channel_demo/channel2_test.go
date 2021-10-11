package channel_demo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
chan int   双向通道
chan<- int 单向通道 只发
<-chan int 单向通道 只收
*/
func producer(producerName string, count int, pipe chan<- string) { // 单向通道
	for i := 0; i < count; i++ {
		pipe <- fmt.Sprintf("[%s %v]", producerName, rand.Int31()) // 向管道内发送数据
	}
}

func consumer(consumerName string, pipe <-chan string) {
	for {
		message := <-pipe // 从管道内接收数据
		fmt.Printf("%s:%s\n", consumerName, message)
	}
}

func node1(chanName chan string) { // 双向通道
	chanName <- "hello channel" // 将字符串 "hello world" 输送到管道里
}

func node2(chanName chan string, content string) {
	chanName <- content
}

func node3(chanName chan string) {
	x := <-chanName // 从管道里获取数据
	fmt.Println("node3:", x)
}

func TestChannel(t *testing.T) {
	// 创建通道
	pipe := make(chan string)
	go producer("生产者A", 2, pipe)
	go producer("生产者B", 5, pipe)
	time.Sleep(time.Millisecond)
	go consumer("消费者1", pipe)
	go consumer("消费者2", pipe)
	go consumer("消费者3", pipe)
	node1(pipe)
	node2(pipe, "国庆节")
	go node3(pipe)
	time.Sleep(time.Second * 1)
}
