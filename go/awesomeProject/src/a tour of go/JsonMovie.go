package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Movie struct {
	Title string `json:"title,omitempty"`
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Acrots []string
}

const (
	Sunday       = iota // 0
	Monday              // 1
	Tuesday             // 2
	Wednesday           // 3
	Thursday            // 4
	Friday              // 5
	Saturday            // 6
	numberOfDays        // 7
)


var movies = []Movie{
	{Title:"Cool Hand Luke", Year:1945, Color:false, Acrots: []string{"kimhyunjin", "khjzzm"}},
	{Title:"little forest", Year:2018, Color:true, Acrots: []string{"kimsanghun", "kshzzm"}},
	{Year:2018, Color:true, Acrots: []string{"kimsanghun", "kshzzm"}},
}

var a1 [5]int = [5]int{32, 29, 78, 16, 81} // int형이며 길이가 5인 배열을 선언하고 초기화
var b1 = [5]int{32, 29, 78, 16, 81}        // 배열을 선언할 때 자료형과 길이 생략


func main() {


	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b) // 변수 b를 사용할 수 있음
	} else {
		fmt.Println(err) // 변수 err을 사용할 수 있음
	}


	arr1 := [5]int{32, 29, 78, 16, 81} // 배열을 선언할 때 var 키워드, 자료형과 길이 생략

	for _, value := range arr1 { // i에는 인덱스, value에는 배열 요소의 값이 들어감
		fmt.Println(value)
	}


	data, _ := json.Marshal(movies)
	fmt.Printf("%s\n", data)

	data2, _ := json.MarshalIndent(movies, "", " ")
	fmt.Printf("%s\n", data2)

	fmt.Println(Thursday)     // 4
	fmt.Println(numberOfDays) // 7

}
