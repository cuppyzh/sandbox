package main

import "fmt"

func main() {
	// mergeTest()
	// removeElementTest()
	// removeDuplicatesTest()
	majorityElementTest()
}

func majorityElementTest() {
	// fmt.Println(float64(2)/3)
	// fmt.Println(float64(1)/3)
    fmt.Println(majorityElement([]int {1,1,2}))
    fmt.Println(majorityElement([]int {2,1,2}))
    fmt.Println(majorityElement([]int {2,2,1,1,1,2,2}))
    fmt.Println(majorityElement([]int {2,1}))
    fmt.Println(majorityElement([]int {1,1}))
    fmt.Println(majorityElement([]int {1,1,2,2}))
    fmt.Println(majorityElement([]int {1}))
    fmt.Println(majorityElement([]int {6,5,5}))
}

func majorityElement(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	candidate, candidateIndex, vote := nums[0], 0, 0
	index := 0

	for candidateIndex < len(nums) {
		if index == len(nums) {
			if vote >= 1 {
				break
			} else {
				candidateIndex++
				if candidateIndex == len(nums) {
					break
				}
				index = 0
				vote = 0
				candidate = nums[candidateIndex]
			}
		} else if candidateIndex == len(nums) {
			break
		}

		if candidate == nums[index] {
			vote++
		} else {
			vote--
		}

		index++
	}

	return candidate
} 

func removeDuplicatesTest(){
    fmt.Println(removeDuplicates([]int {1,1,2}))
    fmt.Println(removeDuplicates([]int {0,0,1,1,1,2,2,3,3,4}))
    fmt.Println(removeDuplicates([]int {0}))
    fmt.Println(removeDuplicates([]int {0,0}))
    fmt.Println(removeDuplicates([]int {0,1}))
    fmt.Println(removeDuplicates([]int {1,1}))
}

func removeDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return 1
		} 

		return 2
	}

	var ptr1, ptr2 = 0, 1;

	for ptr2 < len(nums) {
		// fmt.Println(nums,ptr1,ptr2)
		if nums[ptr1] != nums[ptr2] {
			ptr1++
			nums[ptr1] = nums[ptr2]
		}

		ptr2++
	}
	
	// fmt.Println(ptr1)
	return ptr1+1
}

func removeElementTest(){
	fmt.Println(removeElement([]int {3,2,2,3}, 3))
	fmt.Println(removeElement([]int {0,1,2,2,3,0,4,2}, 2))
	fmt.Println(removeElement([]int {1}, 1))
	fmt.Println(removeElement([]int {3,3}, 3))

}

func removeElement(nums []int, val int) int {
	var ptr1, ptr2 int = 0, len(nums)-1

	for ptr2 >= ptr1{

		if nums[ptr1] == val {
			nums[ptr1] = nums[ptr2]
			ptr2--
			continue
		}

		ptr1++
	}

	return ptr1
}

func mergeTest() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{2, 0}, 1, []int{1}, 1)
	merge([]int{0}, 0, []int{1}, 1)
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	merge([]int{0, 0, 0, 0, 0, 0}, 0, []int{1, 2, 3, 4, 5, 6}, 6)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var ptr1, ptr2, ptr3 int = m - 1, n - 1, m + n - 1

	for ; ptr1 >= 0 && ptr2 >= 0; ptr3-- {
		if nums2[ptr2] >= nums1[ptr1] {
			nums1[ptr3] = nums2[ptr2]
			ptr2--
		} else {
			nums1[ptr3] = nums1[ptr1]
			ptr1--
		}
	}

	if ptr2 >= 0 {
		copy(nums1[:ptr3+1], nums2[:ptr2+1])
	}

	var result string = ""
	for _, value := range nums1 {
		result += fmt.Sprintf("%d ", value)
	}
	fmt.Println(result)
}
