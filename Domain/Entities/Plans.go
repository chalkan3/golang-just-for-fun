package Entities

type Plans struct {
	Id         int
	Type       string
	PriceRenew float32
	PriceOrder float32
	Months     int
}

func NewPlansEmpty() *Plans {
	return &Plans{}
}

func NewPlans(id int,
	types string,
	priceRenew float32,
	priceOrder float32, months int) *Plans {
	return &Plans{
		Id:         id,
		Type:       types,
		PriceRenew: priceRenew,
		PriceOrder: priceOrder,
		Months:     months,
	}
}
