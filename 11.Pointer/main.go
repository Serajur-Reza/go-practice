package main

import "fmt"


type User struct {
	Name string
	Age int
	Salary float64
}

func print(numbers *[3]int){
	fmt.Println(*numbers)
}

func main()  {
	x := 20

	p := &x // amperdand & => address of

	// x = 30

	*p = 30

	// * value at address
	fmt.Println(x ,p)
	fmt.Println(x ,*p)

	arr := [3]int{1,2,3}
	print(&arr)

	radioactive := User{
		Name: "Radioactive",
		Age: 30,
		Salary: 0,
	}

	fmt.Println(radioactive)

	t := &radioactive

	fmt.Println(t.Name)
}
