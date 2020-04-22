package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/knitzsche/test-remodel/snapdapi"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("This command requires path/to/model argument")
		return
	}

	snapC := snapdapi.NewClientAdapter()

	newModel := os.Args[1]
	modelTarget, err :=  ioutil.ReadFile(newModel)
	var chID string
	chID, err = snapC.Remodel(modelTarget)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("chId: %s\n", chID)
}
