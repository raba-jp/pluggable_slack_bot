package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
)

type Bot struct {
	client    *slack.Client
	plugins   []Plugin
	id        string
	channelID string
}

// BotParams is Bot object initializer params
type BotParams struct {
	BotToken          string
	BotID             string
	ChannelID         string
	Port              string
	VerificationToken string
}

func New(params BotParams) *Bot {
	return &Bot{
		client:    slack.New(token),
		plugins:   []Plugin{},
		id:        ID,
		channelID: channelID,
	}
}

func (b *Bot) Run() {
	// Listening slack event and response
	log.Printf("[INFO] Start slack event listening")

	l := newSlackListener(b.client, b.id, b.ChannelID)
	go l.listen()

	queue := l.messageQueue
	messageHandler := newMessageHandler(b.client, queue, b.plugins, b.ChannelID)
	go messageHandler.handle()

	http.Handle("/interaction", interactionHandler{
		verificationToken: env.VerificationToken,
	})

	if err := http.ListenAndServe(":"+b.port, nil); err != nil {
		panic("")
	}
}

func (b *Bot) AddPlugin(p Plugin) {
	b.plugins = append(b.plugins, p)
}

func (b *Bot) PostMessage(m *Message) {
	b.client.PostMessage(m.Channel, m.Text, slack.PostMessageParameters{})
}

func (b *Bot) PostMessageWithAttachment(m *Message) {
	// TODO
	panic("Not Implemented yet")
}

func (b *Bot) ReplyMessage(m *Message) {
	text := fmt.Sprintf("<@%s> %s", m.From, m.Text)
	b.client.PostMessage(m.Channel, text, slack.PostMessageParameters{})
}
