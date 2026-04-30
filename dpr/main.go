package main

import (
	"fmt"
	"os"
)

/*
defer
함수를 바로 실행하지 않고 종료 시점에 실행한다
try catch finally 문의 finally라고 생각하면 됨

panic
실행 시점에서 현재 함수의 defer() 함수를 실행 시킨 후에 리턴한다.
상위로 계속 리턴하여 에러를 출력하고 종료한다.

recover
panic()에 의한 에러를 복구한다.
*/
func main() {

	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
	n, err := f.Read(bytes)
	println(n) // 실제 읽어온 바이트 수

	openFile("Invalid.txt")

}
func openFile(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR:", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

}
