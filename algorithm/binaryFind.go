package main

import "fmt"

/**
	二分查找：首先了解二分查找，首先在长度为n的表范围中查找，
	第一次循环在n/2中查找，第二次在n/2/2中查找，依次循环。
	假设在第X次找到，那么 就是找2的X次方次，有2的X次方=n解出x为log2的n ，
	故时间复杂度为log2N。由于辅助空间是常数级别的所以：空间复杂度是O(1);
	思路：先找到中间的下标middle = (leftIndex + RightIndex) /2 ,然后让中间的下标值和FindVal比较
	a:如果arr[middle] > FindVal,那么就向LeftIndex~(midlle - 1)区间找
	b:如果arr[middle] < FindVal,那么就向middle + 1 ~RightIndex区间找
	c:如果arr[middle] == FindVal,那么直接返回
	②从①的a、b、c递归执行，知道找到位置
	③如果LeftIndex > RightIndex，则表示找不到，退出
 */
/**
   参数1:待查的数组
   参数2:左边的下标
   参数3:右边的下标
   参数4:待查的元素
 */
func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	//首先需要判断左边的leftIndex是否大于rightIndex值
	if leftIndex > rightIndex {
		fmt.Println("没有找到数据")
		return
	}
	//然后我们需要找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标是%v\n", middle)
	}
}
func main() {
	//arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	arr := []int{1, 2, 5, 4, 5}
	fmt.Println(len(arr))
	BinaryFind(&arr, 0, len(arr)-1, 22)
}
