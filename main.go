package main

import (
  "flag"
  "log"

  "github.com/f1337/gmusic.go/gpmdp"
)


const VERSION = "0.0.1"


func main() {
  host := flag.String("host", "localhost:5672", "websocket service address")
  flag.Parse()
  command := flag.Arg(0)
  log.SetFlags(0)

  done := make(chan struct{})

  client, err := gpmdp.NewClient.Connect(*host)
  if err != nil {
    panic(err)
  }

  defer client.Disconnect(done)
  client.ReadMessages(done)


  // TODO: implement the rest of the GMusic API
  // cf: https://github.com/gmusic-utils/gmusic.js#documentation
  switch command {
  case "back":
    client.Back()
  case "hate":
    client.Hate()
  case "love":
    client.Love()
  case "next":
    client.Next()
  case "pause":
    client.Pause()
  case "play":
    client.Play()
  case "playpause":
    client.PlayPause()
  }
}
