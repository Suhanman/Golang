package main

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}
func main() {

	rect := Rect{width: 10, height: 5}
	area := rect.area()
	area2 := rect.area2()
	println(area)
	println(rect.width, area2)

}
