// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transaction.sql

package db

import (
	"context"
	"encoding/json"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (username, context, payload)
VALUES ($1, $2, $3)
RETURNING id, username, context, payload, is_confirmed, created_at
`

type CreateTransactionParams struct {
	Username string          `json:"username"`
	Context  string          `json:"context"`
	Payload  json.RawMessage `json:"payload"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction, arg.Username, arg.Context, arg.Payload)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Context,
		&i.Payload,
		&i.IsConfirmed,
		&i.CreatedAt,
	)
	return i, err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, username, context, payload, is_confirmed, created_at
FROM transactions
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, id int64) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Context,
		&i.Payload,
		&i.IsConfirmed,
		&i.CreatedAt,
	)
	return i, err
}

const listTransactions = `-- name: ListTransactions :many
SELECT id, username, context, payload, is_confirmed, created_at
FROM transactions
WHERE username = $1
ORDER BY id
LIMIT $2 OFFSET $3
`

type ListTransactionsParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) ListTransactions(ctx context.Context, arg ListTransactionsParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransactions, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Context,
			&i.Payload,
			&i.IsConfirmed,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
