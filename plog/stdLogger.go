package plog

/**********************************************
** @Des: stdLogger
** @Author: zhangxueyuan 
** @Date:   2018-09-11 10:43:47
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-09-11 10:43:47
***********************************************/

import "log"

var (
	StdLogger *Logger = newLogger(log.Lshortfile, 3)
)

func ERRORF(format string, v ...interface{}) {
	StdLogger.ERRORF(format, v...)
}

func WARNF(format string, v ...interface{}) {
	StdLogger.WARNF(format, v...)
}

func INFOF(format string, v ...interface{}) {
	StdLogger.INFOF(format, v...)
}

func DEBUGF(format string, v ...interface{}) {
	StdLogger.DEBUGF(format, v...)
}

func ERROR(v ...interface{}) {
	StdLogger.ERROR(v...)
}

func WARN(v ...interface{}) {
	StdLogger.WARN(v...)
}

func INFO(v ...interface{}) {
	StdLogger.INFO(v...)
}

func DEBUG(v ...interface{}) {
	StdLogger.DEBUG(v...)
}
