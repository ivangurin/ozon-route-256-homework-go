-- +goose Up
-- +goose StatementBegin
create table order_item (
    id bigserial 
        constraint order_item_pk primary key,
    order_id bigint not null,
    sku bigint not null,
    quantity int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists order_item;
-- +goose StatementEnd