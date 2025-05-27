package main

import (
	"reflect"
	"testing"
)

func TestLoadoutBuilder(t *testing.T) {
	t.Run("FullLoadout", func(t *testing.T) {
		character := "Solid Snake"
		primaryWeapon := "SOCOM"
		secondaryWeapon := "Nikita"
		item1 := "Rations"
		item2 := "Cardboard Box"
		items := []string{item1, item2}
		buddy := "Otacon"
		vehicle := "None"
		missionTime := "08:00"

		builder := NewLoadoutBuilder()
		loadout := builder.
			SetCharacter(character).
			SetPrimaryWeapon(primaryWeapon).
			SetSecondaryWeapon(secondaryWeapon).
			AddItem(item1).
			AddItem(item2).
			SetBuddy(buddy).
			SetVehicle(vehicle).
			SetMissionTime(missionTime).
			Build()

		if loadout.Character != character {
			t.Errorf("Expected Character %s, got %s", character, loadout.Character)
		}
		if loadout.PrimaryWeapon != primaryWeapon {
			t.Errorf("Expected PrimaryWeapon %s, got %s", primaryWeapon, loadout.PrimaryWeapon)
		}
		if loadout.SecondaryWeapon != secondaryWeapon {
			t.Errorf("Expected SecondaryWeapon %s, got %s", secondaryWeapon, loadout.SecondaryWeapon)
		}
		if !reflect.DeepEqual(loadout.Items, items) {
			t.Errorf("Expected Items %v, got %v", items, loadout.Items)
		}
		if loadout.Buddy != buddy {
			t.Errorf("Expected Buddy %s, got %s", buddy, loadout.Buddy)
		}
		if loadout.Vehicle != vehicle {
			t.Errorf("Expected Vehicle %s, got %s", vehicle, loadout.Vehicle)
		}
		if loadout.MissionTime != missionTime {
			t.Errorf("Expected MissionTime %s, got %s", missionTime, loadout.MissionTime)
		}
		// Test SupportWeapon was not set
		if loadout.SupportWeapon != "" {
			t.Errorf("Expected empty SupportWeapon, got %v", loadout.SupportWeapon)
		}
	})

	t.Run("PartialLoadout", func(t *testing.T) {
		character := "Raiden"
		item1 := "Coolant"
		expectedItems := []string{item1}

		builder := NewLoadoutBuilder()
		loadout := builder.
			SetCharacter(character).
			AddItem(item1).
			Build()

		if loadout.Character != character {
			t.Errorf("Expected Character %s, got %s", character, loadout.Character)
		}
		if !reflect.DeepEqual(loadout.Items, expectedItems) {
			t.Errorf("Expected Items %v, got %v", expectedItems, loadout.Items)
		}

		// Check that other fields are zero-valued
		if loadout.PrimaryWeapon != "" {
			t.Errorf("Expected empty PrimaryWeapon, got %s", loadout.PrimaryWeapon)
		}
		if loadout.SecondaryWeapon != "" {
			t.Errorf("Expected empty SecondaryWeapon, got %s", loadout.SecondaryWeapon)
		}
		if loadout.SupportWeapon != "" {
			t.Errorf("Expected empty SupportWeapon, got %s", loadout.SupportWeapon)
		}
		if loadout.Buddy != "" {
			t.Errorf("Expected empty Buddy, got %s", loadout.Buddy)
		}
		if loadout.Vehicle != "" {
			t.Errorf("Expected empty Vehicle, got %s", loadout.Vehicle)
		}
		if loadout.MissionTime != "" {
			t.Errorf("Expected empty MissionTime, got %s", loadout.MissionTime)
		}
	})

	t.Run("EmptyLoadout", func(t *testing.T) {
		builder := NewLoadoutBuilder()
		loadout := builder.Build()

		if loadout.Character != "" {
			t.Errorf("Expected empty Character, got %s", loadout.Character)
		}
		if loadout.PrimaryWeapon != "" {
			t.Errorf("Expected empty PrimaryWeapon, got %s", loadout.PrimaryWeapon)
		}
		if loadout.SecondaryWeapon != "" {
			t.Errorf("Expected empty SecondaryWeapon, got %s", loadout.SecondaryWeapon)
		}
		if loadout.SupportWeapon != "" {
			t.Errorf("Expected empty SupportWeapon, got %s", loadout.SupportWeapon)
		}
		if loadout.Items != nil { // Unset slice should be nil
			t.Errorf("Expected nil Items, got %v", loadout.Items)
		}
		if loadout.Buddy != "" {
			t.Errorf("Expected empty Buddy, got %s", loadout.Buddy)
		}
		if loadout.Vehicle != "" {
			t.Errorf("Expected empty Vehicle, got %s", loadout.Vehicle)
		}
		if loadout.MissionTime != "" {
			t.Errorf("Expected empty MissionTime, got %s", loadout.MissionTime)
		}
	})
}
