package main

import "fmt"

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func main() {
	p := person{}

	p.name = "Lee"
	p.age = 18

	fmt.Println(p)

	var p1 person
	p1 = person{"김길둥", 20}
	p2 := person{name: "Scan", age: 50}

	fmt.Println(p1)
	fmt.Println(p2)

	p3 := new(person)
	p3.name = "LEE"

	fmt.Println(p3)

	dic := newDict()
	dic.data[1] = "A"
	fmt.Println(dic)

}
