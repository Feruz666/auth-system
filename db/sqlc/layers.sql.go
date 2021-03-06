// Code generated by sqlc. DO NOT EDIT.
// source: layers.sql

package db

import (
	"context"
)

const createLayer = `-- name: CreateLayer :one
INSERT INTO layers (
  username,
  date,
  layer
) VALUES (
  $1, $2, $3
) RETURNING id, username, date, layer
`

type CreateLayerParams struct {
	Username string `json:"username"`
	Date     string `json:"date"`
	Layer    string `json:"layer"`
}

func (q *Queries) CreateLayer(ctx context.Context, arg CreateLayerParams) (Layer, error) {
	row := q.db.QueryRowContext(ctx, createLayer, arg.Username, arg.Date, arg.Layer)
	var i Layer
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Date,
		&i.Layer,
	)
	return i, err
}

const listLayer = `-- name: ListLayer :many
SELECT id, username, date, layer FROM layers
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListLayerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListLayer(ctx context.Context, arg ListLayerParams) ([]Layer, error) {
	rows, err := q.db.QueryContext(ctx, listLayer, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Layer{}
	for rows.Next() {
		var i Layer
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Date,
			&i.Layer,
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
