package openai

import (
	"fmt"
	"time"
)

const (
	ClassifierDefaultTimeout                    = 5 * time.Second
	ClassifierDefaultMaxInputLength             = 500
	ClassifierDefaultMaxNameLength              = 50
	ClassifierDefaultMaxDescriptionLength       = 500
	ClassifierDefaultMaxAdditionalContextLength = 300
	ClassifierDefaultBaseURL                    = "https://api.proxyapi.ru/openai/v1"
	ClassifierDefaultModel                      = GPT3_5Turbo
	ClassifierDefaultTemperature                = 0
	ClassifierDefaultAllowBullshit              = false
)

type ValidateConfig struct {
	MaxInputLength             int
	MaxNameLength              int
	MaxDescriptionLength       int
	MaxAdditionalContextLength int
	AllowBullshit              bool
}

type Config struct {
	ValidateConfig
	Model       GPTModel
	Timeout     time.Duration
	Temperature float64
	APIKey      string
	BaseURL     string
}

func prepareConfig(cfg Config) (Config, error) {
	if cfg.Timeout == 0 {
		cfg.Timeout = ClassifierDefaultTimeout
	} else if cfg.Timeout < 0 {
		return Config{}, fmt.Errorf("%w: timeout value cannot be less than 0, got: %d", ErrWrongConfig, cfg.Timeout)
	}

	if cfg.MaxInputLength == 0 {
		cfg.MaxInputLength = ClassifierDefaultMaxInputLength
	} else if cfg.MaxInputLength < 0 {
		return Config{}, fmt.Errorf("%w: max input length value cannot be less than 0, got: %d", ErrWrongConfig, cfg.MaxInputLength)
	}

	if len(cfg.Model) == 0 {
		cfg.Model = ClassifierDefaultModel
	} else if !isSupportedModel(cfg.Model) {
		return Config{}, fmt.Errorf("%w: unsupported model, got: %s", ErrWrongConfig, cfg.Model)
	}

	if !cfg.AllowBullshit {
		cfg.AllowBullshit = ClassifierDefaultAllowBullshit
	}

	if cfg.Temperature == 0 {
		cfg.Temperature = ClassifierDefaultTemperature
	} else if cfg.Temperature < 0. || cfg.Temperature > 1. {
		return Config{}, fmt.Errorf("%w: temperature value should be less than 1. and greater than 0., got: %f", ErrWrongConfig, cfg.Temperature)
	}

	if cfg.MaxNameLength == 0 {
		cfg.MaxNameLength = ClassifierDefaultMaxNameLength
	} else if cfg.MaxNameLength < 0 {
		return Config{}, fmt.Errorf("%w: max name length value cannot be less than 0, got: %d", ErrWrongConfig, cfg.MaxNameLength)
	}

	if cfg.MaxDescriptionLength == 0 {
		cfg.MaxDescriptionLength = ClassifierDefaultMaxDescriptionLength
	} else if cfg.MaxDescriptionLength < 0 {
		return Config{}, fmt.Errorf("%w: max description length value cannot be less than 0, got: %d", ErrWrongConfig, cfg.MaxDescriptionLength)
	}

	if cfg.MaxAdditionalContextLength == 0 {
		cfg.MaxAdditionalContextLength = ClassifierDefaultMaxAdditionalContextLength
	} else if cfg.MaxAdditionalContextLength < 0 {
		return Config{}, fmt.Errorf("%w: max description length value cannot be less than 0, got: %d", ErrWrongConfig, cfg.MaxAdditionalContextLength)
	}

	if len(cfg.BaseURL) == 0 {
		cfg.BaseURL = ClassifierDefaultBaseURL
	}

	return cfg, nil
}
