package command

import (
	"fmt"

	"github.com/ankush953/proto-tutorial.git/utils"
)

func ListPeople(database string) {
	fmt.Println("You have selected operation: list people")
	addressBook := utils.GetAll(database)
	for _, person := range addressBook.People {
		fmt.Printf("Name: %v, email: %v\n", person.Name, person.Email)
	}
}