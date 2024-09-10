package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello Gophers Again!")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
