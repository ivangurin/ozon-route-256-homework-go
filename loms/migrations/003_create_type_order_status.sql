-- +goose Up
-- +goose StatementBegin
create type order_status_type as enum (
    'new',
    'awaiting_payment',
    'payed',
    'cancelled',
    'failed'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists order_status_type;
-- +goose StatementEnd