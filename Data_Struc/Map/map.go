package main

import (
	"fmt"
)

// 해쉬 충돌 : 하나의 인덱스에  리스트로 집어 넣어서 해결한다
// 해쉬는 file-checksum 으로 많이쓴다.
type Product struct {
	Name  string
	Price int
}

const M = 10

func hash(d int) int {
	return d % M
}

// map 은 index로는 안된다.
func main() {
	m := make(map[string]string)
	m2 := make(map[int]Product)
	m3 := [M]string{}

	m3[hash(23)] = "송하나"
	m3[hash(259)] = "백두산"

	fmt.Printf("%d = %s\n", 23, m3[hash(23)])

	m2[16] = Product{"볼펜", 500}
	m2[46] = Product{"지우개", 200}
	m2[78] = Product{"자", 2100}
	m2[178] = Product{"샤프", 3100}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "청주시시 상당구"

	fmt.Printf("송하나의 주소는 %s입니다.\n", m["송하나"])

}

// HashMap : 정렬 x
// sortedMap : 정렬 보장
// 삭제 delete(map,key)
// v, ok :=m[3] 존재여부
