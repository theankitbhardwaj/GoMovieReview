package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type Storage interface {
	CreateReview(*Review) *MyError
	DeleteReview(uuid.UUID) *MyError
	GetReviewByID(uuid.UUID) (*Review, *MyError)
	GetReviews() (*[]Review, *MyError)
	UpdateReview(*Review) *MyError
}

type MyStore struct {
	tempDB []Review
}

func NewMyStore() (*MyStore, error) {
	tempDB := make([]Review, 0)
	return &MyStore{tempDB: tempDB}, nil
}

func (s *MyStore) setupDB() error {
	for i := 0; i < 10; i++ {
		review := Review{
			Id:          uuid.New(),
			MovieName:   fmt.Sprintf("Movie %v", i),
			Rating:      rand.Intn(6),
			Description: fmt.Sprintf("Description of review for movie %v", i),
		}
		s.tempDB = append(s.tempDB, review)
	}

	return nil
}

func (s *MyStore) GetReviews() (*[]Review, *MyError) {
	return &s.tempDB, nil
}

func (s *MyStore) DeleteReview(uuid.UUID) *MyError {
	return nil
}

func (s *MyStore) UpdateReview(*Review) *MyError {
	return nil
}

func (s *MyStore) CreateReview(*Review) *MyError {
	return nil
}

func (s *MyStore) GetReviewByID(reviewId uuid.UUID) (*Review, *MyError) {
	for _, review := range s.tempDB {
		if review.Id == reviewId {
			return &review, nil
		}
	}

	return nil, &MyError{Error: "No review found for given id"}
}
