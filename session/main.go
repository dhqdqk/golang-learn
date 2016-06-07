package main

import (
	"session"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
}
