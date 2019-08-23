package entity

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response describes HTTP response
type Response struct {
	Status  int
	Headers map[string]string
	Body    map[string]interface{}
}

func (r *Response) Write(w http.ResponseWriter) {
	if r.Status == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if b, err := json.Marshal(r.Body); err == nil {
		if r.Headers != nil {
			h := w.Header()
			for key, value := range r.Headers {
				h.Set(key, value)
			}
		}
		w.WriteHeader(r.Status)
		w.Write(b)
	} else {
		log.Printf("Failed to marshal: %s", err)
	}
}
