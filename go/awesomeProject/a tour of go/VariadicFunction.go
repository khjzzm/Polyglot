package main

func main() {
	say("This", "is", "a", "book")
	say("This")
	println(sum(1,2,3,4,5,6,7,8,9,10))

	count, total := sum2(1,2,3,4,5,6,7,8,9,10)
	println(count, total)

	count1, total1 := sum3(1,2,3,4,5,6,7,8,9,10)
	println(count1, total1)

}

func say(msg ...string)  {
	for  _,s := range msg{
		println(s)
	}
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums{
		s+=n
	}
	return  s
}

func sum2(nums ...int) (int, int){
	s := 0
	count := 0
	for _, n := range nums{
		s += n
		count++
	}
	return count, s
}

//Named Return Parameter
func sum3(nums ...int) (count int, total int)  {
	for _, n:= range nums{
		total += n
	}
	count = len(nums)
	return
}