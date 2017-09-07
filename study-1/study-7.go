package main

import "fmt"

func main() {
	//studyArr()
	//rangeArr()
	//pointArr()
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(&arr[2])
	//modifyArr(arr)
	modifyPointArr(&arr)
}

func studyArr() {
	var arr [5]int
	arr = [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{2, 3, 4, 5, 6}
	arr2 := [4]string{1: "hello", 3: "king"}
	arr3 := [4]int{2: 3, 1: 4}
	fmt.Printf("%d\n", arr)
	fmt.Printf("%d\n", arr1)
	fmt.Printf("%d\n", arr2)
	fmt.Printf("%d\n", arr3)
}

func rangeArr() {
	arr1 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	for i, k := range arr1 {
		fmt.Println(i, k)
	}
}

func pointArr() {
	arr := [4]*int{1: new(int), 3: new(int)}
	arr[0] = new(int)
	*arr[0] = 1
	*arr[1] = 3
	fmt.Println(arr[0])
	fmt.Println(*arr[1])
}

func modifyArr(arr [5]int) {
	arr[0] = 100
	fmt.Println(&arr[0])
}

func modifyPointArr(arr *[5]int) {
	arr[2] = 333
	fmt.Println(arr[2])
	fmt.Println(&arr[2])
}
