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
	vlad := User{
		Username: "Влад",
		Friend:   "Юра",
	}

	if arman.Friend == "Паша" {
		fmt.Println("Арман дружит с Пашей")
	}
	if arman.Friend == "Влад" {
		fmt.Println("Арман дружит с Владом")
	}
	if arman.Friend == "Юра" {
		fmt.Println("Арман дружит с Юрой")
	}

	if vlad.Friend == "Арман" {
		fmt.Println("Влад дружит с Арманом")
	}
	if vlad.Friend == "Паша" {
		fmt.Println("Влад дружит с Пашей")
	}
	if vlad.Friend == "Юра" {
		fmt.Println("Влад дружит с Юрой")
	}

	if pavel.Friend == "Арман" {
		fmt.Println("Паша дружит с Арманом")
	}
	if pavel.Friend == "Влад" {
		fmt.Println("Паша дружит с Владом")
	}
	if pavel.Friend == "Юра" {
		fmt.Println("Паша дружит с Юрой")
	}
}
