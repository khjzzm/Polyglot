package main

import "fmt"

func main() {
	rawLiteral := `아리랑\n
		아리랑\n
 		 아라리요`

	interLiteral := "아리랑\n아리랑" +
		"아라리요"

		fmt.Println(rawLiteral)
		fmt.Println(interLiteral)

}