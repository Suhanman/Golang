package main

import (
	"fmt"
	"strings"
)

func main() {
	// 문자열 합치기에는 strings.Join() 이용
	strs := []string{"a", "b", "c", "d"}
	fmt.Println(strings.Join(strs, " :"))

	// 문자열 대치할때에는 strings.replace() 이용합니다.
	str := "a.b.c"
	r := strings.Replace(str, ".", "_", -1)
	fmt.Println(r)
}
