package main

import "fmt"

/**
归并排序算法是一款十分高效的算法，因为用到了二叉树的特性，我们知道能使用二叉树特性的排序都比较高效，我们详细讲解一下

比如有两个已经排好的数组（[5,7],[2,6,10]），这个需要我们进行合并排序，我们改如何排呢

①首先取出第一个数组第一个元素：5，第二个数组第一个元素：2，我们得到最小值2，放入到第三个数组c中，此时c是[2]

②因为第一个数组第一个元素还没有被拿走，所以第二步我们取出第一个数组第一个元素5，第二个数组第二个元素6，我们得到最小值5，放入到第三个数组c中，此时c是[2,5]

③我们取第一个数组第二个元素7，第二个数组第二个元素 6，我们得到最小值6，放入到第三个数组c中，此时c是[2,5,6]

④我们取第一个数组第二个元素7，第二个数组第三个元素10，我们得到最小值7，放入到第三个数组c中，此时c是[2,5,6,7]

⑤此时第一个数组已经没了，第二个数组还有值，因为两个数组本身就是遍历的，则可以直接把还有值的那个数组（此处是第二个数组）剩下的元素按顺序追加到c中，此时c是[2,5,6,7,10]

图例:https://img-blog.csdn.net/20180928232520286?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3N0cm9uZ2x5aA==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70
*/

func lastMergeSort(list1 []int, list2 []int) []int {
	//初始化两个值
	i, j := 0, 0
	var temp []int //定义一个临时存数据的slice
	//这个for循环这样用 就相当于php while 使用
	for i < len(list1) && j < len(list2) {
		if list1[i] <= list2[j] {
			//这里是吧小的数据存到temp的slice
			temp = append(temp, list1[i])
			i++
		} else {
			//这里是吧小的数据存到temp的slice
			temp = append(temp, list2[j])
			j++
		}
		if i == len(list1) {
			for ; j < len(list2); j++ {
				temp = append(temp, list2[j])
			}
		}
		if j == len(list2) {
			for ; i < len(list1); i++ {
				temp = append(temp, list1[i])
			}
		}
	}
	return temp
}

func mergeSort(theArray []int) []int {
	//当切片一个数据时就不用排序了
	if len(theArray) == 1 {
		return theArray
	}
	//吧切片一分为二
	mid := len(theArray) / 2
	leftArray := mergeSort(theArray[:mid])
	//fmt.Println("leftArray:", leftArray)
	rightArray := mergeSort(theArray[mid:])
	//fmt.Println("rightArray:", leftArray)
	return lastMergeSort(leftArray, rightArray)
}
func main() {
	var theArray = []int{234243, 6, 4, 5, 1, 8, 7, 2, 3}
	fmt.Print("排序前")
	fmt.Println(theArray)
	fmt.Print("排序后")
	arrayResult := mergeSort(theArray)
	fmt.Println(arrayResult)
}
