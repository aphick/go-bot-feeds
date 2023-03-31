# go-bot-feeds

RSS, Atom and JSON feed parser plugin for go-chat-bot

Requires on the `FEEDS`, and `FEEDS_CRON` environment variables to be set.

`FEEDS` is a comma seperated list of [gofeed](https://github.com/mmcdole/gofeed) compatible feeds (RSS, ATOM, JSON).

`FEEDS_CRON` is the frequency of polling in cron syntax.
