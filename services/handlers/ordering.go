package handlers

// Acts as the employee taking an order from a customer and building the
// meal. Uses the IngredientLine to fulfill the order.
func takeOrder(orderId int) {
	// TODO: Interaction between employee and customer
	println(orderId)
}

// Goroutine for handling the customer <-> employee interaction
func HandleCustomerQueue(c chan int) {
	for customer := range c {
		takeOrder(customer)
	}
}
