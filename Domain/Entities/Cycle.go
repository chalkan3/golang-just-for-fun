package Entities

type Cycle struct {
	Id    int
	Plans []*Plans
}

func NewCycleEmpty() *Cycle {
	return &Cycle{}
}

func NewCycle(id int, plans []*Plans) *Cycle {
	return &Cycle{
		Id:    id,
		Plans: plans,
	}
}
