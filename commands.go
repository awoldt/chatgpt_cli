package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

const conversationFileName = "conversation.json"

type Command struct {
	name        string
	description string
}

var allCommands = []Command{{name: "sysinstr", description: "Sets the attitude/tone of the bots responses"}, {name: "clear", description: "Clears the entire chat conversation along with the system instructions if one was set"}, {name: "save", description: "Will save your current conversation with the bot to a local json file"}, {name: "model", description: "Change the chatGPT model that is used in chat. Visit https://platform.openai.com/docs/models for all available models."}}

func ListCommands(commands []Command) {
	// make sure in A-Z order
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].name < commands[j].name
	})

	fmt.Println("--- Commands --- ")
	for _, v := range commands {
		fmt.Println("\t" + v.name + " - " + v.description)
	}
}

func ExecuteCommand(str string, systemInstructions *SystemInstruction, chat *ConvesationState, config *Config) {
	if str == "\\" {
		ListCommands(allCommands)
		return
	}

	command := str[1:]

	if !validCommand(command) {
		fmt.Println("\"" + command + "\" is not a valid command")
		return
	}

	switch command {
	case "sysinstr":
		// sets the system instruction and will be set in the conversationstate before every chat (openai docs say to do this)
		{
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("enter system instruction: ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("error: could not parse system instruction")
				return
			}
			*systemInstructions = SystemInstruction{Role: "developer", Content: input}
			fmt.Println("successfully added system instruction")
			break
		}

	case "clear":
		{
			var sb strings.Builder

			// see if theres a conversations file
			var hasSavedChat bool
			_, err := os.ReadFile(conversationFileName)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					hasSavedChat = false
				}
			} else {
				hasSavedChat = true
				// clear conversations file
				err = os.Remove(conversationFileName)
				if err != nil {
					fmt.Println("error: could not remove " + conversationFileName)
				}
			}

			// see if theres saved system instructions
			hasSystemInstruction := systemInstructions.Role != ""

			// clear system instructions
			*systemInstructions = SystemInstruction{Role: "", Content: ""}
			// clear chat
			*chat = ConvesationState{Model: chat.Model, Input: []Input{}}

			fmt.Println(hasSavedChat, hasSystemInstruction)

			if !hasSavedChat && !hasSystemInstruction {
				fmt.Println("successfully cleared chat")
				return
			} else if hasSavedChat && hasSystemInstruction {
				fmt.Println("successfully cleared chat, " + conversationFileName + ", and system instruction")
				return
			}

			sb.WriteString("successfully cleared chat")
			if hasSavedChat {
				sb.WriteString(" and " + conversationFileName)
			}
			if hasSystemInstruction {
				sb.WriteString(" and system instruction")
			}

			fmt.Println(sb.String())

			break
		}

	case "save":
		{
			if len(chat.Input) == 0 {
				fmt.Println("there is no chat history to save")
				return
			}
			jsonChat, err := json.Marshal(chat.Input)
			if err != nil {
				fmt.Println("error: could not save chat")
				return
			}
			os.WriteFile(conversationFileName, jsonChat, 0666)
			fmt.Println("successfully saved chat")
			break
		}

	case "model":
		{
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("enter model: ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("error: could not parse model text")
				return
			}

			// save new config.json
			josnData, err := json.Marshal(Config{Key: config.Key, Model: strings.TrimSpace(input)})
			if err != nil {
				fmt.Println("error: could not update model")
				return
			}
			err = os.WriteFile("config.json", josnData, 0666)
			if err != nil {
				fmt.Println("error: could not save new model in config.json")
				return
			}

			// update the model currently being used by program
			*config = Config{Model: strings.TrimSpace(input), Key: config.Key}
			fmt.Println("successfully chagned model to " + input)
			break
		}
	}
}

func validCommand(command string) bool {
	for _, v := range allCommands {
		if v.name == command {
			return true
		}
	}
	return false
}
