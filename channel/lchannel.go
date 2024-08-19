package channel

import (
	"fmt"
	"time"
)

func ReadData() {
	// 创建通道数据 类型为int型
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			// 写入通道之后 直到有其他的goroutines从通道中读取数据 否则该goroutines一直被阻塞
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Println("received:", value)
	}
}

func WriteData() {

}
