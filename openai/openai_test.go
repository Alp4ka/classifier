package openai

import (
	"context"
	"fmt"
	"github.com/Alp4ka/classifier"
	"os"
	"testing"
)

func TestClassificator_Classify(t *testing.T) {
	apiKey := os.Getenv("CLASSIFIER_API_KEY")
	cls, err := NewClassifier(Config{APIKey: apiKey})
	if err != nil {
		panic(err)
	}

	res, err := cls.Classify(
		context.Background(),
		classifier.Params{
			Classes: []classifier.Class{
				classifier.ClassStruct{
					Name:        "Купить пёсика",
					Description: "Песики имеют повадку пакостить",
				},
				classifier.ClassStruct{
					Name:        "Купить котика",
					Description: "Котики имеют повадку тереться об ногу",
				},
				classifier.ClassStruct{
					Name:        "Помощь оператора",
					Description: "Этот вариант выбирается, если все остальные классы не подходят, либо если пользователь напрямую просит помощи оператора",
				},
			},
			Input: "Купить тепловизор",
		},
	)
	if err != nil {
		panic(err)
	}

	best, _ := res.Best()
	fmt.Printf("%+v\n", best.Class())
}
