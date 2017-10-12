/**  
* Date: 2017-10-12 
* Time: 09:44 
* Description: RPC调用
*/
package main

import (
	"net/rpc"
	"fmt"
)

type Tx struct {
	Name    string
	Age     int
	Address string
	Id      int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("DialHTTP error", err)
	}
	target := &Tx{}
	err = client.Call("Rpc.GetTarget", 12, target)
	if err != nil {
		fmt.Println("Call error", err)
	}
	fmt.Println(target)
}
