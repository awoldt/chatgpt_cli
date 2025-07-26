package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Command struct {
	name        string
	description string
}

var allCommands = []Command{{name: "sysinstr", description: "Sets the attitude/tone of the bots responses"}, {name: "clear", description: "Clears the entire chat conversation along with the system instructions if one was set"}, {name: "save", description: "Will save your current conversation with the bot to a local json file"}}

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

func ExecuteCommand(str string, systemInstructions *SystemInstruction, chat *ConvesationState) {
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
			hasSystemInstruction := systemInstructions.Role != ""
			*systemInstructions = SystemInstruction{Role: "", Content: ""}
			*chat = ConvesationState{Model: chat.Model, Input: []Input{}}
			if hasSystemInstruction {
				fmt.Println("successfully cleared chat and system instruction")
			} else {
				fmt.Println("successfully cleared chat")
			}

			break
		}

	case "save":
		{
			jsonChat, err := json.Marshal(chat.Input)
			if err != nil {
				fmt.Println("error: could not save chat")
				return
			}
			os.WriteFile("conversation.json", jsonChat, 0666)
			fmt.Println("successfully saved chat")
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
