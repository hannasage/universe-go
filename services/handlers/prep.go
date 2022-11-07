package handlers

import (
	ingredients "universe-go/models"
)

func sliceOrDice(sliceOutcome ingredients.PreparedIngredient, diceOutcome ingredients.PreparedIngredient, instruction ingredients.ChopInstruction) ingredients.PreparedIngredient {
	if instruction == ingredients.Dice {
		return diceOutcome
	} else {
		return sliceOutcome
	}
}

func prep(item ingredients.ChopItem) ingredients.PreparedIngredient {
	var prepared ingredients.PreparedIngredient
	switch item.Raw {
	case ingredients.Avocado:
		prepared = sliceOrDice(ingredients.SlicedAvocado, ingredients.DicedAvocado, item.Instruction)
	case ingredients.BellPepper:
		prepared = ingredients.SlicedBellPepper
	case ingredients.Cilantro:
		prepared = ingredients.ChoppedCilantro
	case ingredients.RedOnion:
		prepared = sliceOrDice(ingredients.SlicedRedOnion, ingredients.DicedRedOnion, item.Instruction)
	case ingredients.RomaineHead:
		prepared = ingredients.ChoppedRomaine
	case ingredients.Tomato:
		prepared = ingredients.DicedTomato
	}
	return prepared
}

func Chop(chopOrder []ingredients.ChopItem) {
	for _, item := range chopOrder {
		prepared := prep(item)
	}
}
