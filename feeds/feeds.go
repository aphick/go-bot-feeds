package feeds

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-chat-bot/bot"
	"github.com/mmcdole/gofeed"
	"github.com/robfig/cron/v3"
)

// when called, this function polls the feeds and grabs their contents
func feeds(channel string) (msg string, err error) {
        feeds := strings.Split(os.Getenv("FEEDS"), ",")

        cronTime :=  os.Getenv("FEEDS_CRON")
        currTime := time.Now()

        // Parse cron time and figure out last time job would have run.
	// Currently this does not support DoW!! Has to be *
        sched, err := cron.ParseStandard(cronTime)
	if err != nil {
		return err
	}
        nextRun := sched.Next(currTime)
        tDiff := time.Until(nextRun)
        yTime := currTime.Add(-1 * tDiff)
        fp := gofeed.NewParser()

        var sb strings.Builder

        // Iterate through feeds and entries until we should have seen them before
        for _, feed := range feeds {
                pFeed, err := fp.ParseURL(feed)
                if err == nil {
                        for i, _ := range pFeed.Items {
                                fItem := pFeed.Items[i]
                                fTime := fItem.PublishedParsed
                                // Compare published time to last time cron ran
                                if fTime.After(yTime) {
                                        tStr := fmt.Sprintf("*%s* @ %s \n %s - %s\n", pFeed.Title, fTime.Format(time.
UnixDate), fItem.Title, fItem.Link)
                                        sb.WriteString(tStr)
                                } else {
                                        break
                                }
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
		cronTime :=  os.Getenv("FEEDS_CRON")
		// Reads from subscribed feeds and posts updates to channels at 1:01 mon-fri
		config := bot.PeriodicConfig{
			CronSpec: cronTime,
			Channels: channels,
			CmdFunc:  feeds,
		}

		bot.RegisterPeriodicCommand("feeds", config)
	}
}
