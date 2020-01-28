package models

import (
	"errors"
	"fmt"
)

// User type
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers implementation
func GetUsers() []*User {
	return users
}

// AddUser implementation
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include an ID or must me different than 0")
	}

	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}

// GetUserByID implementation
func GetUserByID(ID int) (*User, error) {
	for _, u := range users {
		if u.ID == ID {
			return u, nil
		}
	}

	return &User{}, fmt.Errorf("User with ID '%v' not found", ID)
}

// UpdateUser implementation
func UpdateUser(ID int, u User) (User, error) {
	for i, user := range users {
		if ID == user.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// DeleteUserByID implementation
func DeleteUserByID(ID int) error {
	for i, user := range users {
		if user.ID == ID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", ID)
}
