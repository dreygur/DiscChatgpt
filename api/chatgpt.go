package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	PresensePenalty  int     `json:"presense_penalty"`
}

type GPTResponse struct {
	Choices []string `json:"choices"`
}

// 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
func RequestToChatGPT(query, apiKey string) (string, error) {
	url := "https://api.openai.com/v1/completions"

	gptReqObj := &GPTRequest{
		Model:            "text-davinci-003",
		Prompt:           query,
		Temperature:      0.5,
		MaxTokens:        256,
		Top:              1,
		FrequencyPenalty: 0,
		PresensePenalty:  0,
	}

	jsonObj, err := json.Marshal(gptReqObj)
	if err != nil {
		log.Println(err)
	}

	reqObj := strings.NewReader(string(jsonObj))

	client := &http.Client{Timeout: time.Millisecond * 6000}
	req, err := http.NewRequest("POST", url, reqObj)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))

	return "", nil
}
