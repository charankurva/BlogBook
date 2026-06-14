package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	_ "github.com/lib/pq"
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
func (app *Application) ReadJson(w http.ResponseWriter, r *http.Request, dst any) error {

	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		var syntaxerr *json.SyntaxError
		//var unmarshelltypeerr *json.UnmarshalTypeError
		var invalidunmarshellerr *json.InvalidUnmarshalError
		switch {
		case errors.As(err, &syntaxerr):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxerr.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidunmarshellerr):
			panic(err)
		default:
			return err
		}

	}
	return nil
}
func (app *Application) OpenDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", app.config.data.dsn)
	if err != nil {
		app.logger.Error(err.Error())
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		app.logger.Error(err.Error())
		return nil, err
	}
	app.logger.Info("establised connection to db")
	return db, nil
}
