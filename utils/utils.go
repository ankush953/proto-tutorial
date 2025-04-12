package utils

import (
	"fmt"
	"os"

	"github.com/ankush953/proto-tutorial.git/proto/generated_code"
	"google.golang.org/protobuf/proto"
)

func GetAll(database string) *generated_code.AddressBook {
	content, err := os.ReadFile(database)
	if err != nil {
		fmt.Printf("error when reading database: %v", err)
		os.Exit(1)
	}
	// people := make([]generated_code.Person, 0)
	addressBook := &generated_code.AddressBook{}
	err = proto.Unmarshal(content, addressBook)
	if err != nil {
		fmt.Printf("error when unmarshaling database: %v", err)
		os.Exit(1)
	}
	return addressBook
}