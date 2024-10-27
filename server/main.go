package main

import (
	"fmt"
	"shared/dto"
)

func main() {
	test := dto.Test{
		Ty: "hello",
	}
	fmt.Println("hello", test)
}
