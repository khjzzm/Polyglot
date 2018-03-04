package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	bytes, err := ioutil.ReadFile("textDirectory/fileio3.txt")
	if(err != nil) {
		panic(err)
	}
	fmt.Println(string(bytes))

	key := strings.SplitAfter(string(bytes), ",")
	fmt.Print(key)
}
