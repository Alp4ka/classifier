package openai

import (
	"fmt"
	"github.com/Alp4ka/classifier"
)

func validateParams(cfg ValidateConfig, params classifier.Params) (classifier.Params, error) {
	if len(params.Classes) == 0 {
		return params, fmt.Errorf("no classes provided, length was 0")
	}

	for i := range params.Input {
		if _, ok := _allowedSymbolsSet[params.Input[i]]; !ok {
			return params, fmt.Errorf("%w in input, '%s'", ErrBullshitSymbol, string(params.Input[i]))
		}
	}

	for i := range params.AdditionalContext {
		if _, ok := _allowedSymbolsSet[params.AdditionalContext[i]]; !ok {
			return params, fmt.Errorf("%w in additional context, '%s'", ErrBullshitSymbol, string(params.AdditionalContext[i]))
		}
	}

	classNames := make(map[string]struct{})
	for i := range params.Classes {
		curCls := params.Classes[i].Class()
		if _, ok := classNames[curCls.Name]; ok {
			return params, fmt.Errorf("duplicated classes, %+v", params.Classes[i].Class())
		}
		classNames[curCls.Name] = struct{}{}

		for j := range curCls.Name {
			if _, ok := _allowedSymbolsSet[curCls.Name[j]]; !ok {
				return params, fmt.Errorf("%w in class name, '%s', name: %s", ErrBullshitSymbol, string(curCls.Name[j]), curCls.Name)
			}
		}

		for j := range curCls.Description {
			if _, ok := _allowedSymbolsSet[curCls.Description[j]]; !ok {
				return params, fmt.Errorf("%w in class description, '%s', name: %s", ErrBullshitSymbol, string(curCls.Description[j]), curCls.Name)
			}
		}
	}

	if len(params.Input) > cfg.MaxInputLength {
		return params, fmt.Errorf("input was too long, %d>%d", len(params.Input), cfg.MaxInputLength)
	}

	if len(params.AdditionalContext) > cfg.MaxAdditionalContextLength {
		return params, fmt.Errorf("additional context was too long, %d>%d", len(params.AdditionalContext), cfg.MaxAdditionalContextLength)
	}

	for i := range params.Classes {
		if len(params.Classes[i].Class().Name) > cfg.MaxNameLength {
			return params, fmt.Errorf("class name was too long; %s, %d>%d", params.Classes[i].Class().Name, len(params.Classes[i].Class().Name), cfg.MaxNameLength)
		}
		if len(params.Classes[i].Class().Description) > cfg.MaxDescriptionLength {
			return params, fmt.Errorf("class description was too long; %s, %d>%d", params.Classes[i].Class().Description, len(params.Classes[i].Class().Description), cfg.MaxDescriptionLength)
		}
	}

	return params, nil
}
