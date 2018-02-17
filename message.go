package main

import (
	"github.com/nlopes/slack"
)

type Message struct {
	From    string
	Channel string
	Text    string
}

type messageHandler struct {
	client       *slack.Client
	messageQueue chan *Message
	plugins      []Plugin
	channelID    string
}

func NewMessage(ev *slack.MessageEvent) *Message {
	return &Message{
		From:    ev.User,
		Channel: ev.Channel,
		Text:    ev.Msg.Text,
	}
}

func CopyMessage(m *Message, t string) *Message {
	return &Message{
		From:    m.From,
		Channel: m.Channel,
		Text:    t,
	}
}

func newMessageHandler(cl *slack.Client, q chan *Message, p []Plugin, ch string) *messageHandler {
	return &messageHandler{
		client:       cl,
		messageQueue: q,
		plugins:      p,
		channelID:    ch,
	}
}

func (h *messageHandler) handle() {
	for {
		select {
		case m := <-h.messageQueue:
			h.execPlugins(m)
		}
	}
}

func (h *messageHandler) execPlugins(m *Message) {
	for _, p := range h.plugins {
		if !p.CheckMessage(m) {
			continue
		}
		p.DoAction(m)
		break
	}
}
