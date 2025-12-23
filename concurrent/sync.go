package concurrent

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

func Cond() {
	var (
		status int64
		mux    sync.Mutex
		cond   = sync.NewCond(&mux)
		ctx    = context.Background()
	)
	ctx.Deadline()
	for i := 0; i < 10; i++ {
		go listen(cond, &status)
	}

	go broadcast(cond, &status)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcast(c *sync.Cond, status *int64) {
	c.L.Lock()
	atomic.StoreInt64(status, 1)
	c.Broadcast()
	c.L.Unlock()
}

func listen(c *sync.Cond, status *int64) {
	c.L.Lock()
	for atomic.LoadInt64(status) != 1 {
		c.Wait()
	}
	fmt.Println("listent")
	c.L.Unlock()
}

func WaitGroup() {
	var (
		wg sync.WaitGroup
		ch = make(chan struct{}, 10)
	)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(index int) {
			defer wg.Done()
			fmt.Printf("id: %d \n", index)
			<-ch
		}(i)
	}
	wg.Wait()
}

func TrunPrint() {
	var (
		wg               sync.WaitGroup
		strChan, numChan = make(chan struct{}), make(chan struct{})
	)

	wg.Add(2)
	// 打印字母
	go func() {
		defer wg.Done()
		for i := 'a'; i <= 'z'; i++ {
			<-strChan
			fmt.Println(string(i))
			numChan <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-numChan
			fmt.Println(i)
			strChan <- struct{}{}
		}
	}()

	strChan <- struct{}{}
	wg.Wait()
}

func TimeOut() {
	var (
		ch = make(chan int, 1000)
		wg sync.WaitGroup
	)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		ch <- i
		go func(index int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				fmt.Printf("id: %d timeout \n", index)
				return
			default:
				fmt.Printf("id: %d \n", index)
			}
		}(i)
	}

	<-ctx.Done()
	fmt.Println(" timeout")
	close(ch)
	wg.Wait()
}

// 开启100个协程，顺序打印1-1000，且保证协程号1的，打印尾数为1的数字
func OrderPrint() {
	var (
		chMap = make(map[int]chan int, 100)
		done  = make(chan struct{}, 100)
	)
	for i := 1; i <= 100; i++ {
		// 带缓冲chan减少阻塞
		chMap[i] = make(chan int)
	}

	for i := 1; i <= 100; i++ {
		go func(index int) {
			for num := range chMap[index] {
				fmt.Printf("协程id: %d, 打印数字: %d \n", index, num)
			}
			done <- struct{}{}
		}(i)
	}

	for i := 1; i <= 1000; i++ {
		if i%100 == 0 {
			chMap[100] <- i
			continue
		}
		chMap[i%100] <- i
	}

	for _, ch := range chMap {
		close(ch)
	}
	
	for i := 1; i <= 100; i++ {
		<- done
	}
}
