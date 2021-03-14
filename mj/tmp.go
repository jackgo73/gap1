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

	d := []int{1,2,3,4,5}
	f := []int{7,8,9,10}
	g := []int{}
	g = append(g, d...)
	g = append(g, f...)
	fmt.Println("f")
	f = append(f, d...)
	fmt.Println(f)
	fmt.Println("g")
	fmt.Println(g)



	return true
}

func sss(a, b [][]int) [][]int {
	r := [][]int{}
	r = append(r, a...)
	r = append(r, b...)



	return r
}
