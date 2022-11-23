package logger

import (
	"os"
)

var writingLogs = true

func writeLog(log string) {
	if !writingLogs {
		return
	}
	if _, err := os.Stat("logs"); err != nil {
		err = os.Mkdir("logs", 0777)
		if err != nil {
			Error(err)
		}
	}
	logFile, err := os.OpenFile("logs/"+startTime+".log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err == nil {
		_, err = logFile.WriteString(log)
		if err != nil {
			Error(err)
		}
		err = logFile.Close()
		if err != nil {
			Error(err)
		}
	} else {
		logFile, err := os.Create("logs/" + startTime + ".log")
		if err != nil {
			Error(err)
		}
		_, err = logFile.WriteString(log)
		if err != nil {
			Error(err)
		}
		err = logFile.Close()
		if err != nil {
			Error(err)
		}
	}
}

func SetLogs(value bool) {
	writingLogs = value
}
