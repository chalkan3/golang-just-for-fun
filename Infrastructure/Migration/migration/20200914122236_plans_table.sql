-- +goose Up
-- +goose StatementBegin
CREATE TABLE Plans
(
	Id INT auto_increment,
	IdCycle INT NULL,
	Type VARCHAR(30) NULL,
	PriceRenew DOUBLE NULL,
	PriceOrder DOUBLE NULL,
	Months INT NULL,
	CONSTRAINT Plans_pk
		PRIMARY KEY (Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Plans cascade;
-- +goose StatementEnd
