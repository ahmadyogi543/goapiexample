package main

import (
	"log"
	"net/http"
)

type APIHandlerFunc func(http.ResponseWriter, *http.Request) error

func MakeHandleFunc(h APIHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("server: %s %s\n", r.Method, r.URL.Path)

		if err := h(w, r); err != nil {
			if apierr, ok := err.(APIErrorResponse); ok {
				WriteJSON(w, apierr.Errors.Status, apierr)
				return
			}
			log.Println(err)
			WriteHttpInternalServerErrorJSON(w)
		}
	}
}
