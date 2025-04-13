package command

import (
	"fmt"
	"os"

	"github.com/ankush953/proto-tutorial.git/proto/generated_code"
	"github.com/ankush953/proto-tutorial.git/utils"
	"google.golang.org/protobuf/proto"
)

func AddPeople() {
	fmt.Printf("\n=========ADD PEOPLE============\n")
	addressBook := utils.GetAll(utils.DATABASE)
	takeInput(addressBook)
	content, err := proto.Marshal(addressBook)
	if err != nil {
		fmt.Printf("error when marshalling person: %v", err)
		os.Exit(1)
	}
	err = os.WriteFile(utils.DATABASE, content, 0666)
	if err != nil {
		fmt.Printf("error when writing address book: %v to database. err: %v", addressBook, err)
		os.Exit(1)
	}
	fmt.Printf("records written successfully\n")
}

func takeInput(ab *generated_code.AddressBook){
	person := &generated_code.Person{}
	fmt.Printf("Please enter name: ")
	fmt.Scanln(&person.Name)
	fmt.Printf("Please enter email: ")
	fmt.Scanln(&person.Email)
	ab.People = append(ab.People, person)	
}