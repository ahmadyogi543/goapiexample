package main

import "strings"

func ValidateUserForm(username, password string) map[string]string {
	errs := map[string]string{}

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	if len(username) < 1 {
		errs["username"] = "username is empty"
	} else if len(username) < 8 {
		errs["username"] = "username is too short"
	}

	if len(password) < 1 {
		errs["password"] = "password is empty"
	} else if len(password) < 8 {
		errs["password"] = "password is too short"
	}

	return errs
}
