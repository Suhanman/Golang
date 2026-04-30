package main

import (
	"fmt"
	"strings"
)

// 함수 호출의 경우 다른 언어와 동일하고 몇가지의 특이사항이 존재합니다.
// slice 파라미터의 호출
func divideStrings(value string) (string, string) {
	values := strings.Split(value, "_")
	return values[0], values[1]
}

// 다중값의 반환
func addTen(value int) int {
	return value + 10
}

func changeValue(message *string) {
	*message = fmt.Sprintf("your name is %s", *message)
}

// 익명 함수
func sum(values ...int) int {
	result := 0
	for _, value := range values {
		result = result + value
	}

	return result
}

// 일급 함수

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	fmt.Println(addTen(100))
	fmt.Println(sum(1, 2, 3))
	prefix, suffix := divideStrings("Hello_World")
	fmt.Println(prefix, suffix)
	message := "jax"
	changeValue(&message)
	fmt.Println(message)
	sum := func(x int, y int) int {
		return x + y
	}
	fmt.Println(sum(5, 10))

	add := func(i int, j int) int {
		return i + j
	}

	minus := func(i int, j int) int {
		return i - j
	}

	fmt.Println(calc(add, 1, 10))
	fmt.Println(calc(minus, 10, 1))

}
