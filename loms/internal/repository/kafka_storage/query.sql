-- name: InsertOutbox :one
insert into "kafka_outbox" ("event", entity_type, entity_id, "data") 
    values ($1, $2, $3, $4)
    returning id;

-- name: SelectOutboxMessages :many
select * 
    from "kafka_outbox"
    where status = $1
    for update skip locked;

-- name: UpdateOutboxStatus :exec
update "kafka_outbox"
    set status = $1,
        updated_at = now()
    where id = $2;





