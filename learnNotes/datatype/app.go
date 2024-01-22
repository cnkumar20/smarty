package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var arr = make([]string, 2)
	arr[0] = "kumar"
	arr[1] = "Nandiesh"
	fmt.Println(arr)

	var arr1 [2]int
	arr1[0] = 0
	arr1[1] = 1
	fmt.Println(arr1)

	arr2 := &arr1

	arr2[1] = 3

	fmt.Println("arr2 : ", arr2)
	fmt.Println("arr1", arr1)

	var var1 = "Lovely"
	var var2 = &var1
	fmt.Println(*var2)
	*var2 = "kumar"
	fmt.Println(var1)

}
