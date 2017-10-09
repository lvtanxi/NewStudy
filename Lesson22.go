/**  
* Date: 2017-10-09 
* Time: 16:54 
* Description: 文件操作
*/
package main

import (
	"os"
	"fmt"
)

func main() {
	testMkdir()
	testWrite()
	testRead()
	testRemove()
}

//删除文件
func testRemove() {
	err:=os.Remove("test.txt")
	if err != nil {
		fmt.Println(err)
	}
}

//打开文件，并读取文件
func testRead() {
	userFile := "test.txt"
	file, err := os.Open(userFile)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	buf :=make([]byte,1024)
	for  {
		n,err :=file.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}else if 0==n{
			break
		}
		os.Stdout.Write(buf[:n])
	}

}

//创建文件，并写入信息
func testWrite() {
	userFile := "test.txt"
	fileOut, err := os.Create(userFile)
	defer fileOut.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		fileOut.WriteString("Just a test \r\n")
		fileOut.Write([]byte("Just a test \r\n"))
	}
}

//简单创建
func testMkdir() {
	os.Mkdir("as", 0777)
	os.MkdirAll("as/as/as", 0777)
	err := os.Remove("as")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("as")
}
