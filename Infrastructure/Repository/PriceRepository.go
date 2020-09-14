package Repository

import (
	"database/sql"
	"fmt"

	d "../Database"
)

type PriceRepository struct {
	dbConnection *sql.DB
}

func (priceRepository *PriceRepository) Test() string {
	var schema_name string
	err := priceRepository.dbConnection.QueryRow("SELECT schema_name FROM information_schema.schemata LIMIT 1").Scan(&schema_name)

	fmt.Println(err)
	return schema_name
}

func NewPriceRepository(database *d.Database) *PriceRepository {
	return &PriceRepository{
		dbConnection: database.GetConnection(),
	}
}
