package main

import (
	"github.com/google/uuid"
)

type Storage interface {
	CreateReview(*Review) error
	DeleteReview(uuid.UUID) error
	GetReviewByID(uuid.UUID) (*Review, error)
	GetReviews() ([]*Review, error)
	UpdateReview(*Review) error
}
