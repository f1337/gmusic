gmusic.go
=========

A command-line remote for the unofficial [Google Play Music Desktop Player](https://github.com/MarshallOfSound/Google-Play-Music-Desktop-Player-UNOFFICIAL-).

## Usage

```
Usage:
  gmusic <command>
  gmusic -h | --help
  gmusic -v | --version

Options:
  -h, --help            Show this screen.
  -v, --version         Show version.

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
