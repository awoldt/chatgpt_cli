package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type SystemInstruction struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

var systemInstructions SystemInstruction

type Content struct {
	Text string `json:"text"`
}

type Output struct {
	Content *[]Content `json:"content"` // possible undefined in some resposnes
}

type OpenAiResponse struct {
	Output []Output `json:"output"`
}

type Input struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ConvesationState struct {
	Model string  `json:"model"`
	Input []Input `json:"input"`
}

const apiEndpoint = "https://api.openai.com/v1/responses"

func main() {
	var chat ConvesationState

	for {
		fmt.Print("(you): ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error: could not get input")
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if []rune(input)[0] == '\\' {
			ExecuteCommand(input, &systemInstructions, &chat, &UserConfig)
		} else {
			// if theres a system instruction set, append BEFORE the new chat
			if systemInstructions.Role != "" {
				chat.Input = append(chat.Input, Input(systemInstructions))
			}

			chat.Input = append(chat.Input, Input{Role: "user", Content: input})
			newRequest(&chat)
		}
	}
}

func newRequest(chat *ConvesationState) {
	// pass up the entire conversation state (including all previous user and assistant messages)
	payload, err := json.Marshal(ConvesationState{Model: UserConfig.Model, Input: chat.Input})
	if err != nil {
		log.Fatal("error: could not parse json body")
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("error: there was an error while constructing new response to to open")
	}

	req.Header.Add("Authorization", "Bearer "+UserConfig.Key)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer res.Body.Close()
	var response OpenAiResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("error: could not parse json response")
	}

	var assistantResponse string
	for i := 0; i < len(response.Output); i++ {
		if response.Output[i].Content != nil {
			assistantResponse = (*response.Output[i].Content)[0].Text
			break
		}
	}

	fmt.Println("(bot): " + assistantResponse)

	// add the assistants response to the conversation state
	chat.Input = append(chat.Input, Input{Role: "assistant", Content: assistantResponse})
}
