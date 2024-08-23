package base

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 简单的 Web 爬虫
//这个示例将创建一个简单的 Web 爬虫，使用 goroutine 并发地抓取多个 URL，并使用 channel 收集结果。

type Result struct {
	URL  string
	Body string
	Err  error
}

func fetch(url string) Result {
	resp, err := http.Get(url)
	if err != nil {
		return Result{URL: url, Err: err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Result{URL: url, Err: err}
	}
	return Result{URL: url, Body: string(body)}
}

// GrapData 抓取数据
func GrapData() {
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
	}
	var wg sync.WaitGroup
	// 创建一个带缓冲区的通道,通道类型为Result,这样可以避免通道之间阻塞的问题 允许数据先暂存到通道 主goroutine可稍后进行处理 类似Java中的缓冲队列
	results := make(chan Result, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			results <- fetch(url)
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Println("err result: ", result)
		} else {
			fmt.Println("success result:", result)
		}
	}
}

// goroutine 通道 与 select
func testSelectAndChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "channel1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "channel2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
