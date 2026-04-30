package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(15)

	fmt.Println(i)

	if i < 10 {
		fmt.Println("A")
	} else if i == 10 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// if condition 문을 통해서 err 처리를 한 라인으로 할 수있음
	str := "Hello World!"
	if err := ioutil.WriteFile(str, []byte(str), 0644); err != nil {
		fmt.Println(err)
	}

	switch i {
	case 0, 1:

		fmt.Println("A")

	case 2, 3, 4:
		fmt.Println("B")

	case 5:
		fmt.Println("C")

	default:
		fmt.Println("D")
	}
}
