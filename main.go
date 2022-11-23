package main

import (
	"cursogoorientado/encapsulamiento/course"
	"fmt"
)

func main() {

	Go := course.New("Go desde cero", 12.34, false)

	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introduccion",
		2: "Estructuras",
		3: "Maps",
	})

	Go.SetName("Programacion Orientada Objetos con Go")
	fmt.Println(Go.Name())
	Go.PrintClasses()

}
