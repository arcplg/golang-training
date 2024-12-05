package repository

import (
	"database/sql"
	"time"
	model "to-do-app/internal/models"
)

type toDoRepository struct {
	db *sql.DB
}

// Hàm tạo mới repository
func NewToDoRepository(db *sql.DB) ToDoRepository {
	return &toDoRepository{db: db}
}

// Implement GetAll
func (r *toDoRepository) GetAll() ([]model.ToDo, error) {
	var todos []model.ToDo
	rows, err := r.db.Query("SELECT id, title, description, is_completed, user_id, created_at, updated_at FROM to_do_list")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.ToDo
		rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.UserId, &todo.CreatedAt, &todo.UpdatedAt)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *toDoRepository) Create(todo model.ToDo) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO to_do_list (title, description, is_completed) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(todo.Title, todo.Description, todo.IsCompleted)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *toDoRepository) GetDetailById(id int) (*model.ToDo, error) {
	var todo model.ToDo
	var createdAt []byte
	var updatedAt []byte

	err := r.db.QueryRow("SELECT id, title, description, is_completed, user_id, created_at, updated_at FROM to_do_list WHERE id = ?", id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.UserId, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	// Convert byte slice to time.Time
	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt)) // Adjust format based on your DB format
	if err != nil {
		return nil, err
	}
	// Convert byte slice to time.Time
	parsedTime2, err := time.Parse("2006-01-02 15:04:05", string(updatedAt)) // Adjust format based on your DB format
	if err != nil {
		return nil, err
	}

	todo.CreatedAt = parsedTime
	todo.UpdatedAt = parsedTime2

	return &todo, nil
}

func (r *toDoRepository) Update(toDo *model.ToDo) error {
	query := `UPDATE to_do_list SET title = ?, description = ?, is_completed = ? WHERE id = ?`
	_, err := r.db.Exec(query, toDo.Title, toDo.Description, toDo.IsCompleted, toDo.ID)
	return err
}

func (r *toDoRepository) Delete(id int) error {
	query := "DELETE from to_do_list where id = ?"
	_, err := r.db.Exec(query, id)

	return err
}
