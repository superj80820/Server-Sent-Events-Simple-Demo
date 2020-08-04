# Server-Sent-Events-Simple-Demo

此`Repository`透過`Server-Sent-Events(SSE)`來實現`Golang Server`主動推送訊息給`Web Client`

## 什麼是SSE

特性:

* 透過`HTTP Stream`的特性，實現`Server`主動推播的功能
* 由`HTML5`規範原生支持
* 瀏覽器端已原生實作斷線時的`retry connection handle`

場景:

* 天氣即時佈告欄推播
* 聊天室: 有許多人認為聊天室就用`Websocket`即可，但`Websocket`傳遞訊息的方式非常簡單，無法像`Restful API`擁有4xx等`Error Code`，所以要怎麼表示這些錯誤要開發者手動實作。如果透過`SSE`推播 + `Restful API`取值，就可以善用`Restful API`的`Error Code`特性來達到聊天效果。甚至搭配對`Protocol`更嚴謹的[GRPC-Web](https://github.com/grpc/grpc-web)也是很好的方案。

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
