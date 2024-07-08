package main

import (
	"fmt"

	"github.com/google/uuid"
)

type TempStore struct {
	tempDB []*Review
}

func NewTempStore() (*TempStore, error) {
	tempDB := make([]*Review, 0)
	return &TempStore{tempDB: tempDB}, nil
}

func (s *TempStore) init() error {
	for i := 0; i < 10; i++ {
		review := Review{
			Id:        uuid.New(),
			MovieName: fmt.Sprintf("Movie %v", i),
			Rating:    3,
			Comment:   fmt.Sprintf("Description of review for movie %v", i),
		}
		s.tempDB = append(s.tempDB, &review)
	}

	return nil
}

func (s *TempStore) GetReviews() ([]*Review, error) {
	return s.tempDB, nil
}

func (s *TempStore) DeleteReview(reviewId uuid.UUID) error {
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

func (s *TempStore) UpdateReview(updatedReview *Review) error {
	for i, review := range s.tempDB {
		if review.Id == updatedReview.Id {
			s.tempDB[i] = updatedReview
			return nil
		}
	}

	return fmt.Errorf("no review found for id %v", updatedReview.Id)
}

func (s *TempStore) CreateReview(review *Review) error {
	s.tempDB = append(s.tempDB, review)

	return nil
}

func (s *TempStore) GetReviewByID(reviewId uuid.UUID) (*Review, error) {
	for _, review := range s.tempDB {
		if review.Id == reviewId {
			return review, nil
		}
	}

	return nil, fmt.Errorf("no review found for given id %v", reviewId)
}
