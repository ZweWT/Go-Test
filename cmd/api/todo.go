package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZweWT/Go-Test.git/internal/data"
)

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"title"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	todo := data.Todo{
		ID:   id,
		Name: "Gopppp",
	}

	app.successResponse(w, r, http.StatusOK, todo, nil)
}
