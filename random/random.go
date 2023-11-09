package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mySlice := []string{
		"a",
		"b",
		"c",
	}
	name := rand.Intn(len(mySlice))
	fmt.Println(name)
}
