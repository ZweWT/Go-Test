package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	//overwriting the built-in notfound and method not allowed error messages from httprouter
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/index", app.statusHandler)
	router.HandlerFunc(http.MethodPost, "/v1/api/todo", app.createTodoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/api/todo/:id", app.showTodoHandler)

	return router
}
