package main

import (
	"bufio"
	"fmt"
	"os"
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
