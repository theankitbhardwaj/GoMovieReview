package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type Storage interface {
	CreateReview(*Review) error
	DeleteReview(uuid.UUID) error
	GetReviewByID(uuid.UUID) (*Review, error)
	GetReviews() (*[]Review, error)
	UpdateReview(*Review) error
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

func (s *MyStore) GetReviews() (*[]Review, error) {
	return &s.tempDB, nil
}

func (s *MyStore) DeleteReview(reviewId uuid.UUID) error {
	indexToDelete := -1
	for i, review := range s.tempDB {
		if review.Id == reviewId {
			indexToDelete = i
		}
	}
	if indexToDelete != -1 {
		s.tempDB = append(s.tempDB[:indexToDelete], s.tempDB[indexToDelete+1:]...)
		return nil
	}

	return fmt.Errorf("no review found for id %v", reviewId)
}

func (s *MyStore) UpdateReview(updatedReview *Review) error {
	for i, review := range s.tempDB {
		if review.Id == updatedReview.Id {
			s.tempDB[i] = *updatedReview
			return nil
		}
	}

	return fmt.Errorf("no review found for id %v", updatedReview.Id)
}

func (s *MyStore) CreateReview(review *Review) error {
	s.tempDB = append(s.tempDB, *review)

	return nil
}

func (s *MyStore) GetReviewByID(reviewId uuid.UUID) (*Review, error) {
	for _, review := range s.tempDB {
		if review.Id == reviewId {
			return &review, nil
		}
	}

	return nil, fmt.Errorf("no review found for given id %v", reviewId)
}
