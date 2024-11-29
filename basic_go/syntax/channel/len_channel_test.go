package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func producer(c chan<- int) {
	var i = 1

	for {
		ok := trySend(c, i)
		if ok {
			i++
			fmt.Printf("[producer]: send [%d] to channel\n", i)
			continue
		}
		time.Sleep(time.Second)
		fmt.Printf("[producer]: try send [%d], but channel is full\n", i)
	}
}

func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

func consumer(c <-chan int) {
	for {
		v, ok := tryRev(c)
		if !ok {
			fmt.Println("[consumer]: try to recv from channel, but the channel is empty")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("[consumer]: recv [%d] to channel\n", v)
		if v > 3 {
			fmt.Println("[consumer]: exit")
			return
		}
	}
}

func tryRev(c <-chan int) (int, bool) {
	select {
	case v := <-c:
		return v, true
	default:
		return 0, false
	}
}

func TestLenChannel(t *testing.T) {
	c := make(chan int, 3)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		producer(c)
		wg.Done()
	}()

	go func() {
		consumer(c)
		wg.Done()
	}()

	wg.Wait()
}
