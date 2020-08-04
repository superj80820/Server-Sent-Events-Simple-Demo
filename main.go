package main

import (
	"io"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SSEHandler(c *gin.Context) {
	// 創建一個Channel，讓訊息以Server以外的goroutine傳送
	chanStream := make(chan int)
	go func() {
		defer close(chanStream)
		for i := 0; i < 5; i++ {
			chanStream <- i
			time.Sleep(time.Second * 1)
		}
	}()
	// 接收到goroutine的訊息後，透過c.SSEvent以`push-message`的event名字傳送
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			c.SSEvent("push-message", msg)
			return true
		}
		return false
	})
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// 透過一般Get API連線，即可使用SSE功能
	r.GET("/", SSEHandler)
	r.Run("0.0.0.0:3000")
}
