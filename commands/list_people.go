package command

import (
	"fmt"

	"github.com/ankush953/proto-tutorial.git/utils"
)

func ListPeople() {
	fmt.Printf("\n=========LIST PEOPLE============\n")
	addressBook := utils.GetAll(utils.DATABASE)
	for _, person := range addressBook.People {
		fmt.Printf("Name: %v, email: %v\n", person.Name, person.Email)
	}
}