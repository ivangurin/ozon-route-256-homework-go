-- +goose Up
-- +goose StatementBegin
create table "order" (
    id bigserial 
        constraint order_pk primary key,
    "user" bigint not null,
    "status" order_status_type,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now() 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists order;
-- +goose StatementEnd