package main

import "fmt"

/**
	它的核心思想就是选取一个基准元素(通常已需要排序的数组第一个数)，
	然后通过一趟排序将比基准数大的放在右边，
	比基准数小的放在左边，接着对划分好的两个数组再进行上述的排序。
	快速排序：时间复杂度nlogN,最坏是n平方
*/

func QuickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			//if i <= j {
			//
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
			//}
		}
		if start < j {
			QuickSort(arr, start, j)
		}
		if end > i {
			QuickSort(arr, i, end)
		}
	}
}
func main() {
	var nums = []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
