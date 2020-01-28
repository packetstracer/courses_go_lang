package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 4
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 4}
	fmt.Println(arr2)

	arr3 := [5]int{1, 2, 3, 4, 5}
	slice := arr3[:]
	arr3[1] = 8
	slice[3] = 16
	fmt.Println(arr3, slice)

	s2 := []int{1, 2, 3}
	s2 = append(s2, 7, 8, 9, 100, 200, 300)
	fmt.Println(s2)

	s3 := s2[2:]
	s4 := s2[:4]
	s5 := s2[2:4]
	fmt.Println(s3, s4, s5)

	m := map[string]int{"foo": 24, "bar": 43, "mike": 88}
	fmt.Println(m, m["foo"])
	delete(m, "foo")
	fmt.Println(m)

	type user struct {
		id   int
		name string
		age  int
	}
	var u1 user
	u1.id = 1
	u1.name = "Ivan"
	u1.age = 41
	fmt.Println(u1)

	u2 := user{id: 2, name: "Peter", age: 43}
	fmt.Println(u2)
}
