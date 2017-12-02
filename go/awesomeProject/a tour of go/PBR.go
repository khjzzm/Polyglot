package main
func main() {
	msg := "Hello"

	sayValue(msg)
	println(msg)

	sayReference(&msg)
	println(msg) //변경된 메시지 출력
}

func sayReference(msg *string) {
	println(*msg)
	*msg = "Changed" //메시지 변경
}

func sayValue(msg string)  {
	println(msg)
	msg = "Changed"
}