package main

import (
	"errors"
	"slices"
)

var ErrRecordNotFound = errors.New("models: record not found")

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

type UsersModel struct {
	Users   []User
	Counter int
}

func (m *UsersModel) GetAll() []User {
	return m.Users
}

func (m *UsersModel) GetOne(id int) (*User, error) {
	for _, user := range m.Users {
		if id == user.ID {
			return &user, nil
		}
	}

	return nil, ErrRecordNotFound
}

func (m *UsersModel) Add(username, password string) User {
	user := User{
		ID:       m.Counter,
		Username: username,
		Password: password,
		Balance:  0,
	}

	m.Users = append(m.Users, user)
	m.Counter++
	return user
}

func (m *UsersModel) Update(id int, username, password string) error {
	for i := range m.Users {
		if id == m.Users[i].ID {
			m.Users[i].Username = username
			m.Users[i].Password = password
			return nil
		}
	}
	return ErrRecordNotFound
}

func (m *UsersModel) Delete(id int) error {
	for i := range m.Users {
		if id == m.Users[i].ID {
			m.Users = slices.Delete(m.Users, i, i+1)
			return nil
		}
	}
	return ErrRecordNotFound
}
