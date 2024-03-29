package main

import "fmt"

/**
1.梳排序算法讲解
	嗯，历史山这些算法的名字总是，嗯，不知道如何说，反正充满了想象力，不过也恰恰形象的说明了这个算法

	梳排序，跟梳子一样，齿和齿中间有间隙，这个间隙是多少呢，是1.3，嗯，记住就好，为啥是1.3，数学功底比较强的人可以去看看哈，我们这就不做太多讨论，直接举例

	有数组[6, 4, 5, 1, 8, 7, 2, 3]

	我们先来算算间隙，上面的数组长度为8，8/1.3=6，则间隙为6

	我们开始遍历数组

第一遍，取间隙6

           取下标0和6的数，分别是6和2，6比2大，进行互换，[2, 4, 5, 1, 8, 7, 6, 3]

           取下标1和7的数，分别是4和3，4比3大，进行互换，[2, 3, 5, 1, 8, 7, 6, 4]

第二遍，取间隙6/1.3=4

           取下标0和4的数，分别是2和8，顺序正确，则不互换，[2, 3, 5, 1, 8, 7, 6, 4]

           取下标1和5的数，分别是3和7，顺序正确，则不互换，[2, 3, 5, 1, 8, 7, 6, 4]

           取下标2和6的数，分别是5和6，顺序正确，则不互换，[2, 3, 5, 1, 8, 7, 6, 4]

           取下标3和7的数，分别是1和4，顺序正确，则不互换，[2, 3, 5, 1, 8, 7, 6, 4]

第三遍，取间隙4/1.3=3

           取下标0和3的数，分别是2和1，2比1大，进行互换，[1, 3, 5, 2, 8, 7, 6, 4]

           取下标1和4的数，分别是3和8，顺序正确，则不互换，[1, 3, 5, 2, 8, 7, 6, 4]

           取下标2和5的数，分别是5和7，顺序正确，则不互换，[1, 3, 5, 2, 8, 7, 6, 4]

           取下标3和6的数，分别是2和6，顺序正确，则不互换，[1, 3, 5, 2, 8, 7, 6, 4]

           取下标4和7的数，分别是8和4，8比4大，进行互换，[1, 3, 5, 2, 4, 7, 6, 8]

第四遍，取间隙3/1.3=2

           取下标0和2的数，分别是1和5，顺序正确，则不互换，[1, 3, 5, 2, 4, 7, 6, 8]

           取下标1和3的数，分别是3和2，3比2大，进行互换，[1, 2, 5, 3, 4, 7, 6, 8]

           取下标2和4的数，分别是5和4，5比4大，进行互换，[1, 2, 4, 3, 5, 7, 6, 8]

           取下标3和5的数，分别是3和7，顺序正确，则不互换，[1, 2, 4, 3, 5, 7, 6, 8]

           取下标4和6的数，分别是5和6，顺序正确，则不互换，[1, 2, 4, 3, 5, 7, 6, 8]

           取下标5和7的数，分别是7和8，顺序正确，则不互换，[1, 2, 4, 3, 5, 7, 6, 8]

第五遍，取间歇2/1.3=1

           取下标0和1的数，分别是1和3，顺序正确，则不互换，[1, 3, 5, 2, 4, 7, 6, 8]

           取下标1和2的数，分别是2和5，顺序正确，则不互换，[1, 2, 5, 3, 4, 7, 6, 8]

           取下标2和3的数，分别是4和3，4比3大，进行互换，[1, 2, 3, 4, 5, 7, 6, 8]

           取下标3和4的数，分别是3和7，顺序正确，则不互换，[1, 2, 3, 4, 5, 7, 6, 8]

           取下标4和5的数，分别是5和6，顺序正确，则不互换，[1, 2, 3, 4, 5, 7, 6, 8]

           取下标5和6的数，分别是7和8，顺序正确，则不互换，[1, 2, 3, 4, 5, 7, 6, 8]

           取下标6和7的数，分别是7和8，顺序正确，则不互换，[1, 2, 3, 4, 5, 7, 6, 8]

因为1/1.3<1了，所以没法弄了，排序结束，这个时候我们看到，数组其实已经排好了，太神奇了哈
 */
func combSort(theArray []int) []int {
	theLen := len(theArray)
	i := int(float32(theLen) / 1.3)
	for i >= 1 {
		for j := 0; j < theLen; j++ {
			if i+j >= theLen {
				i = (int)(float32(i) / 1.3)
				break
			} else {
				if theArray[j] > theArray[i+j] {
					theArray[j], theArray[i+j] = theArray[i+j], theArray[j]
				}
			}
		}
	}
	return theArray
}
func main() {
	var theArray = []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 17,255}
	fmt.Print("排序前")
	fmt.Println(theArray)
	fmt.Print("排序后")
	arrayResult := combSort(theArray)
	fmt.Println(arrayResult)
}
