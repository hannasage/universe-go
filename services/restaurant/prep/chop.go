package prep

import (
	"errors"
	ingredients "universe-go/models"
)

func verifyNeedsChop(list []ingredients.RawIngredient) ([]ingredients.RawIngredient, error) {
	var toChop []ingredients.RawIngredient
	for _, v := range list {
		switch v {
		case ingredients.BellPepper:
		case ingredients.Cilantro:
		case ingredients.RedOnion:
		case ingredients.RomaineHead:
		case ingredients.Avocado:
		case ingredients.Tomato:
			toChop = append(toChop, v)
		}
	}
	if len(list) != len(toChop) {
		return toChop, errors.New("un-choppable items have been removed from the chop list")
	} else {
		return toChop, nil
	}
}

func Chop(chopOrder []ingredients.RawIngredient) {
	// Ensure we're only chopping the good stuff
	chopList, chopErrors := verifyNeedsChop(chopOrder)
	if chopErrors != nil {
		println(chopErrors)
	}
}
