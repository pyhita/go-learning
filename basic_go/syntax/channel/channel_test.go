package channel

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}

// 测试channel 进行信号传递：一对一的通信
type signal struct{}

func worker() {
	println("worker is working ... ")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)

	go func() {
		println("worker start to work ... ")
		f()
		c <- signal{}
	}()

	return c
}

func TestChannel2(t *testing.T) {
	println("main start a worker ... ")

	c := spawn(worker)
	<-c
	println("worker finish working ... ")
}

// 测试channel 进行信号传递：一对一的通信
func worker2(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-groupSignal
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
		}(i + 1)
	}

	go func() {
		wg.Wait()
		// 所有的 worker 结束工作，发送信号
		c <- signal{}
	}()

	return c
}

func TestChannel3(t *testing.T) {
	println("main starts a group workers ... ")
	groupSignal := make(chan signal)
	c := spawnGroup(worker2, 5, groupSignal)
	time.Sleep(5 * time.Second)

	close(groupSignal)
	<-c
	println("worker finish working ... ")
}

// 测试channel 用作计数信号量
var active = make(chan signal, 3)
var jobs = make(chan int, 10)

func TestChannel4(t *testing.T) {
	// 一个 goroutine 不停创建job
	go func() {
		for i := 0; i < 10; i++ {
			jobs <- i + 1
		}
		close(jobs)
	}()

	// 主goroutine 创建其他goroutine 进行消费
	var wg sync.WaitGroup
	for n := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- signal{}
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(n)
	}

	wg.Wait()
}
