package main

import "fmt"

func main() {
	p := []int{1,2,3,4,5,6}
	fmt.Println("p==",p)

	for i := 0; i <len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}


	fmt.Println(len(p), cap(p)) // 6 6: 길이가 5이며 용량이 5인 슬라이스

	p = append(p, 6, 7)         // 슬라이스 a에 값 6, 7을 추가
	fmt.Println(len(p), cap(p)) // 7 10: 길이가 7이며 용량이 10인 슬라이스, 용량이 늘어남!

	for i := 0; i <len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	for _, value := range p {
		fmt.Printf("%v",value)
	}
}
