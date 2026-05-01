package main

import (
	"fmt"
	"time"
)

/*
채널은 데이터를 주고 받을 수 있는 통로다. make()함수로 미리 생성되고 <- 데이터를 보내고. 받을 수 있다.
주로 Go 루틴에서 함수들 사이에서 데이터를 주고 받는 것에 이용된다. 데이터를 받을때 까지 대기하며, 동기화를 구현하는 것에 이용한다.
*/

func sendChan(ch chan<- string) {
	ch <- "Data"
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func main() {
	// 정수형 채널을 생성한다.
	ch := make(chan int)

	go func() {
		ch <- 123 //채널에 123을 보낸다.
	}()

	var i int
	i = <-ch // 채널로부터 123을 받는다.
	println(i)

	/*
		기본적으로 채널은 버퍼를 사용하지 않고 생성됩니다.
		버퍼를 이용하지 않으면 채널의 데이터를 바로 수신하지 않을 시 에 오류가 발생하게 됩니다.
		make() 함수에 인자를 추가하여 버퍼를 생성할 수 있습니다.
		버퍼를 이용할 때는 채널에 데이터를 수신후 다른 Go 루틴에서 데이터를 수신하지 않아도 다른 작업을 수행할 수 있습니다.
	*/

	//c := make(chan int)
	//c <- 1
	//fmt.Println(<-c)

	ch1 := make(chan int, 1)
	// 수신자가 없더라도 보낼 수 있다
	ch1 <- 101
	fmt.Println(<-ch1)

	// 그니까 Go 내장 RabbitMQ라는거죠...?

	ch2 := make(chan string, 1)
	sendChan(ch2)
	receiveChan(ch2)

	close(ch2)
	println(<-ch2)
	println(<-ch2)

	// 채널이 닫힌 것을 감지할 때까지 계속 수신
	if _, success := <-ch2; !success {
		println("더이상의 데이터 없음")
	}

	// 위 표현과 동일한 채널의 range문
	for i := range ch2 {
		println(i)
	}

	// select : 키워드를 이용하여 여러 개의 채널에서 수신한 데이터를 이용해서 작업을 진행하는 방법제공
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")

		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}
