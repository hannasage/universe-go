package models

type MealIngredient int
type RawIngredient int

const (
	// Grains
	WhiteRice MealIngredient = iota
	BrownRice
	// Beans
	BlackBeans
	PintoBeans
	// Protein
	Chicken
	Steak
	Barbacoa
	Carnitas
	Sofritas
	// Extras
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

const (
	BrownRiceGrain RawIngredient = iota
	WhiteRiceGrain
	Cilantro
	CitrusJuiceMix
	Salt
	Pepper
	RedOnion
	BellPepper
	Tomato
	RawCorn
	RomaineHead
	Avocado
)
