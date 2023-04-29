package entity

import (
	"errors"
)

type (
	Pagination struct {
		Offset int `json:"offset" example:"0"`
		Limit  int `json:"limit" example:"1"`
		Total  int `json:"total" example:"100"`
	}

	Collection[T any] struct {
		Pagination
		Edges []T `json:"edges" example:"[{ 'userID': 1, 'name': 'Luffy' }]"`
	}

	AccountId int
)

var (
	ErrNoRows     = errors.New("error no rows")
	ErrDuplicated = errors.New("error duplicated")
)

func NewCollection[T any](edges []T, offset, limit, total int) Collection[T] {
	return Collection[T]{
		Edges: edges,
		Pagination: Pagination{
			Offset: offset,
			Limit:  limit,
			Total:  total,
		},
	}
}