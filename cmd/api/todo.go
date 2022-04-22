package main

import (
	"fmt"
	"net/http"

	"github.com/ZweWT/Go-Test.git/internal/data"
	"github.com/ZweWT/Go-Test.git/internal/validator"
)

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"title"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	todo := &data.Todo{
		Name:   input.Name,
		UserID: user.ID,
	}

	v := validator.New()

	if data.ValidateTodo(v, todo); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Todo.Insert(todo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
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

func (app *application) listTodoHandler(w http.ResponseWriter, r *http.Request) {

	todos, err := app.models.Todo.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.successResponse(w, r, http.StatusOK, todos, nil)
}
