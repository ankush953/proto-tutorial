package main

import (
	"fmt"
	"os"

	command "github.com/ankush953/proto-tutorial.git/commands"
)

const DATABASE = "gateways/database.txt"

func main() {
	fmt.Print("Please enter\n1. To list people\n2.To add person\nPress any other for exit ")
	var operation string
	fmt.Scanln(&operation)
	if operation == "1" {
		command.ListPeople(DATABASE)
	}else if operation == "2" {
		command.AddPeople(DATABASE)
	}else {
		fmt.Printf("Thank you for using address book services")
		os.Exit(0)
	}
}
