-- +goose Up
-- +goose StatementBegin
insert into stock (sku, total_count, reserved)
values 
    (773297411, 150, 10),
    (1148162, 100, 10),
    (1625903, 100, 10),
    (2618151, 100, 10),
    (2956315, 100, 10),
    (2958025, 100, 10),
    (3596599, 100, 10),
    (3618852, 100, 10),
    (4288068, 100, 10),
    (4465995, 100, 10),
    (4487693, 100, 10),
    (4669069, 100, 10),
    (4678287, 100, 10),
    (4678816, 100, 10);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table stock;
-- +goose StatementEnd