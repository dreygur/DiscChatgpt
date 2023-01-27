package api

type GPTRequest struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	Top              int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresensePenalty  int     `json:"presense_penalty"`
}

type GPTResponse struct {
	Choices []string `json:"choices"`
}

// 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
func RequestToChatGPT(query, apiKey string) (string, error) {
	return "", nil
}
