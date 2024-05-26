package main

import (
	"github.com/google/uuid"
)

type Review struct {
	Id          uuid.UUID `json:"id"`
	MovieName   string    `json:"movie_name"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
}

type MyError struct {
	Error string `json:"message"`
}
