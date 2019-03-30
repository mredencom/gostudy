package main

import "fmt"
/**
	插入排序的思想就是：
		从数组的下标为0的元素出发，
		每次向后取一个元素，将该元素
        插入到前面已排好的子数组中，
        以排升序为例，将所要插入的元素
        插在左边小于该元素和右边大于该元
		素之间的位置，进而形成新的子数组，
        直到所有元素全插进来为止。
2.插入排序 时间复杂度：O(n^2)，空间复杂度：1
 */
func InsertionSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}
func main() {
	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3,232}
	fmt.Println(s)
	InsertionSort(s)
	fmt.Println(s)
}
