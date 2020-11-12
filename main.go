package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const token string = "NzYzMDY2MTQ4NjAxNDYyNzk0.X3ySug.xL7ipJW--H5eczwjJ9O0OQAmEtU"

func main() {
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
