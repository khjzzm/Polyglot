package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	var hello func(a int, b int) int = sum // 함수를 저장하는 변수를 선언하고 함수 대입
	world := sum                           // 변수 선언과 동시에 함수를 바로 대입
	fmt.Println(hello(1, 2)) // 3: hello 변수에 저장된 sum 함수 호출
	fmt.Println(world(1, 2)) // 3: world 변수에 저장된 sum 함수 호출


	f := []func(int, int) int{sum, diff} // 함수를 저장할 수 있는 슬라이스를 생성한 뒤 함수로 초기화
	fmt.Println(f[0](1, 2)) // 3: 배열의 첫 번째 요소로 함수 호출
	fmt.Println(f[1](1, 2)) // -1: 배열의 두 번째 요소로 함수 호출


	f1 := map[string]func(int, int) int{ // 함수를 저장할 수 있는 맵을 생성한 뒤 함수로 초기화
		"sum":  sum,
		"diff": diff,
	}
	fmt.Println(f1["sum"](1, 2))  // 3: 맵에 sum 키를 지정하여 함수 호출
	fmt.Println(f1["diff"](1, 2)) // -1: 맵에 diff 키를 지정하여 함수 호출


	f2 := make(map[string]func(int, int) int)
	f2["sum"] = sum
	f2["diff"] = diff

	fmt.Printf("%v\n",f2["sum"](1,2))
	fmt.Printf("%v",f2["diff"](1,2))
}
