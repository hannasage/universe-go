package order

import (
	"universe-go/models/ingredients"
	"universe-go/models/vessel"
)

type CustomerChoice struct {
	Ingredient ingredients.MealIngredient
	Multiplier int
	Side       bool
}

type Order struct {
	Vessel  vessel.VesselOption
	Options []CustomerChoice
}
