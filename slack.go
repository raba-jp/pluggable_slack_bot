package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
)

type slackListener struct {
	messageQueue chan *Message
	client       *slack.Client
	botID        string
	channelID    string
}

func newSlackListener(cl *slack.Client, b string, ch string) *slackListener {
	return &slackListener{
		MessageQueue: make(chan *Message),
		client:       cl,
		botID:        b,
		channelID:    ch,
	}
}

// LstenAndResponse listens slack events and response
// particular messages. It replies by slack message button.
func (s *slackListener) listen() {
	rtm := s.client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if !s.validateMessageEvent(ev) {
				return // no op
			}
			s.MessageQueue <- NewMessage(ev)
		}
	}
}

func (s *slackListener) validateMessageEvent(ev *slack.MessageEvent) bool {
	// Only response in specific channel. Ignore else.
	if ev.Channel != s.channelID {
		log.Printf("%s %s", ev.Channel, ev.Msg.Text)
		return false
	}

	// Only response mention to bot. Ignore else.
	if !strings.HasPrefix(ev.Msg.Text, fmt.Sprintf("<@%s> ", s.botID)) {
		return false
	}

	return true
}
