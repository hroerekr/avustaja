package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"avustaja/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/google/shlex"
)

var (
	config map[string]string
)

func init() {
	configFile, err := os.ReadFile(".\\config.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(configFile, &config)
}

func main() {
	s, err := discordgo.New("Bot " + config["token"])
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	err = s.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	s.AddHandler(dispatchCommands)

	s.Identify.Intents = discordgo.IntentsGuildMessages

	s.UpdateWatchStatus(-1, "for &")

	fmt.Println("Avustaja is now up, C-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	s.Close()
}

func dispatchCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	// string -> []byte is O(n)
	if m.Author.Bot || m.Content == "" || string(m.Content[0]) != config["prefix"] || m.GuildID == "" {
		return
	}

	//cmd, _, _ := strings.Cut(strings.TrimPrefix(m.Content, config["prefix"]), " ")

	//_, args, _ := strings.Cut(m.Content, " ")

	args, _ := shlex.Split(strings.TrimPrefix(m.Content, config["prefix"]))

	fmt.Println(args)
	if err := command.Dispatch(s, m, args); err != nil {
		s.ChannelMessageSendReply(m.ChannelID, err.Error(), m.Reference())
	}
}
