package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	SECRET_KEY = "mysecretKey"
)

type macManager struct {
}

func (m macManager) encryptUrl(api string, now int64) string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(now))

	hmac_sha1 := hmac.New(sha1.New, []byte("secret key"))
	hmac_sha1.Write(append([]byte(api)[:], b[:]...))
	mac := hmac_sha1.Sum(nil)

	return fmt.Sprintf("%s?mac=%x#%d", api, mac, now)
}

func (m macManager) check(encryptApi string) bool {
	u, _ := url.Parse(encryptApi)
	apis := strings.Split(encryptApi, "?")
	encryptTime, _ := strconv.ParseInt(u.Fragment, 10, 64)
	if encryptApi == m.encryptUrl(apis[0], encryptTime) {
		return true
	}
	return false
}

func main() {
	MAC := macManager{}
	now := time.Now().Unix()
	macUrl := MAC.encryptUrl("https://www.joinc.co.kr/api/user/profile", now)
	fmt.Println("Request url : ", macUrl)
	if MAC.check(macUrl) {
		fmt.Println("Test OK")
	} else {
		fmt.Println("Test Fail")
	}
}

