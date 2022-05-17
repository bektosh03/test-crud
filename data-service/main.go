package main

import (
	"fmt"

	"github.com/bektosh03/test-crud/genprotos/datapb"
)

func main() {
	request := datapb.CollectPostsRequest{}
	fmt.Println(request)
}
