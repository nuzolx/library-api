package router

import (
	"log"
	"net/http"
	"time"

	"github.com/nuzolx/library-api/app"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func errorHandler(handler Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			if appErr, ok := err.(*app.AppError); ok {
				http.Error(w, appErr.Message, appErr.Code)
				return
			} else {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	}
	return http.HandlerFunc(fn)
}
