package models

import (
	"errors"

	"example.com/exercises/events-api/db"
	"example.com/exercises/events-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := db.DB.Exec(query, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	// id, err := result.LastInsertId()
	// user.ID = id

	return err
}

func (user User) RetrieveUserId() error {
	return nil
}

// TODO: Refactor this method to  not query the database
func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
