//http 1.0
package main

import (
	"net/http"
	"io"
	"log"
)

func main() {
	//设置路由
	http.HandleFunc("/", sayHello)
	//设置监听端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello woeld ,this is version 1")
}
