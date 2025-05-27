// package main implements a Metal Gear Solid V-inspired loadout builder.
package main

import "fmt"

// main is the entry point of the application.
// It demonstrates the usage of the LoadoutBuilder to create and print a loadout.
func main() {
	// Create a new LoadoutBuilder instance.
	builder := NewLoadoutBuilder()

	// Use the builder to configure a Loadout object using a fluent interface.
	// Each SetX or AddX method configures a part of the loadout and returns the builder,
	// allowing for method chaining. The Build() method finalizes the object.
	myLoadout := builder.
		SetCharacter("Venom Snake").
		SetPrimaryWeapon("AM MRS-4").
		SetSecondaryWeapon("WU Silent Pistol").
		AddItem("Magazine").
		AddItem("Water Pistol").
		SetBuddy("D-Dog").
		SetVehicle("Jeep").
		SetMissionTime("18:00").
		Build() // Finalize the loadout creation.

	// Print the details of the configured MGSV loadout.
	fmt.Println("MGSV Loadout:")
	fmt.Printf("Character: %s\n", myLoadout.Character)
	fmt.Printf("Primary Weapon: %s\n", myLoadout.PrimaryWeapon)
	fmt.Printf("Secondary Weapon: %s\n", myLoadout.SecondaryWeapon)
	fmt.Printf("Items: %v\n", myLoadout.Items) // Items slice will be printed in its default format.
	fmt.Printf("Buddy: %s\n", myLoadout.Buddy)
	fmt.Printf("Vehicle: %s\n", myLoadout.Vehicle)
	fmt.Printf("Mission Time: %s\n", myLoadout.MissionTime)
}
