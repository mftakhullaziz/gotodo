package helpers

import (
	"github.com/gorilla/mux"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	PanicIfError(err)

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

func LoggingMiddleware(next http.Handler) http.Handler {
	log := LoggerParent()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request Handler: %s %s %s", r.Host, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func LogRoutes(router *mux.Router) {
	log := LoggerParent()
	log.Infoln("Registered Handler Router:")
	// Walk through all the registered routes and log their respective URLs
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, _ := route.GetMethods()
		path, _ := route.GetPathTemplate()
		log.Println(methods, path)
		return nil
	})

	if err != nil {
		log.Infoln("Error Walking Routes:", err)
	}
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
