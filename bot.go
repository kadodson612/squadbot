package main

import (
    "fmt"
    "regexp"
    "strings"
    "io/ioutil"
    "os"

    "github.com/nlopes/slack"
    "github.com/peterhellberg/giphy"
)

//MAIN FUNCTION -- instantiate bot
func main() {

    // fetch token file
    data, err := ioutil.ReadFile("token")
    if err != nil {
        return
    }

    token := string(data)
    token = strings.TrimSuffix(token, "\n")

    api := slack.New(token)
    rtm := api.NewRTM()
    go rtm.ManageConnection()

Loop:
    for {
        select {

        case msg := <-rtm.IncomingEvents:

            switch ev := msg.Data.(type) {

            case *slack.MessageEvent:
                text := ev.Text
                text = strings.TrimSpace(text)
                text = strings.ToLower(text)

                check_messages(text, ev, rtm)

            case *slack.RTMError:
                fmt.Printf("Error: %s\n", ev.Error())

            case *slack.InvalidAuthEvent:
                fmt.Printf("Invalid credentials")
                break Loop

            default:
                // Take no action

            }
        }
    }

}

func check_messages(text string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check if the message string exactly matches "meow"
    info := rtm.GetInfo()
    matched, _ := regexp.MatchString("^meow$", text)
    gif_needed, _ := regexp.MatchString("^gif .*",text)

    if ev.User != info.User.ID && matched {
        rtm.SendMessage(rtm.NewOutgoingMessage("Is that Maru I hear?", ev.Channel))
    }

    if gif_needed {
        text = text[3:]
        g := giphy.DefaultClient
        res, err := g.Search([]string{text})
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        if len(res.Data) > 0 {
            rtm.SendMessage(rtm.NewOutgoingMessage(res.Data[0].BitlyGifURL, ev.Channel))
        }
    }


}
