package command

import (
	"fmt"
	"os"

	"github.com/ankush953/proto-tutorial.git/proto/generated_code"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func AnyDemo() {
	var name, path string
	var details anypb.Any
	fmt.Printf("enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter file path: ")
	fmt.Scan(&path)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("error while reading file %v %v", path, err))
	}
	proto.Unmarshal(data, &details)
	demo := generated_code.AnyDemo{
		Name:    &name,
		Details: &details,
	}
	encoded, err := proto.Marshal(&demo)
	if err != nil {
		panic(fmt.Errorf("error while marshaling demo %v %v", demo.Details.Value, err))
	}
	parsedDemo := &generated_code.AnyDemo{}
	proto.Unmarshal(encoded, parsedDemo)
	fmt.Printf("Name: %v, Details: %v\n", *(parsedDemo.Name),
		*(&parsedDemo.Details.Value))
}
