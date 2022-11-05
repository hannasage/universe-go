package customer

import (
	"math/rand"
	"time"
	"universe-go/models/order"
	"universe-go/services/ordering"
)

// Randomly select an ingredient from an array of choices
// mimicking the behavior of a customer choosing between white or brown rice
// for example.
func chooseIngredient(choices []order.IngredientOption) order.IngredientOption {
	index := rand.Intn(len(choices) - 1)
	return choices[index]
}

// Randomly decide (with bias to false) whether item is on the side or not
// mimicking a baseless assumption that there's a 20% chance of of somebody
// asking for an ingredient on the side.
func chooseOnTheSide() bool {
	return rand.Intn(10) < 2
}

// Compiles the Option structure, like a customer choosing ingredients for
// their meal.
func MakeOption(choices []order.IngredientOption) order.CustomerChoice {
	ingredient := chooseIngredient(choices)
	multiplier := rand.Intn(2) + 1
	side := chooseOnTheSide()

	return order.CustomerChoice{
		Ingredient: ingredient,
		Multiplier: multiplier,
		Side:       side,
	}
}

// Creates a channel and starts a Goroutine that adds a new customer to the
// queue every random-n seconds. The queue uses a shared channel to coordinate
// with the OrderTaker.
func Queue() {
	customerQueue := make(chan int)
	go ordering.HandleCustomerQueue(customerQueue)
	nextCustomerInQueue := 0
	for {
		nextCustomerInQueue += 1
		customerQueue <- nextCustomerInQueue
		nextCustomerTimeout := rand.Intn(10) // Simulate randomly timed queue growth
		time.Sleep(time.Second * time.Duration(nextCustomerTimeout))
	}
	// TODO: How to keep queue running while the restaurant moves through the line
}
