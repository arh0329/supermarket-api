package logging

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// env variables
const (
	logLevel  = "LOG_LEVEL"
	logFormat = "LOG_FORMAT"
)

// variable for logger
var (
	logger *logrus.Logger
)

func Log() *logrus.Logger {
	return logger
}

func fileFormat(f *runtime.Frame) (string, string) {
	filename := path.Base(f.File)
	function := path.Base(f.Function)
	return fmt.Sprintf("%s()", function), fmt.Sprintf("%s:%d", filename, f.Line)
}

// init function sets up logging for stdout
func init() {
	logger = logrus.New()
	logger.SetReportCaller(true)
	_ = viper.BindEnv(logFormat)
	format := strings.ToLower(viper.GetString(logFormat))
	if format == "text" {
		logger.Formatter = &logrus.TextFormatter{
			CallerPrettyfier: fileFormat,
			FullTimestamp:    true,
		}
	} else {
		logger.Formatter = &logrus.JSONFormatter{
			CallerPrettyfier: fileFormat,
		}
	}
	_ = viper.BindEnv(logLevel)
	level := strings.ToLower(viper.GetString(logLevel))
	switch level {
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

}
