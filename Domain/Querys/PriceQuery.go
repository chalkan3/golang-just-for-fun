package Querys

import "text/template"

type PriceQuery struct {
}

func (query *PriceQuery) GetAll() *template.Template {
	tpl, err := template.New("query").Parse(`
		SELECT
			*
		FROM
			Product
		INNER JOIN Cycle ON Product.Id = Cycle.Product
		INNER JOIN Plans ON Cycle.Id = Plans.IdCycle
		ORDER BY Product.Id
	`)

	if err != nil {
		panic(err)
	}

	return tpl

}

func NewPriceQuery() *PriceQuery {
	return &PriceQuery{}
}
