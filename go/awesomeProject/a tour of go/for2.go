package main

import "fmt"

// Go에서 "While" 사용하기
func main() {
	sum := 1
	for sum < 1000 {
		sum = sum + sum
	}
	fmt.Println(sum)

}
