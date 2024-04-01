// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqlc

import (
	"context"
)

const cancelReserve = `-- name: CancelReserve :exec
update stock 
    set reserved = reserved - $2  
    where sku = $1
`

type CancelReserveParams struct {
	Sku      int64
	Reserved int32
}

func (q *Queries) CancelReserve(ctx context.Context, arg CancelReserveParams) error {
	_, err := q.db.Exec(ctx, cancelReserve, arg.Sku, arg.Reserved)
	return err
}

const getBySKU = `-- name: GetBySKU :one
select sku, total_count, reserved 
    from stock 
    where sku = $1 limit 1
`

func (q *Queries) GetBySKU(ctx context.Context, sku int64) (Stock, error) {
	row := q.db.QueryRow(ctx, getBySKU, sku)
	var i Stock
	err := row.Scan(&i.Sku, &i.TotalCount, &i.Reserved)
	return i, err
}

const removeReserve = `-- name: RemoveReserve :exec
update stock 
    set total_count = total_count - $2,
        reserved = reserved - $2  
    where sku = $1
`

type RemoveReserveParams struct {
	Sku        int64
	TotalCount int32
}

func (q *Queries) RemoveReserve(ctx context.Context, arg RemoveReserveParams) error {
	_, err := q.db.Exec(ctx, removeReserve, arg.Sku, arg.TotalCount)
	return err
}

const reserve = `-- name: Reserve :exec
update stock 
    set reserved = reserved + $2  
    where sku = $1
`

type ReserveParams struct {
	Sku      int64
	Reserved int32
}

func (q *Queries) Reserve(ctx context.Context, arg ReserveParams) error {
	_, err := q.db.Exec(ctx, reserve, arg.Sku, arg.Reserved)
	return err
}
