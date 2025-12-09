package main

import "fmt"

func main() {
	var a int8 = 5
	var b int8 = 6

	var x uint8 = 255 // unsigned (0 and only positive numbers)

	var j float32 = 10.22746
	var k float64 = 16.22346

	var flag bool = true

	r := '$'

	var s string = "I am Radioactive"

	fmt.Println(r)

	fmt.Printf("%c", r) // character -> c

	fmt.Println()

	fmt.Printf("%d", a) // decimal -> d

	fmt.Println()

	fmt.Printf("%.2f\n", j) // floating -> f -> Printf("%.2f", j)

	fmt.Printf("%v\n", flag) // floating -> f -> Printf("%.2f", j)

	fmt.Printf("%s\n", s)

	fmt.Printf("%T\n", s)

	fmt.Printf("**Type of variable s = %T\n", s)

	fmt.Printf("**Type of variable s = %T\n", flag)

	fmt.Println(a, b, x, j, k, flag)
}
