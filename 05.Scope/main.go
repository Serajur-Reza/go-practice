package main

import (
	calculate_pack "custom/Math"
	"fmt"
)

var (
	a = 20
	b = 30
)


func add(x int, y int){
	z := x + y

	fmt.Println(z)
}

func main()  {
	// p := 30
	// q := 40

	// add(p,q)

	// add(a,b)

	// add(a,p)

	// add(b,z)

	// fmt.Println(c)
	// c := 10
	// fmt.Println(c)

	// x := 18

	// if(x >= 18){
	// 	p := 10
	// 	fmt.Println(p)
	// }

	// fmt.Println(p)

	res := calculate_pack.Add(10,25)
	fmt.Println(res)
}
