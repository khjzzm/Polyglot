package main

import "fmt"

type Vertex2 struct {
	X int
	Y int
}


func main(){
	p := Vertex2{1,2}
	q := &p
	q.X = 5e4
	fmt.Println(p)
}


