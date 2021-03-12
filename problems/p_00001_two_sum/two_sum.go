package p_00001_two_sum

func twoSumS1(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func twoSumS2(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumS3(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		if _, ok := hashTable[need]; ok {
			return []int{hashTable[need], i}
		} else {
			hashTable[nums[i]] = i
		}
	}
	return nil
}