package main

import (
	"fmt"
	"time"
)

// Go 키워드를 이용하여 함수를 호출하면 Go 루틴을 생성한다.

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// 동기 실행
	say("Sync")

	// 비동기 실행
	go say("ASync1")
	go say("ASync2")
	go say("ASync3")

	time.Sleep(time.Second * 3)
}
