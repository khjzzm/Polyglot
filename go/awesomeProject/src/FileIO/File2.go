package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	//func Open(name string) (file *File, err error): 기존 파일을 열기
	fi, err := os.Open("hello1.txt")
	if err != (nil) {
		panic(err)
	}
	//func (f *File) Close() error: 열린 파일을 닫음
	defer fi.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if (cnt == 0) {
			break
		}
		fmt.Print(cnt)
	}
}
