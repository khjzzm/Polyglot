package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func SumAndDiff(a int, b int) (r0 int, r1 int) {
	return a + b, a - b // 리턴할 값 또는 변수를 차례대로 나열
}

func SumAndDiff2(a int, b int) (sum2 int, diff2 int) { // 리턴값을 각각 sum, diff로 이름을 정함
	sum2 = a + b  // 리턴값 변수 sum에 합 대입
	diff2 = a - b // 리턴값 변수 diff에 차 대입
	return
}

func main() {
	fmt.Printf("%d\n",add(42, 13))

	_, diff := SumAndDiff(6, 2) // 첫 번째 리턴값은 _으로 생략, 두 번째 리턴값만 사용
	fmt.Println(diff)


	_, diff2 := SumAndDiff2(6, 2) // 첫 번째 리턴값은 _으로 생략, 두 번째 리턴값만 사용
	fmt.Println(diff2)


}
