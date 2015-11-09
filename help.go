package bot

import (
	"fmt"
	"strings"
)

const (
	helpDescripton    = "Description: %s"
	helpUsage         = "Usage: %s%s %s"
	availableCommands = "Available commands: %v"
	helpAboutCommand  = "Type: '%shelp <command>' to see details about a specific command."
	helpCommand       = "help"
)

func help(c *Cmd) {
	cmd := parse(CmdPrefix+c.RawArgs, c.Channel, c.User)
	if cmd == nil {
		showAvailabeCommands(c.Channel, c.User)
		return
	}

	command := commands[cmd.Command]
	if command == nil {
		showAvailabeCommands(c.Channel, c.User)
		return
	}

	showHelp(cmd, command)
}

func showHelp(c *Cmd, help *customCommand) {
	if help.Description != "" {
		handlers.Response(c.Channel, fmt.Sprintf(helpDescripton, help.Description), c.User)
	}
	handlers.Response(c.Channel, fmt.Sprintf(helpUsage, CmdPrefix, c.Command, help.ExampleArgs), c.User)
}

func showAvailabeCommands(channel string, sender *User) {
	var cmds []string
	for k := range commands {
		cmds = append(cmds, k)
	}
	handlers.Response(channel, fmt.Sprintf(helpAboutCommand, CmdPrefix), sender)
	handlers.Response(channel, fmt.Sprintf(availableCommands, strings.Join(cmds, ", ")), sender)
}
