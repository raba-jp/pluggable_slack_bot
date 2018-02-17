package main

type Plugin interface {
	CheckMessage(m *Message) bool
	DoAction(m *Message)
}
