package main

import (
	"fmt"
	"strings"
)

var allCommands = [1]string{"\\sysinstr - set the assistants system instructions which can define behavior and personality"}

func ListCommands() {
	fmt.Println("--- Commands --- ")
	fmt.Println("commands have a key/value pair format. So for example, if you wanted to change the system instructions for the assistant, the command would be \\sysinstr=\"Be very informative in your responses.\"")
	for _, v := range allCommands {
		fmt.Println("\n\t" + v)
	}
}

func ExecuteCommand(str string) {
	if str == "\\" {
		ListCommands()
		return
	}

	command := strings.Split(str[1:], "=")
	if len(command) == 1 {
		fmt.Println("invalid command format (missing \"=\")")
		return
	}

	var commandKey, commandValue string
	commandKey = command[0]
	commandValue = command[1]

	if commandValue == "" {
		fmt.Println("you must assign a value to the command")
		return
	}

	switch commandKey {
	case "sysinstr":
		{
			fmt.Println("sstem instruc")
			break
		}

	default:
		{
			break
		}
	}
}
