package logger

import (
	"bytes"
	"log"
	"path"
	"runtime"
	"strconv"
)

// create single ton logger with different log level methods/extentions

// print log msg with caller's file name & line# info
func Logger(msg interface{}) {
	_, file, lineNo, ok := runtime.Caller(1)

	suffix := "<suffix>"
	if ok {
		var buffer bytes.Buffer
		buffer.WriteString(" @")
		buffer.WriteString(path.Base(file))
		buffer.WriteString("#")
		buffer.WriteString(strconv.Itoa(lineNo))

		suffix = buffer.String()
	}
	log.Print(msg, suffix)
}
