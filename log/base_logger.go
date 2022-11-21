// @Title  
// @Description  
// @Author  MrXiang 2022/11/8 14:18
package log

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type BaseLogger struct {
	Logger *log.Logger
	Level  Level
	Fields map[string]interface{}
	Ctx    context.Context
}

var ErrorKer = "error"

// Set Log Directory
func (b *BaseLogger) SetOutput(path string) Entry {
	today := time.Now().Format("2006-01-02")
	logPath := fmt.Sprintf("%s%s.log", path, today)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return b
	}
	b.Logger.SetOutput(logFile)
	return b
}

func (b *BaseLogger) WithContext(ctx context.Context) Entry {
	b.Ctx = ctx
	return b
}

func (b *BaseLogger) WithField(key string, value interface{}) Entry {
	return b.WithFields(map[string]interface{}{key: value})
}

func (b *BaseLogger) WithFields(fields map[string]interface{}) Entry {
	data := make(map[string]interface{}, len(b.Fields)+len(fields))
	for k, v := range b.Fields {
		data[k] = v
	}
	for k, v := range fields {
		data[k] = v
	}
	b.Fields = data
	return b
}

func (b *BaseLogger) WithError(err error) Entry {
	return b.WithField(ErrorKer, err.Error())
}

func (b *BaseLogger) SetLevel(lvl Level) {
	b.Level = lvl
}

func (b *BaseLogger) getLevel() Level {
	return b.Level
}

func (b *BaseLogger) Fatal(format string, v ...interface{}) {
	b.log(FatalLevel, format, v)
}

func (b *BaseLogger) Panic(format string, v ...interface{}) {
	b.log(PanicLevel, format, v)
}

func (b *BaseLogger) Error(format string, v ...interface{}) {
	b.log(ErrorLevel, format, v)
}

func (b *BaseLogger) Warn(format string, v ...interface{}) {
	b.log(WarnLevel, format, v)
}

func (b *BaseLogger) Info(format string, v ...interface{}) {
	b.log(InfoLevel, format, v)
}

func (b *BaseLogger) Debug(format string, v ...interface{}) {
	b.log(DebugLevel, format, v)
}

func (b *BaseLogger) log(level Level, format string, v []interface{}) {
	_, file, line, _ := runtime.Caller(2)
	b.WithField("file", fmt.Sprintf("%s, line:%d", file, line))
	if level > b.getLevel() {
		return
	}
	msg := format
	if len(v) > 0 {
		for _, s := range v {
			msg = fmt.Sprintf("%s%s", msg, s)
		}
	}
	b.output(b, level, msg)
}

func (b *BaseLogger) output(entry *BaseLogger, lvl Level, msg string) {
	var sb strings.Builder
	b.WithField("level", lvl.String())
	b.WithField("time", time.Now().Format("2006-01-02 15:04:05"))
	b.WithField("msg", msg)
	if lvl != ErrorLevel {
		delete(entry.Fields, ErrorKer)
	}
	fs, err := json.Marshal(entry.Fields)
	if err == nil {
		sb.Write(fs)
	}
	if lvl == FatalLevel {
		b.Logger.Fatal(sb.String())
	} else if lvl == PanicLevel {
		b.Logger.Panic(sb.String())
	}
	b.Logger.Println(sb.String())
}
