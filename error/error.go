package main

import (
	"fmt"
	"log"
	"os"
)

// 1. 사용자 정의 에러 구조체 정의
type MyError struct {
	Code    int
	Message string
}

// 2. error 인터페이스 구현 (Error() 메서드)
func (e MyError) Error() string {
	return fmt.Sprintf("에러 코드: %d, 메시지: %s", e.Code, e.Message)
}

// 3. 간단한 리스트와 에러를 반환하는 함수
func otherFunc() ([]string, error) {
	// 예시를 위해 리스트 생성
	items := []string{"apple", "banana", "cherry"}

	// 특정 조건에서 사용자 정의 에러 반환 (여기서는 강제 발생)
	return items, MyError{
		Code:    404,
		Message: "데이터를 찾을 수 없습니다.",
	}
}

func main() {
	// 파일 오픈 예시 (경로는 환경에 맞게 수정 필요)
	f, err := os.Open("test.txt")
	if err != nil {
		// 실제 파일이 없으면 여기서 멈추므로, 테스트를 위해 출력만 함
		log.Println("파일 오픈 에러:", err)
	} else {
		fmt.Println("파일명:", f.Name())
		f.Close()
	}

	fmt.Println("--- otherFunc 호출 ---")

	// 4. 사용자 정의 에러 처리
	list, err := otherFunc()

	if err == nil {
		fmt.Println("결과 리스트:", list)
	} else {
		// Type Assertion을 이용한 에러 분기
		switch e := err.(type) {
		case MyError:
			log.Printf("[커스텀 에러 발생] %v\n", e)
		case error:
			log.Fatal("일반 에러:", e.Error())
		default:
			println("에러 없음")
		}
	}
}
