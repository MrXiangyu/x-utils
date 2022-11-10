// @Title  
// @Description  
// @Author  MrXiang 2022/11/8 14:19
package log

import "context"

type (
	Fields map[string]interface{}
	Entry  interface {
		SetOutput(path string) Entry
		SetLevel(lvl Level)
		WithContext(ctx context.Context) Entry
		WithField(key string, value interface{}) Entry
		WithFields(fields map[string]interface{}) Entry
		WithError(err error) Entry
		//
		Fatal(format string, v ...interface{})
		Panic(format string, v ...interface{})
		Error(format string, v ...interface{})
		Warn(format string, v ...interface{})
		Info(format string, v ...interface{})
		Debug(format string, v ...interface{})
	}
)
