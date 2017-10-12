/**  
* Date: 2017-10-12 
* Time: 13:53 
* Description: RPC的创建
*/
package main

import (
	"qiniupkg.com/x/errors.v7"
	"time"
	"net/rpc"
	"net/http"
	"fmt"
)

type Target struct {
	Name    string
	Age     int
	Address string
	Id      int
}

type Rpc int

func (r *Rpc) GetTarget(id int, target *Target) error {
	if id == 0 {
		return errors.New("id is zero")
	}
	fmt.Println("开始处理数据！")
	target.Age = time.Now().Hour()
	target.Id = id
	target.Address = "四川成都双流"
	target.Name = "念阳枭"
	return nil
}

func main() {
	targetRpc := new(Rpc)
	rpc.Register(targetRpc)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err)
	}
}

