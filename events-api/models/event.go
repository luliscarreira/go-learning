package models

import (
	"time"

	"example.com/exercises/events-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	query := `
    INSERT INTO events(name, description, location, date_time, user_id)
    VALUES(?, ?, ?, ?, ?)
    `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	// e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventById(id int64) (Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return event, err
	}

	return event, nil
}

func UpdateEvent(id int64, event Event) error {
	query := `
    UPDATE events
    SET name = ?, description = ?, location = ?, date_time = ?
    WHERE id = ?
    `

	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (event Event) Update() error {
	query := `
    UPDATE events
    SET name = ?, description = ?, location = ?, date_time = ?
    WHERE id = ?
    `

	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := db.DB.Exec(query, event.ID)
	if err != nil {
		return err
	}

	return nil
}
