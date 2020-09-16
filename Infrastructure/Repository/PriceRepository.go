package Repository

import (
	"database/sql"
	"fmt"

	Entity "../../Domain/Entities"
	Querys "../../Domain/Querys"
	d "../Database"
)

type PriceRepository struct {
	dbConnection *sql.DB
	priceQuery   *Querys.PriceQuery
}

func (repository *PriceRepository) GetAll() []*Entity.Products {

	var products []*Entity.Products
	var tmplQuery = repository.priceQuery.GetAll()

	var schema_name string
	err := repository.dbConnection.QueryRow(tmplQuery.String()).Scan(&schema_name)

	fmt.Println(err)

	return products
}

func NewPriceRepository(database *d.Database, priceQuery *Querys.PriceQuery) *PriceRepository {
	return &PriceRepository{
		dbConnection: database.GetConnection(),
		priceQuery:   priceQuery,
	}
}
