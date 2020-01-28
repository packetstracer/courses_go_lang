package main

type HTTPRequest struct {
	Method string
}

func main() {
	// for 0..n
	println("LOOP i")
	i := 0
	for i < 5 {
		println(i)
		i++

		if i == 3 {
			println("i equals 3 continue")
			continue
		} else if i == 4 {
			println("i equals 4 break")
			break
		}

		println("next iteration")
	}

	println("\n\nLOOP j")
	for j := 0; j < 5; j++ {
		println(j)
	}

	println("\n\nLOOP k")
	// infinites loops
	k := 0
	for {
		println(k)
		k++

		if k == 5 {
			break
		}

	}

	println("\n\nLOOP Colecciones")
	// looping over collections
	slice := []int{2, 4, 10}
	for i, v := range slice {
		println(i, v)
	}

	myMap := map[string]int{"uno": 3, "dos": 6, "tres": 9}
	for k, v := range myMap {
		println(k, v)
	}

	myMap2 := map[string]int{"aaa": 13, "bbb": 16, "ccc": 19}
	for _, v := range myMap2 {
		println(v)
	}

	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("GET Method")
	case "POST":
		println("POST Method")
	case "DELETE":
		println("POST DELETE")
	default:
		println("Unhandled method")
	}
}
