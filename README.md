Basic go program that lets you interact with ChatGPT through the terminal. Each chat initiated will store the entire conversations context so you can get answers based on previous questions asked.

## Installation

Simply run the main.go file and a "config.json" will be generated in the root of your project. Add your OpenAI API key to this config and start using! You can also select which model you want to use in the config (defaults to gpt-4.1-nano).

## Commands

You can enter custom commands to customize your experience while chatting with the bot

- **\clear** - Will clear the entire chat conversation you've had with the bot. This will also clear any custom system instructions if you had set one.

- **\model** - Change the chatGPT model that is used during chat. Defaults to gpt-4.1-nano when running app for the first time.

- **\save** - Saves your current conversation with the bot in a json file.

- **\sysinstr** - Set a system instruction for the bot during your conversation. A system instruction is a rule or guideline given to a chatbot that tells it how to behave or respond. Each system instruction given to the bot will persist during the entire conversation unless cleared.

