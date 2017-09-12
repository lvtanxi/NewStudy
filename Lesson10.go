//Panic
package main

import "fmt"

func main() {
	pa()
	pb()
	pc()
}

func pa() {
	fmt.Println("Func a")
}

func pb() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in b")
}

func pc() {
	fmt.Println("Func c")
}
