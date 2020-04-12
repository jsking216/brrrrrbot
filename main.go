package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type WSBBot struct {
	bot reddit.Bot
}

func (r *WSBBot) Comment(p *reddit.Comment) error {
	matched, _ := regexp.MatchString(`(.*?\d){3, }`, p.Body)
	if strings.Contains(p.Body, "SPY") && matched {
		<-time.After(10 * time.Second)
		fmt.Println(fmt.Sprintf("Replying to comment body: %s at link %s", p.Body, p.Permalink))
		return r.bot.Reply(p.Name, "Money printer go BRRRBRRRBRRR")
	}
	return nil
}

func main() {
	bot, err := reddit.NewBotFromAgentFile("creds", 1*time.Minute)
	if err != nil {
		fmt.Println("Failed to create new bot from file", err)
		return
	}

	cfg := graw.Config{SubredditComments: []string{"wallstreetbets"}}
	handler := &WSBBot{bot: bot}

	if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
		fmt.Println("Failed to start graw run: ", err)
	} else {
		fmt.Println("graw run failed: ", wait())
	}
}
