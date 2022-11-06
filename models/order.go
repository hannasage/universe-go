package models

type CustomerChoice struct {
	Ingredient MealIngredient
	Multiplier int
	Side       bool
}

type Order struct {
	Vessel  VesselOption
	Options []CustomerChoice
}
