package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var prefix string

func main() {

	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	prefix = os.Getenv("PREFIX")

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, prefix) {
		content := m.Content[len(prefix):len(m.Content)]

		args := strings.Fields(content)
		if len(args) != 2 {
			s.ChannelMessageSend(m.ChannelID, "Lệnh sai thằng ngu ạ!")
			return
		}
		cmd := strings.ToLower(args[0])

		if cmd == "check" {
			// r, _ := http.PostForm("https://trade.vndirect.com.vn/chung-khoan/danh-muc",
			// 	url.Values{"searchMarketStatisticsView.symbol": {args[1]}})
			// bytes, _ := ioutil.ReadAll(r.Body)

			msg := &discordgo.MessageEmbed{Title: "Bảng giá mã " + strings.ToUpper(args[1])}

			s.ChannelMessageSendEmbed(m.ChannelID, msg)
		}
	}

}
