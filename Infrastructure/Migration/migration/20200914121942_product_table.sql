-- +goose Up
-- +goose StatementBegin
CREATE TABLE Product
(
	Id INT auto_increment,
	Name VARCHAR(30) NULL,
	constraint Product_pk
		PRIMARY KEY (Id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Product cascade;
-- +goose StatementEnd
