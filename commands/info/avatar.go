package info

import (
	"github.com/bwmarrin/discordgo"
)

func Avatar(s *discordgo.Session, m *discordgo.Message, global bool, user string) {
	switch global {
    case false:
        switch {
        case user == "": // &avatar
            reqUser, _ := s.GuildMember(m.GuildID, m.Author.ID)
            s.ChannelMessageSendReply(m.ChannelID, reqUser.AvatarURL("1024"), m.Reference())
            break
        case len(m.Mentions) == 1: // &avatar <@ID>
            reqUser, _ := s.GuildMember(m.GuildID, m.Mentions[0].ID)
            s.ChannelMessageSendReply(m.ChannelID, reqUser.AvatarURL("1024"), m.Reference())
            break
        default: // &avatar ID | &avatar ???
            if reqUser, ok := s.GuildMember(m.GuildID, user); ok == nil {
                s.ChannelMessageSendReply(m.ChannelID, reqUser.AvatarURL("1024"), m.Reference())
            } else {
                s.ChannelMessageSendReply(m.ChannelID, "argument is not a user", m.Reference())
            }
        }
        break
    case true:
        switch {
        case user == "":
            s.ChannelMessageSendReply(m.ChannelID, m.Author.AvatarURL("1024"), m.Reference())
            break
        case len(m.Mentions) == 1:
            s.ChannelMessageSendReply(m.ChannelID, m.Mentions[0].AvatarURL("1024"), m.Reference())
            break
        default:
            if reqUser, ok := s.User(user); ok == nil {
                s.ChannelMessageSendReply(m.ChannelID, reqUser.AvatarURL("1024"), m.Reference())
            } else {
                s.ChannelMessageSendReply(m.ChannelID, "argument is not a user", m.Reference())
            }
        }
    }
}