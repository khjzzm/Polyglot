package main

import "fmt"

type Vertex3 struct {
	X, Y int
}

var (
	p = Vertex3{1,2} //has type Vertex
	q = &Vertex3{1,2} //has type *Vertex
	r = Vertex3{X:1} // Y:0 is implicit
	s = Vertex3{}
)

func main() {
	fmt.Println(p,q,r,s)
}
