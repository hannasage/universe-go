package models

type RawIngredient int
type ChopIngredient int
type PreparedIngredient int
type MealIngredient int
type ChopInstruction int

// Used by grill
const (
	BrownRiceGrain RawIngredient = iota
	WhiteRiceGrain
	CitrusJuiceMix
	Salt
	Pepper
	RawCorn
)

// Used by prep
const (
	Avocado ChopIngredient = iota
	BellPepper
	Cilantro
	RedOnion
	RomaineHead
	Tomato
)

// Used by kitchen to stock line (make MealIngredient)
const (
	SlicedBellPepper PreparedIngredient = iota
	ChoppedCilantro
	DicedRedOnion
	SlicedRedOnion
	ChoppedRomaine
	SlicedAvocado
	DicedAvocado
	DicedTomato
)

// Used by line to serve customers
const (
	WhiteRice MealIngredient = iota
	BrownRice
	BlackBeans
	PintoBeans
	Chicken
	Steak
	Barbacoa
	Carnitas
	Sofritas
	FejitaVeg
	Queso
	SalsaPico
	SalsaCorn
	SalsaVerde
	SalsaRoja
	SourCream
	Cheese
	Lettuce
	Guac
)

// Instruction on how to cut a ChopIngredient
const (
	Dice ChopInstruction = iota
	Slice
)

// Struct for a single request for Prep to chop
type ChopItem struct {
	Raw         ChopIngredient
	Instruction ChopInstruction
}
