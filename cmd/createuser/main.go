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
	opts := &client.CreateUserOptions{}
	opts.Email = os.Args[1]
	opts.ForceManaged = true
	result := &client.CreateUserResult{}
	result, err := snapd.CreateUser(opts)
	if err != nil{
		fmt.Println(err)
	} else {
		_, e := json.MarshalIndent(result, "", "    ")
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Printf("Username: %#v\n", result.Username)
			fmt.Printf("Pub ssh: %#v\n", result.SSHKeys[0])
		}
	}
}
