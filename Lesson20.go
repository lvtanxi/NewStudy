//http 2.0
package main

import (
	"net/http"
	"io"
	"log"
	"os"
)

func main() {
	//设置mux
	mux := http.NewServeMux()

	mux.Handle("/", &muxHander{})
	//设置访问控制器
	mux.HandleFunc("/hello", sayHello2)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//设置静态文件路径
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))
	//设置监听端口
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type muxHander struct{}

func (*muxHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL"+r.URL.String())
}

func sayHello2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello woeld ,this is version 2")
}
