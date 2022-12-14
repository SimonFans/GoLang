package Method

import "fmt"

// LogLevel is a logging level
type LogLevel uint8

// Possible log levels
const (
	DebugLevel LogLevel = iota + 1
	WarningLevel
	ErrorLevel
)

// String implements the fmt.Stringer interface
func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	}

	return fmt.Sprintf("unknown log level: %d", l)
}
