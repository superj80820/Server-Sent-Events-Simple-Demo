# Server-Sent-Events-Simple-Demo

此`Repository`透過`Server-Sent-Events(SSE)`來實現`Golang Server`主動推送訊息給`Web Client`

## 需要安裝

* `docker`
* `docker-compose`

## 怎麼運作？

1. `docker-compos up`
2. 瀏覽器連接至`localhost:5000`

## 流程

![](https://imgur.com/AXWDxK7.jpg)

## 解釋

整體我都用註釋寫在`code`裡面了，主要在此兩個檔案:

* [index.html](./index.html)
* [main.go](./main.go)

## 參考

* [使用 server-sent 事件 - Web APIs | MDN](https://developer.mozilla.org/zh-TW/docs/Web/API/Server-sent_events/Using_server-sent_events)
* [Server-Sent Events 教程 - 阮一峰的网络日志](https://www.ruanyifeng.com/blog/2017/05/server-sent_events.html)
* [go - How to write a stream API using gin-gonic server in golang? Tried c.Stream didnt work - Stack Overflow](https://stackoverflow.com/questions/44825244/how-to-write-a-stream-api-using-gin-gonic-server-in-golang-tried-c-stream-didnt)
* [JavaScript - Polling、WebSocket 與 SSE 介紹 - iT 邦幫忙::一起幫忙解決難題，拯救 IT 人的一天](https://ithelp.ithome.com.tw/articles/10230335)
