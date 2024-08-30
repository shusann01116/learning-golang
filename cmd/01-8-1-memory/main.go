package main

import "fmt"

func main() {
	// allocate actual memory
	s1 := make([]int, 1000)
	fmt.Println(len(s1)) // 1000
	fmt.Println(cap(s1)) // 1000

	s2 := make([]int, 0, 1000)
	fmt.Println(len(s2)) // 0
	fmt.Println(cap(s2)) // 1000

	m := make(map[string]string, 1000)
	fmt.Println(len(m)) // 0
}
