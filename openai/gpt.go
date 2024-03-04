package openai

type GPTModel string

const (
	GPT4        GPTModel = "gpt-4"
	GPT4Turbo   GPTModel = "gpt-4-1106-preview"
	GPT3_5      GPTModel = "gpt-3.5"
	GPT3_5Turbo GPTModel = "gpt-3.5-turbo"
)

var _supportedModels = map[GPTModel]struct{}{
	GPT4:        {},
	GPT4Turbo:   {},
	GPT3_5:      {},
	GPT3_5Turbo: {},
}

func isSupportedModel(model GPTModel) bool {
	_, ok := _supportedModels[model]
	return ok
}
