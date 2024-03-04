package openai

import "fmt"

var (
	ErrRequestFailed  = fmt.Errorf("request failed")
	ErrWrongConfig    = fmt.Errorf("wrong config")
	ErrUnknownResult  = fmt.Errorf("unknown result")
	ErrBullshitSymbol = fmt.Errorf("bullshit symbol")
)
