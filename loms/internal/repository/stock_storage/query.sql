-- name: GetBySKU :one
select sku, total_count, reserved 
    from stock 
    where sku = $1 limit 1;

-- name: Reserve :exec
update stock 
    set reserved = reserved + $2  
    where sku = $1;

-- name: RemoveReserve :exec
update stock 
    set total_count = total_count - $2,
        reserved = reserved - $2  
    where sku = $1;

-- name: CancelReserve :exec
update stock 
    set reserved = reserved - $2  
    where sku = $1;
