package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const maxLogLines = 100

type LevelFileHook struct {
	Writer     io.Writer
	Formatter  logrus.Formatter
	LogLevels  []logrus.Level
	ShouldExec func() bool
}

func (h *LevelFileHook) Levels() []logrus.Level {
	return h.LogLevels
}

func (h *LevelFileHook) Fire(entry *logrus.Entry) error {
	if h.ShouldExec != nil && !h.ShouldExec() {
		return nil
	}

	dataBytes, err := h.Formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = h.Writer.Write(dataBytes)
	if err != nil {
		return err
	}

	return nil
}

func LoggerParent() *Logger {
	logs := logrus.New()

	dir, err := os.Getwd()
	PanicIfError(err)

	loggerDir := dir + "/logs/application.log"

	if isRunningInTest() {
		// Condition Testing environment will be not record logger test
		logs.SetOutput(ioutil.Discard) // Discard log output when not running tests
		_, err := os.OpenFile(loggerDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			if !os.IsNotExist(err) {
				// Error other than "no such file or directory", log the error
				logs.WithError(err).Warn("Failed to open log file. Logging to console.")
			} else {
				// File does not exist, log the message
				logs.Warn("Log file does not exist. Logging to console.")
			}
			// Set clean console
			//nullFile, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)
			//if err != nil {
			//	log.Fatal(err)
			//}
			// Redirect os.Stdout to /dev/null
			//os.Stdout = nullFile
		}
	} else {
		fmt.Println("Testing run")
		logs.SetOutput(ioutil.Discard)

		// Create a custom formatter for file logging
		fileFormatter := &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
			FullTimestamp:   true,
			ForceColors:     true,
		}

		// Create a file hook to write logs to the log file
		file, err := os.OpenFile(loggerDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		PanicIfError(err)

		// Create the custom hook with desired log levels and formatter
		fileHook := &LevelFileHook{
			Writer:     file,
			Formatter:  fileFormatter,
			LogLevels:  []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel},
			ShouldExec: func() bool { return true },
		}

		// Add the file hook to the logger
		logs.AddHook(fileHook)

		// Check log file size and reset if necessary
		ResetLogFileIfNeeded(loggerDir)
	}

	// Create a console logger that writes to os.Stdout
	consoleLog := logrus.New()
	consoleLog.SetOutput(os.Stdout)

	// Return a custom struct that encapsulates both loggers
	return &Logger{
		Log:        logs,
		ConsoleLog: consoleLog,
	}
}

type Logger struct {
	Log        *logrus.Logger
	ConsoleLog *logrus.Logger
}

func (l *Logger) Println(args ...interface{}) {
	// Print to both log and console
	l.Log.Println(args...)
	l.ConsoleLog.Println(args...)
}

func LoggerQueryInit(db *gorm.DB) {
	log := LoggerParent()
	// Set up a logger to print SQL statements
	newLogger := logger.New(log.Log, logger.Config{
		SlowThreshold:             time.Second, // Log slow queries
		LogLevel:                  logger.Info, // Log SQL Statement
		IgnoreRecordNotFoundError: true,        // Ignore "not found" errors
		Colorful:                  true,        // Enable colorful output
	})
	db.Logger = newLogger
}

// LoggerMiddleware function to log requests and responses
func LoggerMiddleware(next http.Handler) http.Handler {
	log := LoggerParent()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details
		log.Log.Infof("Received request: %s %s", r.Method, r.URL.Path)
		// Call the next handler
		next.ServeHTTP(w, r)
		// Log response details
		log.Log.Infof("Sent response: %s %s", r.Method, r.URL.Path)
	})
}

// Start Login define is Running test
func isTestFile(file string) bool {
	matched, _ := regexp.MatchString(".+_test\\.go", file)
	return matched
}

func getCallerFileName() string {
	// Get the caller's PC (Program Counter)
	pc, _, _, _ := runtime.Caller(1)
	// Get the corresponding function
	fn := runtime.FuncForPC(pc)
	// Get the file name and line number
	file, _ := fn.FileLine(pc)
	// Normalize the file path
	file = filepath.Base(file)
	return file
}

func isRunningInTest() bool {
	if os.Getenv("ENV") == "test" {
		return true
	}
	pc, _, _, _ := runtime.Caller(1)
	caller := runtime.FuncForPC(pc)
	if caller == nil {
		return false
	}
	file := getCallerFileName()
	file = strings.ToLower(file)
	return isTestFile(file)
}

// End Logic

func ResetLogFileIfNeeded(logPath string) {
	fileInfo, err := os.Stat(logPath)
	if err != nil {
		return // Unable to get file info, skip resetting log file
	}

	if fileInfo.Size() > 0 && fileInfo.Size() >= maxLogLines*calculateAverageLogLineSize() {
		err := os.Truncate(logPath, 0)
		if err != nil {
			logrus.Warn("Failed to reset log file:", err)
		}
	}
}

func calculateAverageLogLineSize() int64 {
	// Calculate average log line size based on previous log entries
	// Return an estimated size in bytes
	// You can customize this implementation based on your log format and typical log message lengths
	return 100 // Assuming an average log line size of 100 bytes
}
