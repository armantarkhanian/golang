package main

import "fmt"

func main() {
	var x = 10

	switch x {
	case 0:
		fmt.Println("x = 0")
	case 1:
		fmt.Println("x = 1")
	case 2:
		fmt.Println("x = 2")
	case 3:
		fmt.Println("x = 3")
	case 4:
		fmt.Println("x = 4")
	case 5:
		fmt.Println("x = 5")
	case 6:
		fmt.Println("x = 6")
	case 7:
		fmt.Println("x = 7")
	case 8:
		fmt.Println("x = 8")
	case 9:
		fmt.Println("x = 9")
	case 10:
		fmt.Println("x = 10")
	default:
		fmt.Println("Неизвестное число")
	}
}
