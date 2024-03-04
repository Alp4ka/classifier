package openai

import (
	"github.com/Alp4ka/classifier"
)

type Result struct {
	result classifier.Class
}

func NewResult(result classifier.Class) *Result {
	return &Result{
		result: result,
	}
}

func (r *Result) Best() (classifier.Class, error) {
	return r.result, nil
}

func (r *Result) Worst() (classifier.Class, error) {
	panic("not implemented")
}

func (r *Result) All() ([]classifier.Class, error) {
	panic("not implemented")
}
