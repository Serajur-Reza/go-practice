package main

import "fmt"

func main() {
	// arr := [6]string{"This", "is", "go", "interview", "questions"}

	// fmt.Println(arr)

	// s := arr[1:4]

	// fmt.Println(s)

	// s1 := s[1:2]

	// fmt.Println(s1)
	// fmt.Println(len(s1))

	// fmt.Println(cap(s1))

	// nums := []int{1,2,3} //slice literal

	// fmt.Println("slice pointer", nums, "len", len(nums), "cap", cap(nums))

	// nums1 := make([]int, 5,7)
	// nums1[4] =5

	// fmt.Println("slice pointer", nums1, "len", len(nums1), "cap", cap(nums1))

	// var nums2 []int //empty slice/ nil slice

	// nums2 = append(nums2,1, 2 , 3)

	// fmt.Println(nums2)

	// s is slice. slice maintains 3 things 1. pointer, 2. length, 3. capacity. pointer is the reference of starting index. length is end index  - start index. capacity is array length - start index

	// start index inclusive. end index exclusive

	var x []int

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x

	x = append(x, 4)

	fmt.Println(x)
	fmt.Println(y)
	y = append(y, 5)

	z := x
	z = append(z, 8)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	x[0] = 10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)


	// nums4 := []int{1, 2, 3, 4, 5}
	// nums4 = append(nums4, 6)
	// nums4 = append(nums4, 7)

	// a1 := nums4[4:]

	// y := changeSlice(a1)

	// fmt.Println(nums4)
	// fmt.Println(y)

	// fmt.Println(nums4[0:8])

	// print(5,6,8,9,4,5,3,6)
}


func changeSlice(a []int)  []int{
	a[0] = 10
	a = append(a, 11)
	return a
}

func print(numbers ...int){
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}