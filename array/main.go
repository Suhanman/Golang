package main

import (
	"fmt"
)

// 배열은 길이가 고정된 채로 선언되며, 다차원 배열을 이용할 수 있다.
func main() {

	// 배열 선언
	var intArray [3]int
	intArray[0] = 3
	intArray[1] = 4
	intArray[2] = 6

	fmt.Println(intArray[0] + intArray[1] + intArray[2])

	// 선언과 동시에 초기화
	var stringArrayWithInit = [3]string{"A", "B", "C"}
	for index := range stringArrayWithInit {
		fmt.Printf("%s\n", stringArrayWithInit[index])
	}

	// 멀티플 배열
	var intMultipleArray = [2][3]int{
		{3, 4, 5},
		{6, 7, 8},
	}

	for x := range intMultipleArray {
		for y := range intMultipleArray[x] {
			fmt.Println(intMultipleArray[x][y])
		}
	}

}
