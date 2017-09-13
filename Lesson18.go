//json http://www.cnblogs.com/yangshiyu/p/6942414.html
package main

import (
	"encoding/json"
	"fmt"
)

type ColorGroup struct {
	ID        int   `json:"id"` //自定义json字符串中的字段名
	Name      string
	ProductID int64 `json:"-"` // 表示不进行序列化
	Colors    []string
}

//omitempty，tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
//我们在序列化或者反序列化的时候，可能结构体类型和需要的类型不一致，这个时候可以指定,支持string,number和boolean(price,string)
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,omitempty"`
	Number    int     `json:"number"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,omitempty"`
}

//
type Goods struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func main() {
	group := ColorGroup{1, "Reds", 123, []string{"Crimson", "Red", "Ruby", "Maroon"}}
	toJson(group)

	product := Product{}
	product.Name = "Xiao mi 6"
	product.IsOnSale = false
	product.Number = 10000
	product.Price = 2499.00
	product.ProductID = 0

	toJson(product)

	jsonData := `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`

	//json转换实体类需要&
	pr := &Goods{}
	e := json.Unmarshal([]byte(jsonData), pr)
	fmt.Println(e, pr)

	//json转成实体类
	mapData := make(map[string]interface{})
	e2 := json.Unmarshal([]byte(jsonData), &mapData)
	fmt.Println(e2, mapData)


	data :=make(map[string]interface{})
	data["lv"]=1
	data["tan"]="2"
	data["xi"]=string(2)

	toJson(data)


}

func toJson(any interface{}) {
	result, err := json.Marshal(any)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(string(result))
}
