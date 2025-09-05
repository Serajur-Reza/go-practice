package main

import "fmt"

func add(num1 int, num2 int) int {

	sum := num1 + num2
	// fmt.Println(sum)

	return sum
}

func getNumbers(num1 int , num2 int) (int, int){
	sum := num1 + num2

	mul := num1 * num2
	

	return sum , mul
}

func main() {
	a := 10
	b := 20

	res1 := add(a, b)
	res2 := add(5, 7)

	fmt.Println(res1, res2)

	p,q := getNumbers(10,20)
	fmt.Println(p,q)
}
