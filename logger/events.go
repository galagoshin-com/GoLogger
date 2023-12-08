package logger

import "github.com/galagoshin-com/GoUtils/events"

const (
	LogEvent                = events.EventName("LogEvent")
	ErrorEvent              = events.EventName("ErrorEvent")
	RegisterCommandEvent    = events.EventName("RegisterCommandEvent")
	CommandPreProccessEvent = events.EventName("CommandPreProccessEvent")
	PanicEvent              = events.EventName("PanicEvent")
)
