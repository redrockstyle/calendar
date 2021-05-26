package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var eventsSQLSchema = `
CREATE TABLE IF NOT EXISTS events (
	id  INTEGER NOT NULL PRIMARY KEY,
	start TIMESTAMP,
	end TIMESTAMP,
	duration VARCHAR(256),
	name VARCHAR(256),
	description TEXT
)
`

type Event struct {
	Id          int    `db:"id"`
	Start       string `db:"start"`
	End         string `db:"end"`
	Duration    string `db:"duration"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type EventFilterCriterion struct {
	startDay time.Time
	endDay   time.Time
}

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository() (*EventRepository, error) {
	sqlDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	if _, err = sqlDB.Exec(eventsSQLSchema); err != nil {
		return nil, err
	}

	repository := EventRepository{
		db: sqlDB,
	}
	return &repository, nil
}

func (r *EventRepository) Filter(start time.Time, end time.Time) ([]Event, error) {
	stmt, err := r.db.Prepare("SELECT * FROM events WHERE start BETWEEN ? AND ?")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(start.Format("2006-01-02 15:04"), end.Format("2006-01-02 15:04"))
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	collection := make([]Event, 0)
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.Id, &event.Start, &event.End, &event.Duration, &event.Name, &event.Description)
		if err != nil {
			return nil, fmt.Errorf("can't scan struct: %w", err)
		}

		startTime, _ := time.Parse(time.RFC3339, event.Start)
		event.Start = startTime.Format("2006-01-02 15:04")

		endTime, _ := time.Parse(time.RFC3339, event.End)
		event.End = endTime.Format("2006-01-02 15:04")

		collection = append(collection, event)
	}

	return collection, nil
}

func (r *EventRepository) Create(event Event) error {
	stmt, err := r.db.Prepare("INSERT INTO events(start, end, duration, name, description) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.Start, event.End, event.Duration, event.Name, event.Description)

	if err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) Update(event Event) error {
	stmt, err := r.db.Prepare("UPDATE events SET start = ?, end = ?, duration = ?, name = ?, description = ? WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.Start, event.End, event.Duration, event.Name, event.Description, event.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM events WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
