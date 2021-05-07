package sync_demo

import (
	"fmt"
	"sync"
)

// http://topgoer.com/%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B/sync.html#synconce
func task1(m, n string) {
	fmt.Println(m, n)
}

func task1Wrapper(m, n string) func() {
	// 使用闭包做一层封装
	return func() {
		task1(m, n)
	}
}

func SyncDemo1() {
	var once sync.Once
	once.Do(task1Wrapper("hello", "world"))
	once.Do(task1Wrapper("hello", "world"))
	once.Do(task1Wrapper("hello", "world"))
}
