package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func JsonStatus(message string) []byte {
	m, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
	return m
}

func JsonError(w http.ResponseWriter, err error) {
	log.Printf("ERROR: %v", err)
	io.WriteString(w, string(JsonStatus("fail")))
	return
}
