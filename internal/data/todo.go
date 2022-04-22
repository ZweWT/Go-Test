package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ZweWT/Go-Test.git/internal/validator"
)

type Todo struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type TodoModel struct {
	DB *sql.DB
}

func (m TodoModel) Insert(todo *Todo) error {

	query := `
		INSERT INTO todos (name, user_id)
		VALUES ($1, $2)
		RETURNING id, created_at`

	args := []interface{}{todo.Name, todo.UserID}

	return m.DB.QueryRow(query, args...).Scan(&todo.ID, &todo.CreatedAt)

}

func (m TodoModel) Get(id int64) (*Todo, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, name, user_id, created_at
		FROM todos
		WHERE id = $1
	`

	var todo Todo

	err := m.DB.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Name,
		&todo.UserID,
		&todo.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &todo, nil
}

func (m TodoModel) GetAll() ([]*Todo, error) {
	query := `SELECT id, name, user_id, created_at FROM todos ORDER BY id`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*Todo{}

	for rows.Next() {
		var todo Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Name,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (m TodoModel) Update(t *Todo) (*Todo, error) {
	return nil, nil
}

func (m TodoModel) Delete(id int64) error {
	return nil
}

func ValidateTodo(v *validator.Validator, todo *Todo) {
	v.Check(todo.Name != "", "name", "is required")
}
