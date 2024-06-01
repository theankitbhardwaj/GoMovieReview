package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Review struct {
	Id        uuid.UUID `json:"id"`
	MovieName string    `json:"movie_name"`
	Rating    int8      `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateReview struct {
	MovieName string `json:"movie_name"`
	Rating    int8   `json:"rating"`
	Comment   string `json:"comment"`
}

type MyError struct {
	Error string `json:"message"`
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func NewReview(movieName string, rating int8, comment string) *Review {
	return &Review{
		Id:        uuid.New(),
		MovieName: movieName,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now().UTC(),
	}
}
