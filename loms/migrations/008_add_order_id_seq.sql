-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE order_id_manual_seq INCREMENT 10 START 10;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SEQUENCE order_id_manual_seq;
-- +goose StatementEnd