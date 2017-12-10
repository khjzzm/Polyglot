package main

import "fmt"

type Vertex4 struct {
	X, Y int
}

func main() {
	//new(T) 는 모든 필드가 0(zero value) 이 할당된 T 타입의 포인터를 반환합니다.
	//( zero value 는 숫자 타입에서는 0 , 참조 타입에서는 nil 을 뜻합니다 )

	v := new(Vertex4)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}
