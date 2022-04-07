package httpapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP handlers http request.
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := f(w, r.WithContext(r.Context())); err != nil {
		handleError(w, err)
	}
}

func handleError(w http.ResponseWriter, err error) {
	var statusErr Error
	if !errors.As(err, &statusErr) {
		log.Printf("failed request: %s", err)
		statusErr = ErrUnhealthy
	}

	body, err := json.Marshal(statusErr)
	if err != nil {
		log.Printf("failed to marshal status error: %s", err)
		w.WriteHeader(ErrUnhealthy.Status())
		return
	}

	w.WriteHeader(statusErr.Status())
	if _, err := w.Write(body); err != nil {
		log.Printf("failed to write status error: %s", err)
		return
	}
}
