package feeds

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-chat-bot/bot"
	"github.com/mmcdole/gofeed"
)
// when called, this function polls the feeds and grabs their contents
func feeds(channel string) (msg string, err error) {
	feeds := strings.Split(os.Getenv("FEEDS"), ",")
	fp := gofeed.NewParser()

	var sb strings.Builder

	for _, feed := range feeds {
		pFeed, err := fp.ParseURL(feed)
		if err == nil {
			fItem := pFeed.Items[0]
			fTime := fItem.PublishedParsed
			// Use pubdate of latest entry and if it came after the last update add it to msg string.
			yTime := time.Now().Add(-24*time.Hour)

			if fTime.After(yTime) {
				tStr :=	fmt.Sprintf("*%s* @ %s \n %s - %s\n", pFeed.Title, fTime.Format(time.UnixDate), fItem.Title, fItem.Link)
				sb.WriteString(tStr)
			}
		}
	}

	msg = fmt.Sprintf(sb.String())
	return
}


func init() {
	// A comma separated list of channel ids from environment
	channels := strings.Split(os.Getenv("CHANNEL_IDS"), ",")


	if len(channels) > 0 {
		// Reads from subscribed feeds and posts updates to channels at 1:01 mon-fri
		config := bot.PeriodicConfig{
			CronSpec: "1 1 * * * mon-fri",
			Channels: channels,
			CmdFunc:  feeds,
		}

		bot.RegisterPeriodicCommand("feeds", config)
	}
}
