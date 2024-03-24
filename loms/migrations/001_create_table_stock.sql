-- +goose Up
-- +goose StatementBegin
create table stock (
    sku bigint 
        constraint stock_pk primary key,
    total_count int not null,
    reserved int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists stock;
-- +goose StatementEnd