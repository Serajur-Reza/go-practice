package main

import "fmt"

// execution serials -->
// 1. global
// 2. init func
// 3. main func

var a = 10

//init function
func init(){
	fmt.Println("I am the init function", a)
	a = 20
}


//standard function
// func add(a int , b int){
// 	fmt.Println(a+b)
	
// }

// global function expression only syntax
	var add= 	func (a int , b int){
		c:= a+b
		fmt.Println(c)
	}


	//higher order function
	// func process(op func(x int, y int)){
	// 	op(10,25)
	// }

	// returning a function
		func process(op func(x int, y int)) func(x int, y int) {
		return op
	}

func main(){
	// add(5,7)
	// fmt.Println("I am the main function", a)

	// IIFE
	// func (a int , b int){
	// 	c:= a+b
	// 	fmt.Println(c)
	// }(5,7)

	// process(add)

	f := process(add)
	f(23,56)

	
	add := 	func (a int , b int){
		c:= a+b
		fmt.Println(c)
	}
	add(5,8)


}




