# GO JAZZ READER

I couldn't find an RSS reader that I liked, so I wrote my own to learn myself
some Golang for great good.

It's a webserver, but assumes only one user. So no user handling, etc.

To set it up, it needs access to an instance of Postgres, with credentials that
it picks up from `~/.jazz_reader/config.json`. Here is mine (slightly edited):
```js
{
    "app_secret": "[super secret!]", // used to encrypt cookie
    "db_user": "postgres", // defaults to "jazzreader_dev"
    "db_pass": "postgres", // defaults to "jazzreader_dev"
    "db_name": "jazzreader_dev", // required
    "password": "[password]", // to enter the app
    // how often should the app check for new entries in each rss
    "sync_rss_every_x_minute": 60,
    "domain": "http://where-you-run-this.com"
}
```

There are currently zero tests. Go me! :)
