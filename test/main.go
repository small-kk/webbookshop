package main

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
)

func main() {
	req, err := http.NewRequest("GET", "https://www.baidu.com/", nil)
	if err != nil {
		panic(err)
	}

	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			fmt.Println("GetConn:", hostPort)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("GotConn Reused=%v WasIdle=%v IdleTime=%v\n", connInfo.Reused, connInfo.WasIdle, connInfo.IdleTime)
			fmt.Printf("GotConn localAddr: %s\n", connInfo.Conn.LocalAddr())
			fmt.Printf("GotConn remoteAddr: %s\n", connInfo.Conn.RemoteAddr())
		},
		PutIdleConn: func(err error) {
			fmt.Printf("PutIdleConn err=%v\n", err)
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    1, //最大Idle
			MaxConnsPerHost: 2, //每个host最大conns
		}}

	for i := 1; i <= 10; i++ {
		fmt.Printf("==============第[%d]次请求================\n", i)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		//io.ReadAll(resp.Body) //读取body里面的内容
		//resp.Body.Close()
		fmt.Println(resp.Status)
	}
}
