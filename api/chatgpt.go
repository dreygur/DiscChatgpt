package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type GPTRequest struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	Top              int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}

type GPTResponse struct {
	Id      string `json:"id,omitempty"`
	Object  string `json:"object,omitempty"`
	Created int    `json:"created,omitempty"`
	Model   string `json:"model,omitempty"`
	Choices []struct {
		Text         string `json:"text,omitempty"`
		Index        int    `json:"index,omitempty"`
		FinishReason string `json:"finish_reason,omitempty"`
	} `json:"choices,omitempty"`
	Usages struct {
		PromptTokens     int `json:"prompt_token,omitempty"`
		CompletionTokens int `json:"completion_tokens,omitempty"`
		TotalTokens      int `json:"total_tokens,omitempty"`
	} `json:"usages,omitempty"`
	Error struct {
		Message string `json:"message,omitempty"`
		Type    string `type:"message,omitempty"`
	} `json:"error,omitempty"`
}

// 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
func RequestToChatGPT(query string) (*GPTResponse, error) {
	url := "https://api.openai.com/v1/completions"

	gptReqObj := &GPTRequest{
		Model:            "text-davinci-003",
		Prompt:           query,
		Temperature:      0.5,
		MaxTokens:        256,
		Top:              1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	jsonObj, err := json.Marshal(gptReqObj)
	if err != nil {
		return nil, err
	}

	reqObj := strings.NewReader(string(jsonObj))

	client := &http.Client{Timeout: time.Minute * 20}
	req, err := http.NewRequest("POST", url, reqObj)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CHATGPT_TOKEN")))
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response GPTResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
