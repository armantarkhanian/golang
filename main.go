package main

import "fmt"

type User struct {
	Username string
	Friend   string
}

func main() {
	arman := User{
		Username: "Арман",
		Friend:   "Паша",
	}
	pavel := User{
		Username: "Паша",
		Friend:   "Арман",
	}

	if arman.Friend == "Паша" {
		fmt.Println("Арман и Паша друзья")
	}
	if pavel.Friend == "Арман" {
		fmt.Println("Паша и Арман друзья")
	}
}
