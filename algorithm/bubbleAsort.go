package main

import "fmt"

/**
	冒泡排序顾名思义就是整个过程像气泡一样往上升，
	单向冒泡排序的基本思想是（假设由小到大排序）：
	对于给定n个记录，从第一个记录开始依次对相邻的两个记录进行比较，
	当前面的记录大于后面的记录时，交换位置，进行一轮比较和换位后，
	n个记录的最大记录将位于第n位，然后对前（n-1）个记录进行第二轮比较；
	重复该过程，直到记录剩下一个为止。
	冒泡排序 时间复杂度：n^2，空间复杂度：1
*/
func BubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				//把values[i] 移到 values[j]后面  交换值
				values[j], values[i] = values[i], values[j]
			}
		}
	}
	fmt.Println(values)
}
func main() {
	values := []int{400, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(values)
	BubbleSort(values)
}
