package infrastructure

import "log"

type LogController struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
}

var LogSysterm LogController
