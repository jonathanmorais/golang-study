package main

import "fmt"

type peter interface { // pet
	speaker()
}

type canino struct {
	late string
}

type felino struct {
	mia string
}

func (c canino) speaker() string {
	return c.late
}

func (f felino) speaker() string {
	return f.mia
}

func main() {
	cachorro := canino{late: "auau"}
	gato := felino{mia: "miau"}
	fmt.Println(cachorro)
	fmt.Println(gato)

}
