package fun_sort

import (
	"fmt"
	"math/rand"
	"time"
)

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

func quick_sort(nums []int, l, r int) {
	if l >= r {
		return
	}
	rand.Seed(time.Now().Unix())
	p := rand.Intn(r-l+1) + l
	nums[r], nums[p] = nums[p], nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] < nums[r] {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	quick_sort(nums, l, i-1)
	quick_sort(nums, i+1, r)
}


func heap_sort(nums []int) []int {
	lens := len(nums) - 1
	for i := lens/2; i >= 0; i -- { // 建堆 O(n) lens/2后面都是叶子节点，不需要向下调整
		down(nums, i, lens)
	}
	for j := lens; j >= 1; j -- {//堆排序（升序）:堆顶(最大值)交换到末尾
		nums[0], nums[j] = nums[j], nums[0]
		lens --
		down(nums, 0, lens)
	}
	return nums
}
func down(nums []int, i, lens int) {//O(logn)大根堆，如果堆顶节点小于叶子，向下调整
	max := i
	if i<<1+1 <= lens && nums[i<<1+1] > nums[max] {
		max = i<<1+1
	}
	if i<<1+2 <= lens && nums[i<<1+2] > nums[max] {
		max = i<<1 + 2
	}
	if max != i {
		nums[max], nums[i] = nums[i], nums[max]
		down(nums, max, lens)
	}
}

func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n - 1; i++ {
		exchange := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] { //两两比较，前面比后面大
				nums[j], nums[j+1] = nums[j+1], nums[j] //交换
				exchange = true
			}
		}
		if !exchange {
			return nums
		}
	}
	return nums
}

func sortArray(nums []int) []int {
	cnt := [100001]int{}
	for i := 0; i < len(nums); i++ {
		cnt[nums[i] + 50000]++
	}
	idx := 0
	for i := 0; i < 100001; i++ {
		for cnt[i] > 0 {
			nums[idx] = i - 50000
			idx++
			cnt[i]--
		}
	}
	return nums
}
