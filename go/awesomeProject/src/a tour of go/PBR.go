package main
func main() {
	msg := "Hello"

	sayValue(msg)
	println("change sayValue : "+msg)

	sayReference(&msg)
	println("change sayReference : "+msg) //변경된 메시지 출력
}

func sayValue(msg string)  {
	println("sayValue : "+msg)
	msg = "Changed"
}

func sayReference(msg *string) {
	println("sayReference : " + *msg)
	*msg = "Changed" //메시지 변경
}

