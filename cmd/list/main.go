package main

import (
	_"fmt"
	"log"


	"github.com/knitzsche/test-remodel/snapdapi"
	"github.com/snapcore/snapd/client"
)

func main() {

	snap := snapdapi.NewClientAdapter()
	names := make([]string, 50)
	listoptions := &client.ListOptions{}
	snaps, err:= snap.List(names, listoptions)
	if err != nil{
		log.Println(err)
	} else {
		for _,s := range snaps {
			log.Printf("Snap:%s, Revision: %s\n", s.Title, s.Revision)
		}
	}
}
