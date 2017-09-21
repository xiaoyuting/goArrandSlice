package main

import "fmt"

func printSlice(x []int) {
	fmt.Println("len ==%d cap==%d slice==%v\n", len(x), cap(x), x)

}

func main() {

	var blance = [5]int{1, 2, 3, 4, 5}
	blance[4] = 50
	fmt.Println(blance)

	fmt.Println("wewqewq")
	//var numbers = make([]int, 3, 5)
	//var numbers []int
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(numbers)
	printSlice(numbers)
	if numbers == nil {
		fmt.Println("切片是空的")
	}
	fmt.Println("numbers[1:4]", numbers[1:4])
	fmt.Println("numbers[:3]", numbers[:3])
	fmt.Println("numbers[4:]", numbers[4:])
	//numbers1 := make([]int, 10, 20)

	numbers1 := numbers[:4]
	fmt.Println("arr=", numbers1)
	numbers2 := make([]int, 0, 5)
	fmt.Println("numbers2==%S", numbers2)

	var arr []int
	printSlice(arr)
	/* 允许追加空切片 */
	arr = append(arr, 0)
	printSlice(arr)
	/* 向切片添加一个元素 */
	arr = append(arr, 1)
	printSlice(arr)
	arr = append(arr, 2, 3, 4)
	printSlice(arr)

	arr1 := make([]int, len(arr), cap(arr)*2)
	printSlice(arr1)
	copy(arr1, arr)
	printSlice(arr1)
}
