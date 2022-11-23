package main

func main() {

	Go := &course.Course{
		"Go desde cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()

}
