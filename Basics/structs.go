package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	fmt.Println(a.X + a.Y)
	fmt.Println(Vertex{1, 2})
}
