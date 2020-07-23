package main

import (
	"encoding/json"
	"net/http"
)

//WebResponse is response object
type WebResponse struct {
	Message string `json:message`
}

func writeErrorMsg(w http.ResponseWriter, msg string) {
	w.WriteHeader(404)
	writeJSONResponse(w, WebResponse{msg})
}

func writeJSONResponse(w http.ResponseWriter, obj interface{}) {
	json.NewEncoder(w).Encode(obj)
}

func writeResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

}
