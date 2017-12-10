package main

func main() {
	sum := func(n ...int) (s int){
		s = 0
		for _, i := range n{
			s += i
		}
		return
	}
	result := sum(1,2,3,4,5)
	print(result)

}