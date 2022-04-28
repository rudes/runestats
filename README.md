# RuneScape Stats

Static [Runescape](http://www.runescape.com/) [Stat](http://runestats.stream/) images for embedding.

```
http://runestats.stream/your_username.png
```
That's it, we'll take care of the rest.

If you're looking to host your own instance,
runestats is best served with [docker](https://docker.com/)
```bash
docker pull rudes/runestats
docker run -d -p 8080:8080 --name=runestats rudes/runestats
```

## Shout Outs

Kyle Stanley for fixing that annoying rendering bug

## Licensing

Runestats is licensed under the Apache License, Version 2.0. 
See [LICENSE](https://github.com/rudes/runestats/blob/master/LICENSE) for the full text.
 
