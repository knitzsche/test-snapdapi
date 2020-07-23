package main

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/knitzsche/test-remodel/snapdapi"
	"github.com/snapcore/snapd/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no command arguments provided")
		return
	}

	snapC := snapdapi.NewClientAdapter()
	snap := &client.Snap{}
	results := &client.ResultInfo{}
	snap, results, err:= snapC.Snap(os.Args[1])
	if err != nil{
		fmt.Println(err)
	} else {
		b, _ := json.MarshalIndent(snap, "", "    ")
		fmt.Printf("Snap:\n %s\n", b)
		fmt.Printf("ResultInfo\n: %#v\n", results)
	}
}
