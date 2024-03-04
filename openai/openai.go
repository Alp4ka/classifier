package openai

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/Alp4ka/classifier"
	"io"
	"net/http"
	"text/template"
)

type Classifier struct {
	client *http.Client
	tmpl   *template.Template

	cfg Config
}

//go:embed data/instructions.txt
var _templateText string

func NewClassifier(cfg Config) (*Classifier, error) {
	tmpl, err := template.New("").Parse(_templateText)
	if err != nil {
		return nil, err
	}

	cfg, err = prepareConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &Classifier{
			client: &http.Client{
				Timeout: cfg.Timeout,
			},
			tmpl: tmpl,
			cfg:  cfg,
		},
		nil
}

func (c *Classifier) Classify(ctx context.Context, params classifier.Params) (classifier.Result, error) {
	var (
		reqPath = c.cfg.BaseURL + "/chat/completions"
		err     error
	)

	params, err = validateParams(c.cfg.ValidateConfig, params)
	if err != nil {
		return nil, fmt.Errorf("input validation failed: %w", err)
	}

	reqBodyStruct, err := buildRequestBody(c.cfg.Temperature, c.cfg.Model, c.tmpl, params)
	if err != nil {
		return nil, err
	}
	reqBodyRaw, err := json.Marshal(reqBodyStruct)
	if err != nil {
		return nil, err
	}

	// >> HTTP.
	req, err := http.NewRequestWithContext(ctx, "POST", reqPath, bytes.NewBuffer(reqBodyRaw))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.cfg.APIKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w, status=%d", ErrRequestFailed, resp.StatusCode)
	}
	// << HTTP.

	defer func() { _ = resp.Body.Close() }()
	respBodyRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respBodyStruct respBody
	err = json.Unmarshal(respBodyRaw, &respBodyStruct)
	if err != nil {
		return nil, err
	}

	if len(respBodyStruct.Choices) == 0 {
		return nil, fmt.Errorf("choices out of range")
	}

	var resClass classifier.ClassStruct
	err = json.Unmarshal([]byte(respBodyStruct.Choices[0].Message.Content), &resClass)
	if err != nil {
		return nil, err
	}

	for i := range params.Classes {
		if params.Classes[i].Class().Eq(resClass) {
			return NewResult(params.Classes[i]), nil
		}
	}

	return nil, ErrUnknownResult
}

func buildRequestBody(temperature float64, model GPTModel, tmpl *template.Template, params classifier.Params) (*reqBody, error) {
	const messageRole = "user"

	templateData, err := NewTemplateData(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create template data: %w", err)
	}

	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, templateData)

	if err != nil {
		return nil, err
	}

	return &reqBody{
		Model: model,
		Messages: []reqBodyMsg{
			{
				Role:    messageRole,
				Content: buf.String(),
			},
		},
		Temperature: temperature,
	}, nil
}

var _ classifier.Classifier = (*Classifier)(nil)
