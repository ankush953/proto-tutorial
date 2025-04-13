package command

import (
	"fmt"
	"os"

	"github.com/ankush953/proto-tutorial.git/proto/generated_code"
	"google.golang.org/protobuf/proto"
)

func SingularityDemo() {
	fmt.Printf("\n=========SINGULARITY DEMO============\n")
	const SINGLUARITY_DEMO_FILE = "singularity_demo.txt"
	var zero int32 = 0
	demo := &generated_code.SingularCardinalityDemo{
		OptionalField: &zero,
		ImplicitField: 0,
	}
	content, err := proto.Marshal(demo); if err != nil {
		fmt.Printf("Error when marshalling %v", demo)
		os.Exit(1)
	}
	os.WriteFile(SINGLUARITY_DEMO_FILE, content, 0666)

	content, err = os.ReadFile(SINGLUARITY_DEMO_FILE); if err != nil {
		fmt.Printf("Error when un-marshalling %v", SINGLUARITY_DEMO_FILE)
		os.Exit(1)
	}

	parsedDemo := &generated_code.SingularCardinalityDemo{}
	proto.Unmarshal(content, parsedDemo)

	fmt.Println(parsedDemo)
}