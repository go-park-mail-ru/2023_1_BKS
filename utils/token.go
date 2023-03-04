package utils

import (
	"encoding/json"
	"net/http"
)

// Utils.go contain handy utils functions to build json
// messages and return a json response.

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
