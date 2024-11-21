package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "Hello golang http!")
}

func main() {
	exitCh := make(chan int)
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)
	s := &http.Server{
		Addr:    ":8091",
		Handler: http.DefaultServeMux,
	}
	go func() {
		// 启动web服务，监听9090端口
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	go gracefulShutdown(s, exitCh)

	<-exitCh
}

func gracefulShutdown(srv *http.Server, exitCh chan int) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Kill)

	sig := <-signalChan
	log.Printf("catch signal, %+v", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second) // 60秒后退出
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Printf("server exiting")
	close(exitCh)
}
