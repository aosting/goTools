package plog

/**********************************************
** @Des: logger
** @Author: zhangxueyuan 
** @Date:   2018-09-11 10:43:17
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-09-11 10:43:17
***********************************************/

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	LevelDebug         = iota
	LevelInformational
	LevelWarning
	LevelError
)

var Level = LevelInformational

type Logger struct {
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	debug *log.Logger
	depth int
}

func newLogger(flag int, depth int) *Logger {
	Logger := newLogger3(os.Stdout, flag, depth)
	return Logger
}

func newLogger3(w io.Writer, flag int, depth int) *Logger {
	logger := new(Logger)
	logger.depth = depth
	if logger.depth <= 0 {
		logger.depth = 2
	}

	logger.err = log.New(os.Stderr, "[ERROR] ", flag)
	logger.warn = log.New(w, "[WARN] ", flag)
	logger.info = log.New(w, "[INFO] ", flag)
	logger.debug = log.New(w, "[DEBUG] ", flag)

	logger.SetLevel(LevelInformational)

	return logger
}

func (ll *Logger) SetLevel(l int) {
	Level = l
}

func SetLevel(l int) {
	Level = l
}

// 统一设置日志前缀
func (ll *Logger) SetPrefix(prefix string) {
	ll.err.SetPrefix("[ERROR] " + prefix)
	ll.warn.SetPrefix("[WARN] " + prefix)
	ll.info.SetPrefix("[INFO] " + prefix)
	ll.debug.SetPrefix("[DEBUG] " + prefix)
}

func (ll *Logger) ERROR(v ...interface{}) {
	if LevelError < Level {
		return
	}
	ll.err.Output(ll.depth, fmt.Sprintln(v...))

}

func (ll *Logger) WARN(v ...interface{}) {

	if LevelWarning < Level {
		return
	}
	ll.warn.Output(ll.depth, fmt.Sprintln(v...))

}

func (ll *Logger) INFO(v ...interface{}) {
	if LevelInformational < Level {
		return
	}
	ll.info.Output(ll.depth, fmt.Sprintln(v...))

}

func (ll *Logger) DEBUG(v ...interface{}) {
	if LevelDebug < Level {
		return
	}
	ll.debug.Output(ll.depth, fmt.Sprintln(v...))

}

func (ll *Logger) ERRORF(format string, v ...interface{}) {
	if LevelError < Level {
		return
	}

	fmt.Println(format)
	ll.err.Output(ll.depth, fmt.Sprintf(format, v...))

}

func (ll *Logger) WARNF(format string, v ...interface{}) {
	if LevelWarning < Level {
		return
	}
	ll.warn.Output(ll.depth, fmt.Sprintf(format, v...))

}

func (ll *Logger) INFOF(format string, v ...interface{}) {
	if LevelInformational < Level {
		return
	}
	ll.info.Output(ll.depth, fmt.Sprintf(format, v...))

}

func (ll *Logger) DEBUGF(format string, v ...interface{}) {
	if LevelDebug < Level {
		return
	}
	ll.debug.Output(ll.depth, fmt.Sprintf(format, v...))

}

func (ll *Logger) SetFlag(flag int) {
	ll.err.SetFlags(flag)
	ll.warn.SetFlags(flag)
	ll.debug.SetFlags(flag)
}
