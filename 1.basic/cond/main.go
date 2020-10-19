package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// https://mp.weixin.qq.com/s/9FjBJVhRBVSBV4CVpCfChg
// Golang的sync包中的Cond实现了一种条件变量，主要用来解决多个读协程等待共享资源变成ready的场景。
// 在使用Cond的时候，需要特别注意下：每个Cond都会关联一个Lock（*sync.Mutex or *sync.RWMutex），
// 当修改条件或者调用Wait方法时，必须加锁，保护condition。

// https://stackoverflow.com/questions/51371587/broadcast-in-golang

func main() {
	var m sync.Mutex
	c := sync.NewCond(&m)

	ready := make(chan struct{}, 10)
	isReady := false

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			m.Lock()

			time.Sleep(time.Duration(rand.Int63n(2)) * time.Second)

			ready <- struct{}{} // 运动员i准备就绪
			for !isReady {
				c.Wait() //必须加锁，使协程处于阻塞状态
			}
			log.Printf("%d started\n", i)
			m.Unlock()
		}()
	}

	// false broadcast
	//c.Broadcast()

	// 裁判员检查所有的运动员是否就绪
	for i := 0; i < 10; i++ {
		<-ready
	}

	// 运动员都已准备就绪，发令枪响, broadcast
	// m.Lock()
	isReady = true //共享资源
	c.Broadcast() // 共享资源变成ready的场景
	// m.Unlock()

	time.Sleep(time.Second)
}
