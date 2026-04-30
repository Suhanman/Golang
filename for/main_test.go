package main

import "testing"

func TestPrintObject(t *testing.T) {
	list := []Obj{{"Beckham", 11},
		{"Zidane", 7},
		{"Ronaldo", 9}}
	PrintObject(list)
}
