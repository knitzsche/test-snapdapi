package main

import (
	"fmt"

	"github.com/knitzsche/test-remodel/snapdapi"
)

func main() {
	snapC := snapdapi.NewClientAdapter()
	model, err := snapC.CurrentModelAssertion()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("Model (name): %s\n", model.Model())
	fmt.Printf("Brand ID: %s\n", model.BrandID())
	fmt.Printf("Architecture: %s\n", model.Architecture())
	fmt.Printf("Classic: %v\n", model.Classic())
	fmt.Printf("Display name: %s\n", model.DisplayName())
	fmt.Printf("Series: %s\n", model.Series())
	fmt.Printf("Grade: %s\n", model.Grade())
	fmt.Printf("GadgetSnap: %v\n", model.GadgetSnap())
	fmt.Printf("GadgetTrack: %s\n", model.GadgetTrack())
	fmt.Printf("KernelSnap: %v\n", model.KernelSnap())
	fmt.Printf("Kernel: %s\n", model.Kernel())
	fmt.Printf("KernelTrack: %s\n", model.KernelTrack())
	fmt.Printf("Base: %v\n", model.BaseSnap())
	fmt.Printf("Store: %s\n", model.Store())
	fmt.Printf("Required No Essential snaps: %s\n", model.RequiredNoEssentialSnaps())
	fmt.Printf("Required With Essential Snaps: %s\n", model.RequiredWithEssentialSnaps())
	fmt.Printf("Serial Authority: %s\n", model.SerialAuthority())
	fmt.Printf("System User Authority: %s\n", model.SystemUserAuthority())
	fmt.Printf("Timestamp: %s\n", model.Timestamp())
}
