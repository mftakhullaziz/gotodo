package logger

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	error2 "gotodo/internal/utils/errors"
	"net/http"
	"os"
	"time"
)

func LoggerParent() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	})
	log.SetOutput(os.Stdout)

	dir, err := os.Getwd()
	error2.PanicIfError(err)

	logPath := dir + "/logs/application.log"
	fileHook := lfshook.NewHook(lfshook.PathMap{
		logrus.InfoLevel:  logPath,
		logrus.ErrorLevel: logPath,
		logrus.DebugLevel: logPath,
		logrus.PanicLevel: logPath,
		logrus.WarnLevel:  logPath,
		logrus.TraceLevel: logPath,
	}, &logrus.JSONFormatter{})

	// Add the file hook to the logger
	log.AddHook(fileHook)
	return log
}

func LoggerQueryInit(db *gorm.DB) {
	log := LoggerParent()
	// Set up a logger to print SQL statements
	newLogger := logger.New(log, logger.Config{
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
		log.Infof("Received request: %s %s", r.Method, r.URL.Path)
		// Call the next handler
		next.ServeHTTP(w, r)
		// Log response details
		log.Infof("Sent response: %s %s", r.Method, r.URL.Path)
	})
}
