package logger

import (
	"bufio"
	"github.com/galagoshin-com/GoUtils/events"
	"os"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Execute     func(string, []string)
}

var commands = make(map[string]Command)

func RegisterCommand(command Command) {
	if len(commands) == 0 {
		go runCommandListener()
	}
	commands[command.Name] = command
	if len(command.Aliases) > 0 {
		for _, alias := range command.Aliases {
			commands[alias] = command
		}
	}
	events.CallAllEvents(RegisterCommandEvent, command)
}

func doCommand(input string) {
	args := strings.Split(input, " ")
	if len(args) > 0 {
		cmd, ex := commands[args[0]]
		if ex {
			var argsArray []string
			if len(args) > 1 {
				argsArray = strings.Split(input[len(args[0])+1:], " ")
			}
			events.CallAllEvents(CommandPreProccessEvent, args[0], argsArray)
			cmd.Execute(args[0], argsArray)
		} else {
			CommandInfo("Unknow command.")
		}
	}
}

func runCommandListener() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			doCommand(scanner.Text())
		}
	}
}

func GetAllCommands() map[string]Command {
	res := make(map[string]Command)
	for alias, cmd := range commands {
		res[alias] = cmd
	}
	return res
}
