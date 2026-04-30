package main

import (
	"fmt"
)

/* Go에서의 변수 선언법:
1. var 를 이용한 명시적인 선언
변수의 타입을 지정하여 선언
타입을 지정하지 않을수도
var()를 이용하여 다수의 변수 동시 선언이 가능함

2. := 을 이용한 묵시적인 선언
생성과 동시에 초기화한다.
*/

func main() {

	var i1 int = 10
	var s1 string = "hello"

	fmt.Println(i1, s1)

	// 타입 생략가능!
	var i2 = 10
	var s2 = "hello"

	fmt.Println(i2, s2)

	i3 := 10
	s3 := "string"

	fmt.Println(i3, s3)

	var i4, j4, k4 int = 1, 2, 3
	s4, s5, s6 := "string1", "string2", "string3"

	fmt.Println(i4, j4, k4)
	fmt.Println(s4, s5, s6)

	printConsts()

}

const i = 1
const s string = "STRING"

const (
	A = iota
	B
	C
	D
)

func printConsts() {
	fmt.Println(i)
	fmt.Println(s)

	fmt.Println(A, B, C, D)
}
