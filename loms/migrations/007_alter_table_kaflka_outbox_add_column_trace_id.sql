-- +goose Up
-- +goose StatementBegin
alter table kafka_outbox 
    add column if not exists trace_id text,
    add column if not exists span_id text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table order_seller_posting 
    drop column if exists trace_id,
    drop column if exists span_id;
-- +goose StatementEnd