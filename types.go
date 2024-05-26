package main

import (
	"net/http"

	"github.com/google/uuid"
)

type Review struct {
	Id          uuid.UUID `json:"id"`
	MovieName   string    `json:"movie_name"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
}

type CreateReview struct {
	MovieName   string `json:"movie_name"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}

type MyError struct {
	Error string `json:"message"`
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func NewReview(movieName string, rating int, desc string) *Review {
	return &Review{
		Id:          uuid.New(),
		MovieName:   movieName,
		Rating:      rating,
		Description: desc,
	}
}
