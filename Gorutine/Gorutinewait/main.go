package main

import (
	"fmt"
	"runtime"
	"sync"
)

// sync.WaitGroup을 이용하여 다른 작업이 종료되는 것을 대기할 수 있다.

type counter struct {
	i  int64
	mu sync.Mutex
}

// counter의 값을 1씩 증가시킴
func (c *counter) incr() {
	c.mu.Lock()   //i 값을 변경하는 부분(임계영역) 을 뮤텍스로 잠근다.
	c.i += 1      // 공유 데이터의 변경
	c.mu.Unlock() // i 값을 변경 완료한 후에 뮤텍스 잠금의 해제
}

// counter 값의 춝력
func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {

	// 모든 cpu 사용
	// Go 루틴은 기본적으로 하나의 프로세스를 이용하여 시분할로 작업을 처리할 수 있다.
	// 해당 내장함수를 이용하여 여러개의 프로세스로 이용할 수 도 있다.
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}     // 카운터 생성
	wg := sync.WaitGroup{} // WaitGroup 생성

	// c.incr를 실행하는 1000개의 고루틴
	for i := 0; i < 1000; i++ {
		wg.Add(1) //WaitGroup의 고루틴 개수 1 증가
		go func() {
			defer wg.Done() // 고루틴 종료 시 Done()처리
			c.incr()        // 카운터 값의 1증가
		}()
	}

	wg.Wait() // 모든 고루틴이 종료될 때 까지 대기

	c.display() // c의 값이 출력된다.

}
