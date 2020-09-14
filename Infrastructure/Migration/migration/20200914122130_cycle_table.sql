-- +goose Up
-- +goose StatementBegin
CREATE TABLE Cycle
(
	Id INT auto_increment,
	Product INT NULL,
	CONSTRAINT Cycle_pk
		PRIMARY KEY (Id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Cycle cascade;
-- +goose StatementEnd
