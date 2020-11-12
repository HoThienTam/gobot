package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "~") {
		if m.Content == "cuc" {
			s.ChannelMessageSend(m.ChannelID, "cac")
		}
	}

}
