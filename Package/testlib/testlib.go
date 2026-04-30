package testlib

import "fmt"

var pop map[string]string

func init() {
	fmt.Println("[testlib] init() 실행: map 초기화 완료")
	pop = make(map[string]string)
	pop["status"] = "OK"
}

func PrintPop() {
	fmt.Println("[testlib] PrintPop() 호출됨, pop 데이터:", pop)
}

func internalCheck() {
	fmt.Println("이 함수는 외부에서 부를 수 없습니다.")
}
