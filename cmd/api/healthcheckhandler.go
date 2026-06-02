package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) HealthcheckHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//stmt := fmt.Sprintf("version of Blogbook:%s\n Enviroment:%s ", app.config.version, app.config.env)
	data := map[string]string{
		"version": app.config.version,
		"env:":    app.config.env,
	}
	err := app.WriteJson(w, r, http.StatusAccepted, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "the process encouter a problem and doesnot resolved", http.StatusInternalServerError)
	}
}
