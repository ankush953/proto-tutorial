package command

import (
	"fmt"
	"os"

	"github.com/ankush953/proto-tutorial.git/proto/generated_code"
	"google.golang.org/protobuf/proto"
)

func EnumDemo() {
	fmt.Print("\n================ENUM DEMO===============\n")
	indore := &generated_code.EnumContainer{
		Destination: generated_code.DESTINATION_DESTINATION_INDORE.Enum(),
	}
	fmt.Printf("Selected destination is %v\n", indore.Destination.String())
	content, err := proto.Marshal(indore); if err != nil {
		fmt.Printf("Error when marshalling enum %v", indore)
		os.Exit(1)
	}
	os.WriteFile("enum.txt", content, 0666)
	os.Stdin.Close()

	parsedEnum := &generated_code.EnumContainer{}
	content, err = os.ReadFile("enum.txt"); if err != nil {
		fmt.Printf("Error when marshalling enum %v", indore)
		os.Exit(1)
	}
	proto.Unmarshal(content, parsedEnum)
	fmt.Printf("Received destination is %v\n", parsedEnum.Destination.String())
}