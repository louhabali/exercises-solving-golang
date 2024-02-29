package piscine

type food struct {
	burgertime  int
	chippstime  int
	nuggetstime int
}

func setFood(ptr *food) {
	ptr.burgertime = 15
	ptr.chippstime = 10
	ptr.nuggetstime = 12
}

func FoodDeliveryTime(order string) int {
	orders := food{}
	setFood(&orders)
	if order == "burger" {
		return orders.burgertime
	}
	if order == "chips" {
		return orders.chippstime
	}
	if order == "nuggets" {
		return orders.nuggetstime
	}
	return 404
}
