gmusic.go
=========

A command-line remote for the unofficial [Google Play Music Desktop Player](https://github.com/MarshallOfSound/Google-Play-Music-Desktop-Player-UNOFFICIAL-).

## Usage

```
Usage:
  gmusic <command>

Commands:
  back                  Play the previous track in the playlist.
  hate                  Hate (thumbs-down) the currently playing track.
  love                  Love (thumbs-up) the currently playing track.
  next                  Play the next track in the playlist.
  pause                 Pause, but do not play if already paused.
  play                  Play, but do not pause if already playing.
  playpause             Toggle play/pause.
```

## Installation

### Download

Binaries for Linux, MacOS and Windows are available via [Releases](https://github.com/f1337/gmusic.go/releases).

### Homebrew (macOS)

```
brew tap f1337/gmusic
brew install gmusic
```

### Source

```
go get github.com/f1337/gmusic
```

## FAQ

OK, the "frequently" part is a bit misleading. "Asked Questions to which I replied with an essay over chat" is probably more accurate.

> Why type an entire command in to terminal when you can hit a single key?

My original reason for creating this program is for automation: When my RSI timer tells me to take a break, it runs a script that invokes this CLI with `gmusic pause`. I now have several “pause my music when I do X” workflows.

Context-switching is my single greatest productivity loss while coding. For me, typing `gmusic love` or `gmusic hate` *is* faster than switching contexts, finding like/dislike, clicking it, and switching back.

I don't use media hotkeys, and media hotkeys don't typically cover like/dislike. If a native API exists for rating, one could theoretically map a hotkey to like and dislike. Or one could skip the "if a native API exists" exercise, and map hotkeys to `gmusic love` and `gmusic hate`.

Those are my (original) use cases, and those use-cases drove the application development. Which is why `play <track>` isn’t supported yet: I browse/search via the GUI, not via CLI. But being able to rate, skip, pause, resume, etc. without leaving my coding context is a big win for me.
