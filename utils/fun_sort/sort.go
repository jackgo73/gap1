package fun_sort

import "fmt"

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Insertion sort
func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func BubbleSort(data []int) {
	fmt.Println("BubbleSort")
	fmt.Printf("Source Data:%d\n", data)

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}

		fmt.Printf("Middle Data:%2d=>%d\n", i, data)
	}

	fmt.Printf("Result Data:%d\n", data)
}

func InsertSort(data []int) {
	fmt.Println("InsertSort")
	fmt.Printf("Source Data:%d\n", data)

	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			for j := i; j > 0; j-- {
				if data[j] < data[j-1] {
					data[j-1], data[j] = data[j], data[j-1]
				}
			}
		}
		fmt.Printf("Middle Data:%2d=>%d\n", i, data)
	}

	fmt.Printf("Result Data:%d\n", data)
}

func SelectSort(data []int) {
	fmt.Println("SelectSort")
	fmt.Printf("Source Data:%d\n", data)
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[j], data[i] = data[i], data[j]
			}
		}
		fmt.Printf("Middle Data:%2d=>%d\n", i, data)
	}
	fmt.Printf("Result Data:%d\n", data)
}
