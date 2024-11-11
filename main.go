package main

import (
	"fmt"
	"runtime"

	puppyLib "github.com/abudygold/puppy"
)

func main() {
	s1 := puppyLib.Bark()
	s2 := puppyLib.Barks()
	s3 := puppyLib.WhenGrownUp("Guguk")

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(puppyLib.Bark())
	fmt.Println(puppyLib.Barks())
	fmt.Println(puppyLib.WhenGrownUp("Guguk"))
}
