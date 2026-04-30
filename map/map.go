package main

import "fmt"

// map 은 변수 선언으로 생성하거나, make 함수로 생성가능하다. 맵의 정보를 가져올 시 키값이 존재하는 지 확인할 수 있다.
func main() {
	//기본 map 의 선언
	intMap := map[int]string{}
	intMap[100] = "A"
	intMap[101] = "B"
	fmt.Println("%v+\n", intMap)
	fmt.Println(intMap[100])

	// make로 Map의 선언
	var stringMap = make(map[string]string)
	stringMap["A"] = "Hello"
	stringMap["B"] = "World"
	fmt.Println("%v+\n", intMap)
	fmt.Println(stringMap["B"])

	// 선언과 동시에 초기화
	tickers := map[string]string{
		"GooG": "Google",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
	}

	fmt.Printf("%+v\n", tickers)
	fmt.Println(tickers["GooG"])

	val, exists := tickers["FB"]
	fmt.Println(val, exists)

	val, exists = tickers["NON_EXISTING"]
	fmt.Println(val, exists)

	for key, value := range tickers {
		fmt.Println(key, value)
	}
}
