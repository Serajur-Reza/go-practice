package main

import "fmt"

// func a() {
// 	i := 0

// 	fmt.Println("first", i)

// 	defer fmt.Println("second", i) // fmt.Println("second",0) // runs just before return

// 	i = i + 1 // or i++

// 	defer fmt.Println("fourth", i)

// 	defer fmt.Println("fifth", i)

// 	fmt.Println("third", i)

// 	return
// }

func sum(a int, b int) (result int) {
	result = a + b

	// return result

	return
}

func calculate() (result int) {
	fmt.Println("first", result)

	show := func() {
		result = result + 10
		fmt.Println("Defer", result)
	}

	defer show()

	result = 5
	fmt.Println("second", result)

	return
}

func calc() int {
	result := 0
	fmt.Println("first", result)

	show := func() {
		result = result + 10
		fmt.Println("Defer", result)
	}

	defer show()

	result = 5
	fmt.Println("second", result)

	return result
}

func calc1() (result int) {
	fmt.Println("first", result)

	show := func() {
		result = result + 10
	}

	defer show()

	result = 5

	p := func(a int) {
		fmt.Println("ami", a)
	}

	defer p(result)

	defer fmt.Println("the result", result)

	fmt.Println("Second", result)

	defer fmt.Println(5)

	return
}

func calc2() int {
	result := 0
	fmt.Println("first", result)

	show := func() {
		result = result + 10
	}

	defer show()

	result = 5

	p := func(a int) {
		fmt.Println("ami", a)
	}

	defer p(result)

	defer fmt.Println("the result", result)

	fmt.Println("Second", result)

	defer fmt.Println(5)

	result = 100

	return result
}

func main() {
	// a()

	// res := sum(3, 4)

	// fmt.Println(res)

	a := calculate()
	fmt.Println("main first", a)

	b := calc()
	fmt.Println("main second", b)

	c := calc1()
	fmt.Println("final first", c)

	d := calc2()
	fmt.Println("final second", d)

}
