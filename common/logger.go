package common

import (
	"fmt"
	"github.com/lifangjunone/go-micro/utils"
	"github.com/phachon/go-logger"
	"path"
)

var (
	LoggerObj *Logger
)

type LoggerOutMode int32

const (
	LoggerConsole     = 0
	LoggerFile        = 1
	LoggerConsoleName = "console"
	LoggerFileName    = "file"
)

type Logger struct {
	OutMode   LoggerOutMode
	FileName  string
	LoggerObj *go_logger.Logger
}

func NewLogger(outMode LoggerOutMode, fileName string) *Logger {
	return &Logger{
		OutMode:  outMode,
		FileName: fileName,
	}
}

func (l *Logger) Config() {
	logger := go_logger.NewLogger()
	switch l.OutMode {
	case LoggerConsole:
		consoleConfig := &go_logger.ConsoleConfig{
			Color:      true, // Does the text display the color
			JsonFormat: true, // Whether to formatted into a JSON string
			Format:     "",   // JsonFormat is false, logger message output to console format string
		}
		logger.Attach(LoggerConsoleName, go_logger.LOGGER_LEVEL_INFO, consoleConfig)
	case LoggerFile:
		currentPath := utils.GetCurrentAbPathByCaller()
		fileName := fmt.Sprintf("%s.log", l.FileName)
		currentFilePath := path.Join(currentPath, "../logs/", fileName)

		fileConfig := &go_logger.FileConfig{
			Filename: currentFilePath, // The file name of the logger output, does not exist automatically
			//LevelFileName: map[int]string{
			//	// If you want to separate separate logs into files, configure LevelFileName parameters.
			//	logger.LoggerLevel("error"): "./error.log", // The error level log is written to the error.log file.
			//	logger.LoggerLevel("info"):  "./info.log",  // The info level log is written to the info.log file.
			//	logger.LoggerLevel("debug"): "./debug.log", // The debug level log is written to the debug.log file.
			//},
			MaxSize:    1024 * 1024, // File maximum (KB), default 0 is not limited
			MaxLine:    100000,      // The maximum number of lines in the file, the default 0 is not limited
			DateSlice:  "d",         // Cut the document by date, support "Y" (year), "m" (month), "d" (day), "H" (hour), default "no".
			JsonFormat: true,        // Whether the file data is written to JSON formatting
			Format:     "",          // JsonFormat is false, logger message written to file format string
		}
		logger.Attach(LoggerFileName, go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
	}
	l.LoggerObj = logger
}
