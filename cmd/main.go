package main

import (
	"fmt"

	"os"
	"strconv"

	"github.com/mspeer383/go-structs"
)

func main() {

	var argHeight float64
	argHeight, _ = strconv.ParseFloat(os.Args[1], 64)

	var argWidth float64
	argWidth, _ = strconv.ParseFloat(os.Args[2], 64)

	var rectangle structs.Rectangle

	//passed in
	rectangle = structs.Rectangle{Height: argHeight, Width: argWidth}

	fmt.Println("my Rectangle ", rectangle)
	fmt.Println("area is ", rectangle.Area())

	//static
	rectangle = structs.Rectangle{Height: 40, Width: 25}

	fmt.Println("my Rectangle ", rectangle)
	fmt.Println("area is ", rectangle.Area())

}
