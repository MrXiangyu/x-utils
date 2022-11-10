// @Title  
// @Description  
// @Author  MrXiang 2022/11/8 14:35
package log

type Level uint32

const (
	// FatalLevel level. Logs and then calls `os.Exit(1)`.
	FatalLevel Level = iota + 1
	// PanicLevel level. Logs and then calls panic with the message passed to
	// Debug, Info, ...
	PanicLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

// Defines the key when adding errors using WithError.
var basicSeparator = " "

func (lvl Level) String() string {
	if lvl == FatalLevel {
		return "FATAL"
	} else if lvl == PanicLevel {
		return "PANIC"
	} else if lvl == ErrorLevel {
		return "ERROR"
	} else if lvl == WarnLevel {
		return "WARN"
	} else if lvl == InfoLevel {
		return "INFO"
	} else if lvl == DebugLevel {
		return "DEBUG"
	} else {
		return ""
	}
}
