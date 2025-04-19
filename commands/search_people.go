package command

import (
	"fmt"

	"github.com/ankush953/proto-tutorial.git/proto/generated_code"
)

func Search() {
	fmt.Println("==========Welcome to service demo===========")
	query := "ankush"
	request := &generated_code.SearchRequest{
		Q: &query,
	}
	fmt.Print(request)
}