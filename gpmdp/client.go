package gpmdp

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "log"
  "net/url"
  "os"
  "strings"
  "time"

  "github.com/gorilla/websocket"
  homedir "github.com/mitchellh/go-homedir"
)


type Client struct {
  IsPlaying   string

  conn        *websocket.Conn
  token       string
}


var NewClient = &Client{

}


type ws_message struct {
  Namespace string    `json:"namespace"`
  Method    string    `json:"method"`
  Arguments []string  `json:"arguments"`
}



func (gpmdp *Client) Authorize(token string) error {
  arguments := []string{"go-gpmdp-remote"}

  if token != "" {
    arguments = append(arguments, token)
  }

  message := &ws_message{
  Namespace: "connect",
  Method: "connect",
  Arguments: arguments}

  // log.Printf("Authorize: %#v", message)
  err := gpmdp.conn.WriteJSON(message)
  return err
}



func (gpmdp *Client) Back() error {
  message := &ws_message{
  Namespace: "playback",
  Method: "rewind"}

  log.Printf("Back: %#v", message)
  err := gpmdp.conn.WriteJSON(message)
  return err
}



func (gpmdp *Client) Connect(host string) (*Client, error) {
  u := url.URL{Scheme: "ws", Host: host, Path: "/"}
  c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
  gpmdp.conn = c
  return gpmdp, err
}



func (gpmdp *Client) Disconnect(done chan struct{}) error {
  err := gpmdp.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

  select {
  case <-done:
  case <-time.After(time.Second):
  }

  gpmdp.conn.Close()

  // log.Println("Disconnect OK")
  return err
}



func (gpmdp *Client) Hate() error {
  message := &ws_message{
  Namespace: "rating",
  Method: "toggleThumbsDown"}

  // log.Printf("Hate: %#v", message)
  err := gpmdp.conn.WriteJSON(message)
  return err
}



func (gpmdp *Client) Love() error {
  message := &ws_message{
  Namespace: "rating",
  Method: "toggleThumbsUp"}

  // log.Printf("Love: %#v", message)
  err := gpmdp.conn.WriteJSON(message)
  return err
}



func (gpmdp *Client) Next() error {
  message := &ws_message{
  Namespace: "playback",
  Method: "forward"}

  // log.Printf("Next: %#v", message)
  err := gpmdp.conn.WriteJSON(message)
  return err
}



func (gpmdp *Client) Pause() error {
  var err error

  if gpmdp.IsPlaying == "true" {
    err = gpmdp.PlayPause()
    log.Println("Stopped Google Play Music Desktop Player.")
  } else {
    log.Println("Google Play Music Desktop Player was not running.")
  }

  return err
}



func (gpmdp *Client) Play() error {
  var err error

  if gpmdp.IsPlaying == "true" {
    log.Println("Google Play Music Desktop Player is already playing.")
  } else {
    err = gpmdp.PlayPause()
    log.Println("Started Google Play Music Desktop Player.")
  }

  return err
}



func (gpmdp *Client) PlayPause() error {
  message := &ws_message{
  Namespace: "playback",
  Method: "playPause"}

  // log.Printf("PlayPause: %#v", message)
  err := gpmdp.conn.WriteJSON(message)
  return err
}



func (gpmdp *Client) ReadMessage() (string, interface{}, error) {
  var dat map[string]interface{}

  err := gpmdp.conn.ReadJSON(&dat)

  channel := fmt.Sprintf("%v", dat["channel"])
  payload := dat["payload"]

  return channel, payload, err
}



func (gpmdp *Client) ReadMessages(done chan struct{}) {
  var err error
  playState := make(chan string)
  token := make(chan string)

  go func() {
    token <- gpmdp.readPermanentToken()

    defer close(done)
    for {
      var channel string
      var payload interface{}

      channel, payload, err = gpmdp.ReadMessage()
      if err != nil {
        // log.Printf("ReadMessage %#s", err)
        return
      }
      // log.Printf("recv: %#v", dat)

      switch channel {
      case "connect":
        token <- gpmdp.parseTokenResponse(payload.(string))
      // case "library":
      // case "playlists":
      case "playState":
        playState <- fmt.Sprintf("%v", payload)
      // case "queue":
      // case "time":
      // default:
      // 	log.Printf("default: %#v", dat)
      }
    }
  }()

  var ps, tk string
  for {
    // wait until we know the playState,
    // and have a permanent auth token
    if ps != "" && len(tk) > 4 {
      break
    }

    select {
    case ps = <-playState:
      gpmdp.IsPlaying = ps
    case tk = <-token:
      err = gpmdp.Authorize(tk)
      gpmdp.check(err)
    }
  }
}



// private methods
func (gpmdp *Client) parseTokenResponse (payload string) (string) {
  var code string
  var err error

  if payload == "CODE_REQUIRED" {
    code, err = gpmdp.promptForCode()
  } else if len(payload) > 4 {
    code = payload
    err = gpmdp.savePermanentToken(code)
  }

  gpmdp.check(err)
  return code
}



func (gpmdp *Client) promptForCode () (string, error) {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter code displayed in Google Play Music Desktop Player: ")
  code, err := reader.ReadString('\n')

  return strings.TrimRight(code, "\r\n"), err
}

func (gpmdp *Client) check (e error) {
  if e != nil {
    panic(e)
  }
}

func (gpmdp *Client) readPermanentToken () (string) {
  path, _ := homedir.Expand("~/.go-gpmdp-remote")

  token, err := ioutil.ReadFile(path)
  if err != nil {
    return ""
  }

  return string(token)
}

func (gpmdp *Client) savePermanentToken (token string) error {
  path, _ := homedir.Expand("~/.go-gpmdp-remote")
  err := ioutil.WriteFile(path, []byte(token), 0644)
  // log.Printf("token saved: %#v", token)
  return err
}
