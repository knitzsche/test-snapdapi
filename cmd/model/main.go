package main

import (
	"fmt"
	

	"github.com/knitzsche/test-remodel/snapdapi"
)

func main() {
	snapC := snapdapi.NewClientAdapter()
	//model := &asserts.Model{}
	model, err := snapC.CurrentModelAssertion()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("Name: %s\n", model.Model())
	fmt.Printf("Brand ID: %s\n", model.BrandID())
	fmt.Printf("Architecture: %s\n", model.Architecture())
}
