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

	snapd := snapdapi.NewClientAdapter()
	opts := &client.RemoveUserOptions{}
	opts.Username = os.Args[1]
	result, err := snapd.RemoveUser(opts)
	if err != nil{
		fmt.Println(err)
	} else {
		_, e := json.MarshalIndent(result, "", "    ")
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Printf("Username removed: %#v\n", result[0].Username)
		}
	}
}
