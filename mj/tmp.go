package mj

import "fmt"

func test() bool {
	fmt.Println("a")

	r := [][]int{}
	a := [][]int{{1, 2}, {3, 4}}
	b := [][]int{{4, 7}, {9, 14}}
	r = append(r, a...)
	r = append(r, b...)

	fmt.Println("r")
	fmt.Println(r)


	k := sss(a,b)
	fmt.Println("k")
	fmt.Println(k)
	return true
}

func sss(a, b [][]int) [][]int {
	r := [][]int{}
	r = append(r, a...)
	r = append(r, b...)



	return r
}
