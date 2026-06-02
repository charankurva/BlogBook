package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) WriteJson(w http.ResponseWriter, r *http.Request, status int, data any, headers http.Header) error {
	res, err := json.Marshal(data)
	if err != nil {
		app.logger.Error(err.Error())
	}
	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
	return err

}
