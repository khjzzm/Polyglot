package main

import "fmt"

func main() {

	slice_a := []int{32, 29, 78, 16, 81} // 슬라이스를 생성하면서 값을 초기화
	slice_b := []int{
		32,
		29,
		78,
		16,
		81,  // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}
	var s = make([]int, 5, 10) // 타입, 길이, 용량
	fmt.Println(slice_a)
	fmt.Println(slice_b)
	fmt.Println(cap(s))		// 용량
	fmt.Println(len(s))		// 길이


	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...) // 슬라이스 a에 슬라이스 b를 붙일 때는 b...을 씀
	fmt.Println(a) // [1 2 3 4 5 6]


	//배열은 복사
	arr := [3]int{1, 2, 3}
	var brr [3]int

	brr = arr          // 배열의 요소가 모두 복사됨
	brr[0] = 9       // b[0]에 9를 대입하면 b의 첫 번째 요소만 바뀜

	fmt.Println(arr) // [1 2 3]
	fmt.Println(brr) // [9 2 3]

	//슬라이스는 참조
	slicea := []int{1, 2, 3}
	var sliceb []int    // 슬라이스로 선언

	sliceb = slicea          // a를 b에 대입해도 요소가 모두 복사되지 않고 참조만 함
	sliceb[0] = 9       // 슬라이스는 참조이므로 a[0], b[0]의 값이 모두 바뀜

	fmt.Println(slicea) // [9 2 3]
	fmt.Println(sliceb) // [9 2 3]


}
