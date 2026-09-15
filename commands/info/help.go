package info

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/shlex"
)

type CommandInfo struct {
	Description string
	Usage       string
}

var HelpInfo = CommandInfo{
	Description: "Replies with a full list of categories, commands within those categories, or info about a particular command.",
	Usage:       "help [ category | command ]",
}

func Help(s *discordgo.Session, m *discordgo.MessageCreate, args string) {
	categories := map[string]string{
		"info": "help, avatar",
	}

	commands := map[string]CommandInfo{
		"help":   HelpInfo,
		"avatar": AvatarInfo,
		"av":     AvatarInfo,
	}

	//var flagSet *flag.FlagSet

	//flagSet = flag.NewFlagSet("Help Flags", flag.ContinueOnError)

	input, _ := shlex.Split(args)

	//article := flagSet.String("article", "", "User-supplied command or category thereof.")

	if len(input) > 1 {
		s.ChannelMessageSendReply(m.ChannelID, "You have supplied too many arguments.", m.Reference())
		return
	}

	//flagSet.Parse(input)
	if len(args) == 0 {
		s.ChannelMessageSendReply(m.ChannelID, "...", m.Reference())
	} else if cat, ok := categories[input[0]]; ok {
		s.ChannelMessageSendReply(m.ChannelID, "```"+input[0]+": "+cat+"```", m.Reference())
	} else if cmd, ok := commands[input[0]]; ok {
		s.ChannelMessageSendReply(m.ChannelID, input[0]+":\n```"+cmd.Description+"```\nUsage:\n```"+cmd.Usage+"```", m.Reference())
	} else {
		s.ChannelMessageSendReply(m.ChannelID, "Unrecognised argument.", m.Reference())
	}
}
