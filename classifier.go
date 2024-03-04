package classifier

import (
	"context"
)

type Class interface {
	Class() ClassStruct
}

type ClassStruct struct {
	_           [0]func()
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c ClassStruct) Eq(b ClassStruct) bool {
	return c.Name == b.Name
}

func (c ClassStruct) Class() ClassStruct {
	return c
}

var _ Class = (*ClassStruct)(nil)

type Result interface {
	Best() (Class, error)
	Worst() (Class, error)
	All() ([]Class, error)
}

type Params struct {
	Classes           []Class
	Input             string
	AdditionalContext string
}

type Classifier interface {
	Classify(ctx context.Context, params Params) (Result, error)
}
