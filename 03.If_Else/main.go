package main

import "fmt"

func main() {
	// a := 10

	// var x int = 20

	// x = 50

	const temp = "Yellow"

	// if temp == "Radioactive" {
	// 	fmt.Println("You are a legend")
	// }else if temp == "Yellow"{
	// 	fmt.Println("Good choice")
	// }else{
	// 	fmt.Println("You are trash")
	// }

	switch temp {
	case "Radioactive":
			fmt.Println("You are a legend")
		case "Yellow" , "A sky of stars":
			fmt.Println("Good choice")
		default:
			fmt.Println("You are trash")
	}

	
	const marks = 50

	if marks >  50 {
		fmt.Println("Good")
	}else if marks ==  50{
		fmt.Println("Ijjot bachaisos")
	}else {
		fmt.Println("Fail korsos ja bhag")
	}

	// switch temp {
	// case "Radioactive":
	// 		fmt.Println("You are a legend")
	// 	case "Yellow":
	// 		fmt.Println("Good choice")
	// 	default:
	// 		fmt.Println("You are trash")
	// }
}