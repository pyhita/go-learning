package context

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Withcancel：发送取消信号
func Context1() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1

		go func() {
			for {
				select {
				case <-ctx.Done():
					log.Println("process finished, child go routine exit.")
					return
				case dst <- n:
					n++
				}
			}
		}()

		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		log.Printf("get %v from channel", n)
		if n == 5 {
			break
		}
	}

	log.Println("process finished, main go routine exit.")
}

// Withtimeout: 带超时的取消
func Context2() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		log.Println("overslept")
	// 过了超时时间之后，done中的channel 被关闭
	// goroutine 从Done 中苏醒
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

// Withvalue：Context 携带键值对
func Context3() {

	a := context.Background()
	b := context.WithValue(a, "k1", "v1")
	c := context.WithValue(b, "k2", "v2")
	d := context.WithValue(c, "k1", "v3")

	fmt.Printf("k1 of b: %s\n", b.Value("k1"))
	fmt.Printf("k1 of d: %s\n", d.Value("k1"))
	fmt.Printf("k2 of d: %s\n", d.Value("k2"))
}

// cancel 的传播
func Context4() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	slowFunc := func(ctx context.Context, i int) {
		childCtx, cancel := context.WithTimeout(ctx, time.Millisecond*800)
		defer cancel()

		fmt.Printf("query No. %d\n", i)
		select {
		case <-childCtx.Done():
			fmt.Printf("child context err: %v\n", childCtx.Err())
		}
	}

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("parent context err: %v\n", ctx.Err())
			return
		default:
			slowFunc(ctx, i)
		}
	}
}
