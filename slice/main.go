package main

import (
	"fmt"
	"sort"
)

type myDataType struct {
	name string
	age  int
}

// 슬라이스는 배열과 비슷하지만 몇가지의 장점이 있습니다.
// 크기를 지정하지 않고 생성이 가능함
// 크기를 동적으로 증가시킬 수 있음
// 부분 배열의 추출이 가능함
func main() {

	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)

	s1 := make([]int, 0)
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 5)
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 3, 5)
	fmt.Println(s3, len(s3), cap(s3))

	//append
	//append 는 len 으로 설정한 값 뒤에 축가됨
	intArray := []int{100, 101, 102}
	s1 = append(s1, 100)
	s2 = append(s2, intArray...)
	s3 = append(s3, 100, 101, 102)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	//copy
	//copy를 위해서 생성한 슬라이스는 기존 스라이스와 크기가 동일해야 한다.
	lettersCopy1 := make([]string, 0)
	copy(lettersCopy1, letters)
	fmt.Println(lettersCopy1)

	lettersCopy2 := make([]string, len(letters), len(letters))
	copy(lettersCopy2, letters)
	fmt.Println(lettersCopy2)

	lettersCopy2[3] = "="
	fmt.Println(letters)
	fmt.Println(lettersCopy2)

	// 삭제, 추출

	integers := []int{1, 2, 3, 4, 5}
	sub1 := integers[1:4]
	sub2 := integers[2:4]
	fmt.Println(integers, len(integers), cap(integers))
	fmt.Println(sub1, len(sub1), cap(sub1))
	fmt.Println(sub2, len(sub2), cap(sub2))

	sub1[2] = 100
	fmt.Println(integers, len(integers), cap(integers))
	fmt.Println(sub1, len(sub1), cap(sub1))
	fmt.Println(sub2, len(sub2), cap(sub2))

	// 슬라이스의 정렬

	mySlice := make([]myDataType, 0)
	mySlice = append(mySlice, myDataType{"김형준", 42})
	mySlice = append(mySlice, myDataType{"홍길동", 28})
	mySlice = append(mySlice, myDataType{"임꺽정", 38})

	fmt.Println(mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].age < mySlice[j].age
	})

	fmt.Println(mySlice)

}
