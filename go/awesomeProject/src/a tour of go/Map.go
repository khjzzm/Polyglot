package main

import "fmt"

type Member struct {
	number int
	job string
}

var m map[string]Member


func main() {
	a := map[string]int{"Hello": 10, "world": 20}

	b := map[string]int{
		"Hello": 10,
		"world": 20, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}

	c := make(map[string]string)
	c["name"] = "kimhyunjin"
	c["age"] = "28"


	fmt.Println(a["Hello"]) // 10
	fmt.Println(b["world"]) // 10

	fmt.Printf("삭제전 %v\n", c)

	delete(c, "name")
	fmt.Printf("삭제후 %v\n", c)




	m = make(map[string]Member)
	m["kimzzang"] = Member{
		910509,
		"progreamer",
	}

	m["kimzzang2"] = Member{
		19910509,
		"progreamer2",
	}


	fmt.Printf("%v\n", m)

	solarSystem := make(map[string]float32) // 키는 string, 값은 float32인 맵 생성 및 공간 할당

	solarSystem["Mercury"] = 87.969 // 맵[키] = 값
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528

	fmt.Println(solarSystem["Earth"]) // 365.25641
	fmt.Println(solarSystem["Pluto"])

}
