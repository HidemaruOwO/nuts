package log

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Infof(format string, a ...any) {
	logType := color.HiCyanString("INFO")
	fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
}

func Error(err error) {
	logType := color.RedString("ERROR")
	errorMessage := color.RedString(err.Error())
	fmt.Fprintf(os.Stderr, "[%s] %s", logType, errorMessage)
}

func ErrorExit(err error) {
	Error(err)
	panic(err)
}

func Warnf(format string, a ...any) {
	logType := color.HiYellowString("WARN")
	fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
}

func Criticalf(format string, a ...any) {
	logType := color.HiMagentaString("CRITICAL")
	fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
}

func Debugf(isDebug bool, format string, a ...any) {
	if isDebug {
		logType := color.HiGreenString("DEBUG")
		fmt.Printf("[%s] %s", logType, fmt.Sprintf(format, a...))
	}
}
