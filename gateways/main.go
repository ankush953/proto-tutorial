package main

import (
	"fmt"
	"os"

	command "github.com/ankush953/proto-tutorial.git/commands"
)

const DATABASE = "gateways/database.txt"

func main() {
	fmt.Printf("Please enter\n1. list_people \n2.To add_people\n3. singularity_demo\n4. enum_demo\n5. Oneof Demo\nEnter option: ")
	operationsMap := map[string]interface{}{
		"1": command.ListPeople,
		"2": command.AddPeople,
		"3": command.SingularityDemo,
		"4": command.EnumDemo,
		"5": command.OneofDemo,
	}
	var operation string
	fmt.Scanln(&operation)
	fn, ok := operationsMap[operation]
	if ok {
		fn.(func())()
	}
	fmt.Printf("Thank you for using address book services")
	os.Exit(0)
}
