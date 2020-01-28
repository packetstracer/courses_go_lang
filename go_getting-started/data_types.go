package main

import (
	"fmt"
)

// constant block
const (
	example1 = iota
	example2 = iota
	example3
	example4 = iota
	example5 = iota * 2
	example6 = 2 + 3 + iota
)

// another iota constant block
const (
	example7 = iota
	example8
	example9 = iota
)

func main() {
	// basic data types
	var i int
	i = -11121
	fmt.Println(i)

	var f float32 = 3.1245
	fmt.Println(f)

	s := "gamjah!"
	fmt.Println(s)

	b := false
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	x, y, z := 3, 4, 5
	fmt.Println(x, y, z)

	// pointers
	var pointer *string = new(string)
	*pointer = "Mike"
	fmt.Println(*pointer)

	name := "Peter"
	otherName := &name
	fmt.Println(name, &name)
	fmt.Println(otherName, *otherName)
	name = "Tricia"
	fmt.Println(otherName, *otherName)

	// constants
	const pi = 3.141596
	fmt.Println(pi)
	// pi = 2 // cannot reassing a constant value

	const pic int = 3
	fmt.Println(pi + float32(pic))

	// constants with iota (check de constant block at the top of the file)
	fmt.Println(example1, example2, example3, example4, example5, example6, example7, example8)

}
