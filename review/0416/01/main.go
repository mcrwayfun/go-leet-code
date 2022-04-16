package main

import "fmt"

/**
数组：arr=[2, 3, 3, 3, 6, 9, 9]
维护两个指针i和j，其中i遍历数组arr，j指向有效数组的最后一个位置。
1、i=0, j=0, arr[i]=2 == arr[j]=2, 不交换
2、i=1, j=0, arr[i]=3 != arr[j]=2, j=1, arr[j]=arr[i]=3,
此时arr=arr=[2, 3, 3, 3, 6, 9, 9]
3、i=2, j=1, arr[i]=3 == arr[j]=3, 不交换
4、i=3, j=1, arr[i]=3 == arr[j]=3, 不交换
5、i=4, j=1, arr[i]=6 != arr[j]=3, j=2, arr[j]=arr[i]=6,
此时arr=arr=[2, 3, 6, 3, 6, 9, 9]
6、i=5, j=2, arr[i]=9 != arr[j]=6, j=3, arr[j]=arr[i]=9,
此时arr=arr=[2, 3, 6, 9, 6, 9, 9]
7、i=6, j=3, arr[i]=9 == arr[3]=9, 不交换
*/
func remove(arr []int) int {
	var i, j int
	for i = 0; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i]
		}
	}
	return j + 1
}

func main() {
	arr := []int{2, 3, 3, 3, 6, 9, 9}
	fmt.Println(remove(arr))
}
