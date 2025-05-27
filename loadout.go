// package main implements a Metal Gear Solid V-inspired loadout builder.
package main

// Loadout represents a player's customizable loadout for a mission.
// It contains all the equipment and options a player can select.
type Loadout struct {
	PrimaryWeapon   string   // PrimaryWeapon is the main firearm.
	SecondaryWeapon string   // SecondaryWeapon is the sidearm.
	SupportWeapon   string   // SupportWeapon is a special weapon like a rocket launcher or sniper rifle.
	Items           []string // Items is a list of consumable items like grenades, C4, etc.
	Buddy           string   // Buddy is the companion accompanying the player (e.g., D-Dog, Quiet).
	Character       string   // Character is the player character (e.g., Venom Snake).
	Vehicle         string   // Vehicle is the chosen transport for deployment or support.
	MissionTime     string   // MissionTime is the selected time of day for the mission (e.g., "06:00", "18:00").
}

// LoadoutBuilder defines the interface for the builder pattern.
// It provides methods to construct a Loadout object step by step.
type LoadoutBuilder interface {
	// SetPrimaryWeapon sets the primary weapon for the loadout.
	SetPrimaryWeapon(weapon string) LoadoutBuilder
	// SetSecondaryWeapon sets the secondary weapon for the loadout.
	SetSecondaryWeapon(weapon string) LoadoutBuilder
	// SetSupportWeapon sets the support weapon for the loadout.
	SetSupportWeapon(weapon string) LoadoutBuilder
	// AddItem adds a consumable item to the loadout. It can be called multiple times.
	AddItem(item string) LoadoutBuilder
	// SetBuddy sets the buddy for the loadout.
	SetBuddy(buddy string) LoadoutBuilder
	// SetCharacter sets the player character for the loadout.
	SetCharacter(character string) LoadoutBuilder
	// SetVehicle sets the vehicle for the loadout.
	SetVehicle(vehicle string) LoadoutBuilder
	// SetMissionTime sets the mission time for the loadout.
	SetMissionTime(time string) LoadoutBuilder
	// Build finalizes the construction and returns the configured Loadout object.
	Build() Loadout
}

// loadoutBuilder is the concrete implementation of the LoadoutBuilder interface.
// It holds the Loadout object being constructed.
type loadoutBuilder struct {
	loadout *Loadout // loadout is a pointer to the Loadout object being built.
}

// NewLoadoutBuilder creates and returns a new instance of LoadoutBuilder.
// It initializes an empty Loadout object to be configured.
func NewLoadoutBuilder() LoadoutBuilder {
	return &loadoutBuilder{loadout: &Loadout{}}
}

// SetPrimaryWeapon sets the primary weapon for the loadout.
// It supports fluent chaining by returning the builder.
func (b *loadoutBuilder) SetPrimaryWeapon(weapon string) LoadoutBuilder {
	b.loadout.PrimaryWeapon = weapon
	return b
}

// SetSecondaryWeapon sets the secondary weapon for the loadout.
// It supports fluent chaining by returning the builder.
func (b *loadoutBuilder) SetSecondaryWeapon(weapon string) LoadoutBuilder {
	b.loadout.SecondaryWeapon = weapon
	return b
}

// SetSupportWeapon sets the support weapon for the loadout.
// It supports fluent chaining by returning the builder.
func (b *loadoutBuilder) SetSupportWeapon(weapon string) LoadoutBuilder {
	b.loadout.SupportWeapon = weapon
	return b
}

// AddItem adds a consumable item to the loadout's item list.
// It supports fluent chaining by returning the builder.
func (b *loadoutBuilder) AddItem(item string) LoadoutBuilder {
	b.loadout.Items = append(b.loadout.Items, item)
	return b
}

// SetBuddy sets the buddy for the loadout.
// It supports fluent chaining by returning the builder.
func (b *loadoutBuilder) SetBuddy(buddy string) LoadoutBuilder {
	b.loadout.Buddy = buddy
	return b
}

// SetCharacter sets the player character for the loadout.
// It supports fluent chaining by returning the builder.
func (b *loadoutBuilder) SetCharacter(character string) LoadoutBuilder {
	b.loadout.Character = character
	return b
}

// SetVehicle sets the vehicle for the loadout.
// It supports fluent chaining by returning the builder.
func (b *loadoutBuilder) SetVehicle(vehicle string) LoadoutBuilder {
	b.loadout.Vehicle = vehicle
	return b
}

// SetMissionTime sets the mission time for the loadout.
// It supports fluent chaining by returning the builder.
func (b *loadoutBuilder) SetMissionTime(time string) LoadoutBuilder {
	b.loadout.MissionTime = time
	return b
}

// Build finalizes the construction of the Loadout object.
// It returns the fully configured Loadout.
func (b *loadoutBuilder) Build() Loadout {
	return *b.loadout
}
