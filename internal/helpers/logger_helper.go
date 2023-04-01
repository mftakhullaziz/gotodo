package helpers

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func LoggerParent() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	})
	log.SetOutput(os.Stdout)
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
