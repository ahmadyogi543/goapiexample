package main

import (
	"errors"
	"net/http"
	"strconv"
)

func (s *APIServer) HomeHandler(w http.ResponseWriter, r *http.Request) error {
	return WriteHttpOKJSON(w, APIResponse{
		Message: "Welcome to Stupid API!",
	})
}

func (s *APIServer) GetAllUserHandler(w http.ResponseWriter, r *http.Request) error {
	users := s.UsersModel.GetAll()

	return WriteHttpOKJSON(w, APIResponse{
		Message: "list of all users",
		Data:    users,
	})
}

func (s *APIServer) GetOneUserHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return ErrBadRequestResponse("id is in invalid format", nil)
	}

	user, err := s.UsersModel.GetOne(id)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return ErrNotFoundResponse("user record is not found")
		}
		return err
	}

	return WriteHttpOKJSON(w, APIResponse{
		Message: "list of all users",
		Data:    user,
	})
}

func (s *APIServer) AddUserHandler(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return ErrBadRequestResponse("failed to parse form", nil)
	}

	username := r.Form.Get("username")
	pasword := r.Form.Get("password")

	errs := ValidateUserForm(username, pasword)
	if len(errs) > 0 {
		return ErrBadRequestResponse("user form is invalid", errs)
	}

	user := s.UsersModel.Add(username, pasword)
	return WriteHttpOKJSON(w, APIResponse{
		Message: "user is created",
		Data:    user,
	})
}

func (s *APIServer) UpdateUserHandler(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return ErrBadRequestResponse("failed to parse form", nil)
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return ErrBadRequestResponse("id is in invalid format", nil)
	}
	username := r.Form.Get("username")
	pasword := r.Form.Get("password")

	errs := ValidateUserForm(username, pasword)
	if len(errs) > 0 {
		return ErrBadRequestResponse("user form is invalid", errs)
	}

	err = s.UsersModel.Update(id, username, pasword)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return ErrNotFoundResponse("user record is not found")
		}
		return err
	}

	return WriteHttpNoContentJSON(w)
}

func (s *APIServer) DeleteUserHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return ErrBadRequestResponse("id is in invalid format", nil)
	}

	err = s.UsersModel.Delete(id)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return ErrNotFoundResponse("user record is not found")
		}
		return err
	}

	return WriteHttpNoContentJSON(w)
}
