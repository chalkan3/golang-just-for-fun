package Entities

type Products struct {
	Id    int
	Name  string
	Cycle *Cycle
}

func NewProductsEmpty() *Products {
	return &Products{}
}

func NewProducts(id int, name string, cycle *Cycle) *Products {
	return &Products{
		Id:    id,
		Name:  name,
		Cycle: cycle,
	}
}
