# praat

Discord bot that streams audio from stdin to a voice channel.

## Example

```bash
yt-dlp "https://www.youtube.com/watch?v=dQw4w9WgXcQ" -o - \
    | ffmpeg -i - -f opus - \
    | BOT_TOKEN=`cat .token` praat -c CHANNEL_ID
```

Pass `-af acrusher=.1:1:64:0:log` to ffmpeg for extra fun (I am not responsible for permanent hearing damage).

## Author

Kees van Voorthuizen

## License

[GPLv3](./LICENSE)
