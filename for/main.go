package main

import "fmt"

type Obj struct {
	Name string
	Age  int
}

// range 문은 array , slice, map 의 내용을 반복하는 것에 많이 사용하는데 range 문은 반복을 처리하기 전에 객체를 복사하여서 사용합니다.
// 따라서 반복문 안에서 객체의 값을 변경해도 실제 객체의 값은 병경되지 않습니다. 따라서 실제 객체의 값을 변경하고자 할때에는 포인터를 이용하여서 작업을 처리해야합니다.
func PrintObject(list []Obj) {
	for index, object := range list {
		fmt.Printf("index: %d, object: %v\n", index, object)
	}
}

func main() {

	for index := 1; index <= 10; index++ {
		fmt.Printf("index: %d\n", index)
	}

	for index := 1; index <= 10; {
		fmt.Printf("index: %d\n", index)
		index++
	}

	index := 1
	for index <= 10 {
		fmt.Printf("index: %d\n", index)
		index++
	}

	index = 1
	for true {
		fmt.Printf("index: %d\n", index)
		index++

		if index > 10 {
			break
		}
	}

	strArray := []string{"A", "B", "C"}
	for index, str := range strArray {
		fmt.Printf("index: %d, str: %s\n", index, str)
	}

	dictionary := map[string]string{
		"key_A": "value_A",
		"key_B": "value_B",
	}

	for key, value := range dictionary {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	list := []Obj{
		{"Beckham", 11},
		{"Zidane", 7},
		{"Ronaldo", 9},
	}

	for _, object := range list {
		object.Age = object.Age * 2
	}

	PrintObject(list)

	for index := range list {
		object := &list[index]
		object.Age = object.Age * 2
	}
	PrintObject(list)
}
