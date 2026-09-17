package command

import (
    "fmt"
    "strings"

	"avustaja/commands/info"

	"github.com/alecthomas/kong"
	"github.com/bwmarrin/discordgo"
)

func Dispatch(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {

    var helpOut strings.Builder
    writtenHelp := false 

	opts := []kong.Option{ // array of Apply(k *Kong) error interfaces
		//kong.NoDefaultHelp(),
		kong.Exit(func(int) {}),
		kong.Bind(s, m.Message),
        kong.Writers(&helpOut, &helpOut),
        kong.Help(func(options kong.HelpOptions, ctx *kong.Context) error {
        if writtenHelp {
            return nil
        }
        writtenHelp = true
        return kong.DefaultHelpPrinter(options, ctx)}),
	}

	var commands struct {
		List   ListCommand   `cmd:"" help:"Serves a list of bot commands."`
		Avatar AvatarCommand `cmd:"" help:"Serves a requested user's avatar."`
	}

	parser, err := kong.New(&commands, opts...)
	if err != nil {
		return err
	}

	ctx, err := parser.Parse(args)
    if helpOut.Len() > 0 {
        s.ChannelMessageSendReply(m.ChannelID, helpOut.String(), m.Reference())
        return nil
    }
    
	if err != nil {
		return err
	}

	if err := ctx.Run(); err != nil {
		return err
	}

    fmt.Println("Parsing concluded.")
	return nil
}

type ListCommand struct {
	Input string `arg:"" help:"Nothing"`
}

type AvatarCommand struct {
	Global bool   `help:"Serves global avatar." short:"g" long:"global"`
	User   string `arg:"" optional:"" help:"Requested user."`
}

func (opts *ListCommand) Run(s *discordgo.Session, m *discordgo.Message) error {
	// implement list
    fmt.Println("List")
    return nil
}

func (opts *AvatarCommand) Run(s *discordgo.Session, m *discordgo.Message) error {
    info.Avatar(s, m, opts.Global, opts.User)
    
    return nil
}