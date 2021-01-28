package main

import (
	"fmt"
	"encoding/json"

	"github.com/knitzsche/test-snapdapi/snapdapi"
	"github.com/snapcore/snapd/client"
)

func main() {

	snap := snapdapi.NewClientAdapter()
	opts := &client.FindOptions{}
	opts.Refresh = true
	snaps, _, err:= snap.Find(opts)
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
