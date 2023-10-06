package log

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Error(err error) {
	logType := color.RedString("ERROR")
	errorMessage := color.RedString(err.Error())
	fmt.Fprintf(os.Stderr, "[%s] %s\n", logType, errorMessage)
}

func Critical(err error) {
	logType := color.HiMagentaString("CRITICAL")
	errorMessage := color.RedString(err.Error())
	fmt.Fprintf(os.Stderr, "[%s] %s\n", logType, errorMessage)
	panic(err)
}

func Infof(format string, a ...any) {
	logType := color.HiCyanString("INFO")
	fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...any) {
	logType := color.HiYellowString("WARN")
	fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
}

func Debugf(isDebug bool, format string, a ...any) {
	if isDebug {
		logType := color.HiGreenString("DEBUG")
		fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
	}
}

func Info(format string, a ...interface{}) {
	logType := color.HiCyanString("INFO")
	message := fmt.Sprintf(format, a...)
	fmt.Printf("[%s] %s\n", logType, message)
}

func Warn(format string, a ...interface{}) {
	logType := color.HiYellowString("WARN")
	message := fmt.Sprintf(format, a...)
	fmt.Printf("[%s] %s\n", logType, message)
}

func Debug(isDebug bool, format string, a ...interface{}) {
	if isDebug {
		logType := color.HiGreenString("DEBUG")
		message := fmt.Sprintf(format, a...)
		fmt.Printf("[%s] %s\n", logType, message)
	}
}
