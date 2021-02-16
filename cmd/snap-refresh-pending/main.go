package main

import (
	"fmt"
	"os"

	"github.com/knitzsche/test-snapdapi/snapdapi"
	"github.com/snapcore/snapd/client"
)

func main() {

        if len(os.Args) < 2 {
                fmt.Println("Error: no command arguments provided")
                return
        }

	snap := snapdapi.NewClientAdapter()
	opts := &client.FindOptions{}
	snapname := os.Args[1]
	opts.Refresh = true
	snaps, _, _:= snap.Find(opts)
	found := false
	for _, s := range snaps {
		if s.Name == snapname {
			found = true
			fmt.Println(snapname, "has a release pending")
		}
	}
	if ! found {fmt.Println(snapname, "has no release pending") }
}
