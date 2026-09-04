package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Task struct {
	ID   int64
	desc string
	done bool
}

type Store struct {
	conn *sql.DB
}

func (s *Store) Init() error {
	var err error
	// TODO: Use better solution for database location
	s.conn, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return err
	}

	createTableStmt := `CREATE TABLE IF NOT EXISTS tasks (
		id integer not null primary key,
		desc text not null,
		done bool
	);`

	if _, err = s.conn.Exec(createTableStmt); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetTasks() ([]Task, error) {
	rows, err := s.conn.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.desc, &task.done)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Store) UpsertTask(task Task) error {
	if task.ID == 0 {
		// TODO: Use better randomization
		task.ID = time.Now().UTC().UnixNano()
	}

	upsertQuery := `INSERT INTO tasks (id, title, body)
	VALUES (?, ?, ?)
	ON CONFLICT(id) DO UPDATE
	SET desc=excluded.title, done=excluded.done`

	if _, err := s.conn.Exec(upsertQuery, task.ID, task.desc, task.ID); err != nil {
		return err
	}

	return nil
}
