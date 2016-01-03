package utils

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(v)
	return nil
}
