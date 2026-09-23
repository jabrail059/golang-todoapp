package core_http_types

import (
	"encoding/json"

	"github.com/jabrail059/golang-todoapp/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true

	if string(b) == "null" {
		n.Value = nil

		return nil
	}

	var value T

	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.Value = &value

	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{
		Value: n.Value,
		Set:   n.Set,
	}
}

/*
JSON: {}
NULLABLE:
	- Value: *nil
	- Set: false

_________________

JSON: {
	"phone_number": "+79270050505"
}
NULLABLE:
	- Value: *"+79270050505"
	- Set: true

_________________

JSON: {
	"phone_number": null
}
NULLABLE:
	- Value: *nil
	- Set: true
*/
