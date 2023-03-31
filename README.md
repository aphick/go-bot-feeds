# go-bot-feeds

RSS, Atom and JSON feed parser plugin for go-chat-bot

Requires on the `FEEDS`, and `FEEDS_CRON` environment variables to be set.

`FEEDS` is a comma seperated list of [gofeed](https://github.com/mmcdole/gofeed) compatible feeds (RSS, ATOM, JSON).

```
FEEDS="https://www.nasa.gov/rss/dyn/breaking_news.rss,https://blog.archive.org/atom"
```

`FEEDS_CRON` is the frequency of polling in cron syntax. *Currently does not support DoW!!*

```
FEEDS_CRON="8 0 * * * *"
```
