package main

import (
	"net/http"
	"io"
	"time"
	"log"
)

//存储自己处理请求的函数
var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHander{},
		ReadTimeout: 5 * time.Second,
	}
	//初始化map
	mux =make(map[string]func(http.ResponseWriter, *http.Request))
	//设置路由对应的方法
	mux["/"] = sayHello3
	mux["/bye"] = sayBye

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

type myHander struct{}

func (*myHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//利用ok patten获取需要调用的函数
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL"+r.URL.String())
}

func sayHello3(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello woeld ,this is version 3")
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Good bye")
}
