# GoLogger
Functional logger for creating software in Golang.
## Install

`go get -t github.com/Galagoshin/GoLogger`

## How to use

### Logger

#### Print
```go
logger.Print("Message")
```

#### Warning
```go
logger.Warning("Message")
```

#### Error
```go
logger.Error(err)
```

#### Panic
```go
logger.Panic(err) //like Error(), but with exit
```

#### Info
```go
logger.Info("Message") //without writing to logs
```

#### Command info
```go
logger.CommandInfo("Executed command") //use after command executed.
```

#### Debug
```go
logger.Debug(0, true, var1, var2)
```
where

`0` - debug level

`true` - Print stack trace

`var1, var2, var3...` - data to debug

#### Debug level
```go
logger.SetDebugLevel(1)
```

#### Disable/enable logs files.
```go
logger.SetLogs(false) //disable
logger.SetLogs(true) //enable
```

### Commands

#### Example usage
```go
logger.RegisterCommand(logger.Command{Name: "stop", Description: "Stop application.", Aliases: []string{"shutdown"}, Execute: stop})

func stop(string, []string) {
	os.Exit(1)
}
```
and type `stop` or `shutdown` in console to stop application.

### Events

#### LogEvent
```go
events.RegisterEvent(events.Event{
    Name: logger.LogEvent,
    Execute: OnLogEvent,
})

func OnLogEvent(args ...any) {
    msg := args[0].(string)
    fmt.Println(msg)
}
```

#### ErrorEvent
```go
events.RegisterEvent(events.Event{
    Name: logger.ErrorEvent,
    Execute: OnErrorEvent,
})

func OnErrorEvent(args ...any) {
    err := args[0].(error)
    fmt.Println(err.Error())
}
```

#### PanicEvent
```go
events.RegisterEvent(events.Event{
    Name: logger.PanicEvent,
    Execute: OnPanicEvent,
})

func OnPanicEvent(args ...any) {
    err := args[0].(error)
    fmt.Println(err.Error())
}
```

#### RegisterCommandEvent
```go
events.RegisterEvent(events.Event{
    Name: logger.RegisterCommandEvent,
    Execute: OnRegisterCommandEvent,
})

func OnRegisterCommandEvent(args ...any) {
    cmd := args[0].(logger.Command)
    fmt.Println("Register command", cmd.Name)
}
```

#### CommandPreProccessEvent
```go
events.RegisterEvent(events.Event{
    Name: logger.CommandPreProccessEvent,
    Execute: OnCommandPreProccessEvent,
})

func OnCommandPreProccessEvent(args ...any) {
    cmd := args[0].(string)
    cmd_args := args[1].([]string)
    fmt.Println("Input:", cmd, cmd_args)
}
```