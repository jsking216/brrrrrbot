package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type BrrrrrBot struct {
	bot reddit.Bot
}

func (r *BrrrrrBot) Comment(p *reddit.Comment) error {
	fmt.Println(fmt.Sprintf("Examining comment body: %s", p.Body))
	matched, _ := regexp.MatchString(`(.*?\d){3}`, p.Body)
	if strings.Contains(p.Body, "SPY") && matched {
		<-time.After(1 * time.Second)
		fmt.Println(fmt.Sprintf("REPLYING to comment body: %s at link %s", p.Body, p.Permalink))
		err := r.bot.Reply(p.Name, "Money printer go BRRRBRRRBRRR")
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to reply to comment due to error: %s", err.Error()))
			return err
		}
	}
	return nil
}

func main() {
	bot, err := reddit.NewBotFromAgentFile("creds", 1*time.Minute)
	if err != nil {
		fmt.Println("Failed to create new bot from file", err)
		return
	}

	subredditsToWatch := []string{
		"robinhood",
		"options",
		"investing",
		"thewallstreet",
		"tradevol",
	}

	cfg := graw.Config{SubredditComments: subredditsToWatch}
	handler := &BrrrrrBot{bot: bot}

	if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
		fmt.Println("Failed to start graw run: ", err)
	} else {
		fmt.Println("graw run failed: ", wait())
	}
}
