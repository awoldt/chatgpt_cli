package main

import (
	"fmt"
	"strings"
)

type Command struct {
	name        string
	description string
}

var allCommands = []Command{{name: "sysinstr", description: "Sets the attitude/tone of the bots responses"}}

func ListCommands() {
	fmt.Println("--- Commands --- ")
	for _, v := range allCommands {
		fmt.Println("\t" + v.name + " - " + v.description)
	}
}

func ExecuteCommand(str string, systemInstructions *SystemInstruction) {
	if str == "\\" {
		ListCommands()
		return
	}

	command := strings.Split(str[1:], "=")
	var commandKey, commandValue string
	commandKey = command[0]

	if !validCommand(commandKey) {
		fmt.Println("the command \"" + commandKey + "\" is not a valid command")
		return
	}
	if len(command) == 1 {
		fmt.Println("invalid command format (missing \"=\")")
		return
	}

	commandValue = command[1]
	if commandValue == "" {
		fmt.Println("you must assign a value to the command")
		return
	}

	switch commandKey {
	case "sysinstr":
		// sets the system instruction and will be set in the conversationstate before every chat (openai docs say to do this)
		{
			*systemInstructions = SystemInstruction{Role: "developer", Content: commandValue}
			fmt.Println("successfully added system instruction")
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
