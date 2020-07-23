package main

import (
	"fmt"
	"encoding/json"

	"github.com/knitzsche/test-remodel/snapdapi"
	"github.com/snapcore/snapd/client"
)

func main() {

	snap := snapdapi.NewClientAdapter()
	names := make([]string, 50)
	listoptions := &client.ListOptions{}
	listoptions.All = true
	snaps, err:= snap.List(names, listoptions)
	if err != nil{
		fmt.Println(err)
	} else {
		for _,s := range snaps {
			b, e := json.MarshalIndent(s, "", "    ")
			if e != nil {
				fmt.Println(e)
			} else {
				fmt.Printf("%s\n", b)
			}
		}
	}
}
