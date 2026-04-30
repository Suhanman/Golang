package main

import (
	"Package/testlib"
	"fmt"
)

/*
실행 시작 시 점으로 인식
공유 라이브러리를 만들 때에는 main 으로 만들면 안됨.

import 키워드를 통하여 다른 패키지의 함수를 불러올 수 있다.

${GoROOT}/pkg 는 표준패키지
${GoPATH}/pkg 는 서드파티패키지

첫문자 대문자시 public
첫문자 소문자시 private
*/

func main() {

	fmt.Println("메인함수시작")

	testlib.PrintPop()

	//fmt.Println(testlib.pop)
	//testlib.internalCheck()

}
