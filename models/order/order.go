package order

import "universe-go/models/vessel"

type IngredientOption string

const (
	// Grains
	WhiteRice IngredientOption = "rice-white"
	BrownRice IngredientOption = "rice-brown"
	// Beans
	BlackBeans IngredientOption = "beans-black"
	PintoBeans IngredientOption = "beans-pinto"
	// Proteins
	Chicken  IngredientOption = "chicken"
	Steak    IngredientOption = "steak"
	Barbacoa IngredientOption = "barbacoa"
	Carnitas IngredientOption = "carnitas"
	Sofritas IngredientOption = "sofritas"
	// Extras
	FejitaVeg  IngredientOption = "fejita-veggies"
	Queso      IngredientOption = "queso"
	Pico       IngredientOption = "salsa-pico"
	Corn       IngredientOption = "salsa-corn"
	SalsaVerde IngredientOption = "salsa-verde-med"
	SalsaRoja  IngredientOption = "salsa-roja-hot"
	SourCream  IngredientOption = "sour-cream"
	Cheese     IngredientOption = "cheese-monterey-jack"
	Lettuce    IngredientOption = "lettuce-romaine"
	Guac       IngredientOption = "guac"
)

type CustomerChoice struct {
	Ingredient IngredientOption
	Multiplier int
	Side       bool
}

type Order struct {
	Vessel  vessel.VesselOption
	Options []CustomerChoice
}
