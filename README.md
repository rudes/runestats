# RuneScape Stats
[![Build Status](https://travis-ci.org/rudes/runestats.svg?branch=master)](https://travis-ci.org/rudes/runestats)
[![Code Climate](https://codeclimate.com/github/rudes/runestats/badges/gpa.svg)](https://codeclimate.com/github/rudes/runestats)
[![GoDoc](https://godoc.org/github.com/rudes/runestats?status.svg)](https://godoc.org/github.com/rudes/runestats)
[![Docker Automated build](https://img.shields.io/docker/automated/rudes/runestats.svg?maxAge=2592000?style=plastic)](https://hub.docker.com/r/rudes/runestats)
[![Coverage Status](https://coveralls.io/repos/github/rudes/runestats/badge.svg?branch=master)](https://coveralls.io/github/rudes/runestats?branch=master)

[Runescape](http://www.runescape.com/) [Stat](http://runestats.stream/) page for twitch.

If you're a streamer, to use runestats simply create a new panel on your
twitch stream, and type the url in.
    http://runestats.stream/your_username.png
That's it, we'll take care of the rest.

If you're looking to host your own instance,
runestats is best served with [docker](https://docker.com/)
```bash
docker pull rudes/runestats
docker run -d -p 8080:8080 --name=runestats rudes/runestats
```

## Shout Outs

Testing with [Nurrivia](https://www.twitch.tv/nurrivia) on his old school
character [Niriviaa](http://services.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws?user1=niriviaa)

## Licensing

Runestats is licensed under the Apache License, Version 2.0. 
See [LICENSE](https://github.com/rudes/runestats/blob/master/LICENSE) for the full text.
