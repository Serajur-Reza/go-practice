package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUserDetails(usr User){
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}


// receiver function
func (usr User) printDetails(){
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

//receiver function
func (usr User) call(a int){
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {
	var user1 User

	user1 = User{
		Name: "Radioactive",
		Age:  30,
	}

	printUserDetails(user1)

	user2 := User{
		Name: "Raphael",
		Age:  25,
	}

	user2.printDetails()

	user2.call(2)

	// fmt.Println(user1.Name)
	// fmt.Println(user1.Age)
}
