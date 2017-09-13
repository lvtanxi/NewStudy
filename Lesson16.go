//并发
package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

func main() {
	selectTimeOut()
}
//设置select超时
func selectTimeOut() {
	c :=make(chan bool)
	select {
	case v:=<-c:
		fmt.Println(v)
	case <-time.After(3*time.Second):
		fmt.Println("time out")

	}
}
func selectSetVal() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
	//随机向channel中设置值
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
func selectTest() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"

	c1 <- 3
	c2 <- "hello"

	close(c1)
	close(c2)

	<-o
}

func waitGroupTest() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go doTest(&wg, i)
	}
	wg.Wait()
}

func doTest(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}

func cacheChannel() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go goNum(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func goNum(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}

func goxx() {
	c := make(chan bool)
	go func() {
		fmt.Println("goxx go go go")
		<-c
	}()
	c <- true
}

//channel缓存
func goAndChannelWithCache() {
	//定义一个缓存为1的channel,有缓存的goroutine是异步的
	c := make(chan bool, 1)
	go func() {
		fmt.Println("goAndChannelWithCache go go go")
		c <- true
	}()
	<-c
	fmt.Println(">>>>>结果")
}

//channel与range联合使用
func goAndChannelWithRange() {
	c := make(chan bool)
	go func() {
		fmt.Println("goAndChannelWithRange go go go")
		c <- true
		//手动关闭channel
		close(c)
	}()

	//遍历或去channel中的值
	for v := range c {
		fmt.Println(v)
	}
}

//单个channel，实现消息的发送与接收,无缓存是同步的
func goAndChannel() {
	//申明一个channel
	c := make(chan bool)
	//启动goroutine
	go func() {
		fmt.Println("goAndChannel go go go !!!")
		//向channel中填充值
		c <- true
	}()
	//从channel取出值
	<-c
	fmt.Println(">>>>>结果")
}

//投机取巧的方式执行线程
func xx() {
	//启动goroutine(相当于线程)
	go fgo()

	go func() { fmt.Println("匿名 Go go go") }()

	time.Sleep(2 * time.Second)
}

func fgo() {
	fmt.Println("Go go go ")
}
