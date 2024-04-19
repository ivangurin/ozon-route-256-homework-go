-- +goose Up
-- +goose StatementBegin
create table kafka_outbox (
    id varchar(36) primary key not null default gen_random_uuid(),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    "status" text default 'new',
    error text,
    "event" text,
    entity_type text,
    entity_id text,
    "data" text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists kafka_outbox;
-- +goose StatementEnd