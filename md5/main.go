package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	text := "abcd1234"

	t1 := GetMD5(text)
	fmt.Println(t1)

	t2 := GenMD5(text)
	fmt.Println(t2)

	t3 := HashMD5(text)
	fmt.Println(t3)

}

func GetMD5(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenMD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func HashMD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
