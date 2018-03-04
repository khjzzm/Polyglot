package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	//↓ 함수의 호출이 지연됨.
	defer world() // 현재 함수(main())가 끝나기 직전에 호출
	hello()
	defer world() // 현재 함수(main())가 끝나기 직전에 호출
	hello()
	hello()


	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}