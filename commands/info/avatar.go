package info

import (
	"flag"

	"github.com/google/shlex"
	"github.com/bwmarrin/discordgo"
)

var AvatarInfo = CommandInfo{
	Description: "Replies with requested user's avatar if specified, otherwise replies with author's. Flag -g to specify global avatar within a server.",
	Usage:       "avatar | av [ -g ] [ @user | ID ]",
}

func Avatar(s *discordgo.Session, m *discordgo.MessageCreate, args string) {

	input, err := shlex.Split(args)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, "AMBATUKAAAAAAAAM!", m.Reference())
		return
	}

	if len(input) > 2 {
		s.ChannelMessageSendReply(m.ChannelID, "Too many arguments supplied.", m.Reference())
	}

	/*
	switch *global {
	case false:
		switch {
		case len(set.Args()) == 0:
			// &avatar
			if authorMember, ok := s.GuildMember(m.GuildID, m.Author.ID); ok == nil {
				s.ChannelMessageSendReply(m.ChannelID, authorMember.AvatarURL("1024"), m.Reference())
			}
			break
		default:
			if len(m.Mentions) > 0 {
				// &avatar <@USERID>
				requestedUser, _ := s.GuildMember(m.GuildID, m.Mentions[0].ID)
				s.ChannelMessageSendReply(m.ChannelID, requestedUser.AvatarURL("1024"), m.Reference())
			} else if requestedUser, ok := s.GuildMember(m.GuildID, set.Args()[0]); ok == nil {
				// &avatar USERID
				s.ChannelMessageSendReply(m.ChannelID, requestedUser.AvatarURL("1024"), m.Reference())
			} else {
				// &avatar ???
				s.ChannelMessageSendReply(m.ChannelID, "Invalid user supplied.", m.Reference())
			}
		}
	default:
		switch {
		case len(set.Args()) == 0:
			// &avatar -g
			s.ChannelMessageSendReply(m.ChannelID, m.Author.AvatarURL("1024"), m.Reference())
			break
		default:
			if len(m.Mentions) == 1 {
				// &avatar -g <@USERID>
				s.ChannelMessageSendReply(m.ChannelID, m.Mentions[0].AvatarURL("1024"), m.Reference())
			} else if requestedUser, ok := s.User(set.Args()[0]); ok == nil {
				// &avatar -g USERID
				s.ChannelMessageSendReply(m.ChannelID, requestedUser.AvatarURL("1024"), m.Reference())
			} else {
				// &avatar -g ???
				s.ChannelMessageSendReply(m.ChannelID, "Invalid user supplied.", m.Reference())
			}
		}
	}
	*/

	/* QUARANTINED
	switch len(args) {
	case 0:
		if authorMember, ok := s.GuildMember(m.GuildID, m.Author.ID); ok == nil {
			s.ChannelMessageSendReply(m.ChannelID, authorMember.AvatarURL("1024"), m.Reference())
		} else {
			s.ChannelMessageSendReply(m.ChannelID, m.Author.AvatarURL("1024"), m.Reference())
		}

		break
	case 1:
		if args[0] == "-g" {
			s.ChannelMessageSendReply(m.ChannelID, m.Author.AvatarURL("1024"), m.Reference())
		} else if len(m.Mentions) > 0 && string(args[0][1]) == "@" {
			requestedUser, _ := s.GuildMember(m.GuildID, m.Mentions[0].ID)
			s.ChannelMessageSendReply(m.ChannelID, requestedUser.AvatarURL("1024"), m.Reference())
		} else if requestedUser, ok := s.GuildMember(m.GuildID, args[0]); ok == nil {
			s.ChannelMessageSendReply(m.ChannelID, requestedUser.AvatarURL("1024"), m.Reference())
		} else {
			s.ChannelMessageSendReply(m.ChannelID, "Invalid user supplied.", m.Reference())
		}

		break
	case 2:
		if args[1] == "-g" {
			if len(m.Mentions) > 0 && string(args[0][1]) == "@" {
				s.ChannelMessageSendReply(m.ChannelID, m.Mentions[0].AvatarURL("1024"), m.Reference())
			} else if requestedUser, ok := s.User(args[0]); ok == nil {
				s.ChannelMessageSendReply(m.ChannelID, requestedUser.AvatarURL("1024"), m.Reference())
			} else {
				s.ChannelMessageSendReply(m.ChannelID, "Please specify a valid user!", m.Reference())
			}
		} else {
			if len(m.Mentions) > 0 {
				s.ChannelMessageSendReply(m.ChannelID, "Invalid user supplied.", m.Reference())
			} else {
				s.ChannelMessageSendReply(m.ChannelID, "Invalid flag supplied.", m.Reference())
			}
		}

		break
	case 3:
		s.ChannelMessageSendReply(m.ChannelID, "Too many arguments supplied.", m.Reference())
	}
	*/
}
