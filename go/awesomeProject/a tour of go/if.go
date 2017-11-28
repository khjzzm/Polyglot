package main

import "fmt"

func sumFrom1ToGaussRecursive(x int) int {
	var result int = 0
	if(x%2==0){
		result += (1+x) * (x/2)
	}else{
		result += sumFrom1ToGaussRecursive(x-1) + x
	}
	return result
}

func main() {
	fmt.Print(sumFrom1ToGaussRecursive(100))
}
