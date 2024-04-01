-- name: CreateOrder :one
insert into "order" ("user", status, created_at, updated_at) 
    values ($1, $2, now(), now())
    returning id;

-- name: AddOrderItem :exec
insert into order_item (order_id, sku, quantity) 
    values ($1, $2, $3);

-- name: UpdateStatusByOrderID :exec
update "order" 
    set status = $2, updated_at = now()
    where id = $1;

-- name: GetOrderByID :one
select id, "user" as user, status, created_at, updated_at 
    from "order" 
    where id = $1 limit 1;

-- name: GetOrderItemsByOrderID :many
select id, sku, quantity
    from order_item
    where order_id = $1;




