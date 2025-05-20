package main

import (
	"net/http"
)

type APIServer struct {
	UsersModel UsersModel
	*http.Server
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		UsersModel: UsersModel{
			Users: []User{},
		},
		Server: &http.Server{
			Addr: addr,
		},
	}
}

func (s *APIServer) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", MakeHandleFunc(s.HomeHandler))
	mux.HandleFunc("GET /users", MakeHandleFunc(s.GetAllUserHandler))
	mux.HandleFunc("GET /users/{id}", MakeHandleFunc(s.GetOneUserHandler))
	mux.HandleFunc("POST /users", MakeHandleFunc(s.AddUserHandler))
	mux.HandleFunc("PATCH /users/{id}", MakeHandleFunc(s.UpdateUserHandler))
	mux.HandleFunc("DELETE /users/{id}", MakeHandleFunc(s.DeleteUserHandler))

	return mux
}
