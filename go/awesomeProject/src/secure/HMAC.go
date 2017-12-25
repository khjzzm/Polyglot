package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)
const (
	//salt
	SECRET_KEY = "mysecretKey"
)
func main() {
	input := "Hello. My name is A"
	hmac_md5 := hmac.New(md5.New, []byte(SECRET_KEY))
	hmac_md5.Write([]byte(input))
	hmac_sha1 := hmac.New(sha1.New, []byte(SECRET_KEY))
	hmac_sha1.Write([]byte(input))
	hmac_sha256 := hmac.New(sha256.New, []byte(SECRET_KEY))
	hmac_sha256.Write([]byte(input))

	fmt.Printf("md5:\t%x\n", hmac_md5.Sum(nil))
	fmt.Printf("sha1:\t%x\n", hmac_sha1.Sum(nil))
	fmt.Printf("sha256:\t%x\n", hmac_sha256.Sum(nil))
}
