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

## Serviced

This is info for myself about how I run it on my machine.

Service file can be found here:
```
/lib/systemd/system/jazz-reader.service
```
Should look something like
```
[Unit]
Description=Webhook

[Service]
User=g
Group=g
Restart=on-failure
WorkingDirectory=/home/g/code/go/src/github.com/casperin/jazz_reader
ExecStart=/home/g/code/go/src/github.com/casperin/jazz_reader/jazzd

[Install]
WantedBy=multi-user.target
```
After adding this file (we can do it manually in a clean install) we run
```
systemctl enable jazz-reader.service
```
and it'll symlink it up for us. From then on, we can use service to start and stop it as we please.

To "deploy" locally, do
```
go build cmd/jazzd/*.go
sudo service jazz-reader restart
```

Logs can be found with
```
sudo journalctl -u jazz-reader
# Or running logs
sudo journalctl -f -u jazz-reader
```
