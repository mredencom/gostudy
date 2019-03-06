package main

import "fmt"

/**
快速排序
*/

func QuickSort(nums []int, start, end int) {
	//fmt.Print(nums)
	//os.Exit(0)
	if start < end {
		i, j := start, end
		//fmt.Print((start+end)/2)
		mid := nums[(start+end)/2]
		for i < +j {
			for nums[i] < mid {
				i++
			}
			for nums[j] < mid {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		if start < j {
			QuickSort(nums, start, j)
		}
		if end > i {
			QuickSort(nums, i, end)
		}
	}
	//fmt.Print(nums)
}
func main() {
	var nums = []int{2, 34, 57, 12, 4, 64, 12431}
	QuickSort(nums, 2, len(nums)-1)
	fmt.Println(nums)
}
