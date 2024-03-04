package openai

import (
	"encoding/json"
	"github.com/Alp4ka/classifier"
)

type TemplateData struct {
	StructureJSON     string
	ClassesJSON       string
	Input             string
	AdditionalContext string
}

func NewTemplateData(params classifier.Params) (*TemplateData, error) {
	classStructure := make([]byte, len(_classStructure))
	copy(classStructure, _classStructure)

	classesMarshalled, err := json.Marshal(classes(params.Classes).Classes())
	if err != nil {
		return nil, err
	}

	return &TemplateData{
		StructureJSON:     string(classStructure),
		ClassesJSON:       string(classesMarshalled),
		Input:             params.Input,
		AdditionalContext: params.AdditionalContext,
	}, nil
}

var _classStructure []byte

func init() {
	_classStructure, _ = json.Marshal(classifier.ClassStruct{})
}
