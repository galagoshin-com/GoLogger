package logger

import (
	"fmt"
	"github.com/galagoshin-com/GoUtils/events"
	"os"
	"runtime/debug"
	"time"
)

var startTime = time.Now().Format("2006.01.02 15:04:05")

func Print(message string) {
	time := time.Now().Format("2006.01.02 15:04:05")
	msg := fmt.Sprintf("%s* %s[%s] %s> %s%s", colors["IBlue"], colors["IWhite"], time, colors["White"], colors["IWhite"], message)
	fmt.Println(msg)
	writeLog(fmt.Sprintf("[%s] > %s\n", time, message))
	events.CallAllEvents(LogEvent, msg)
}

func Warning(message string) {
	time := time.Now().Format("2006.01.02 15:04:05")
	msg := fmt.Sprintf("%s* %s[%s] %s> %s%s", colors["IYellow"], colors["IWhite"], time, colors["White"], colors["IWhite"], message)
	fmt.Println(msg)
	writeLog(fmt.Sprintf("[%s] WARNING > %s\n", time, message))
	events.CallAllEvents(LogEvent, msg)
}

func Info(message string) {
	msg := fmt.Sprintf("%s* %s%s", colors["IGreen"], colors["IWhite"], message)
	fmt.Println(msg)
	events.CallAllEvents(LogEvent, msg)
}

func CommandInfo(message string) {
	msg := fmt.Sprintf("%s* %s%s", colors["IPurple"], colors["IWhite"], message)
	fmt.Println(msg)
	events.CallAllEvents(LogEvent, msg)
}

func Error(err error) {
	time := time.Now().Format("2006.01.02 15:04:05")
	stack := string(debug.Stack())
	msg := fmt.Sprintf("%s* %s[%s] %s> %s%s\n%s", colors["IRed"], colors["IWhite"], time, colors["White"], colors["IWhite"], err.Error(), stack)
	fmt.Println(msg)
	writeLog(fmt.Sprintf("[%s] ERROR > %s\n%s", time, err.Error(), stack))
	events.CallAllEvents(ErrorEvent, err)
	events.CallAllEvents(LogEvent, msg)
}

func Panic(err error) {
	time := time.Now().Format("2006.01.02 15:04:05")
	stack := string(debug.Stack())
	msg := fmt.Sprintf("%s* %s[%s] %s> %s%s\n%s", colors["Red"], colors["IWhite"], time, colors["White"], colors["IWhite"], err.Error(), stack)
	fmt.Println(msg)
	writeLog(fmt.Sprintf("[%s] PANIC > %s\n%s", time, err.Error(), stack))
	events.CallAllEvents(ErrorEvent, err)
	events.CallAllEvents(PanicEvent, err)
	events.CallAllEvents(LogEvent, msg)
	os.Exit(1)
}

func Debug(lvl int, printStackTrace bool, data ...any) {
	if len(data) == 0 {
		return
	}
	if lvl <= debugLevel {
		var msg string
		for _, el := range data {
			time := time.Now().Format("2006.01.02 15:04:05")
			msg = fmt.Sprintf("%s* %s[%s] %s> %s%+v", colors["ICyan"], colors["IWhite"], time, colors["White"], colors["IWhite"], el)
			fmt.Println(msg)
			writeLog(fmt.Sprintf("[%s] DEBUG > %+v\n", time, el))
		}
		if printStackTrace {
			stack := string(debug.Stack())
			fmt.Println(stack)
			msg += "\n" + stack
			writeLog(stack)
		}
		events.CallAllEvents(LogEvent, msg)
	}
}
