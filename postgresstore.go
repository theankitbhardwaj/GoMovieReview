package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db sqlx.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	db, err := sqlx.Connect("postgres", "postgres://postgres:1234@localhost:5432/GoMovieReview?sslmode=disable")
	if err != nil {
		return nil, err
	}

	db.Ping()
	return &PostgresStore{
		db: *db,
	}, nil
}

func (s *PostgresStore) init() error {
	query := `create table if not exists review (
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		movie_name varchar(50) NOT NULL,
		rating integer NOT NULL,
		comment text,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateReview(review *Review) error {
	query := `insert into review 
		(movie_name, rating, comment, created_at)
		VALUES ($1, $2, $3, $4)`

	res, err := s.db.Query(query,
		review.MovieName,
		review.Rating,
		review.Comment,
		review.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Print(res)

	return nil
}

func (s *PostgresStore) DeleteReview(id uuid.UUID) error {
	query := "delete from review where id = $1"
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostgresStore) GetReviewByID(id uuid.UUID) (*Review, error) {
	query := "select * from review where id = $1"
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanRowToReview(rows)
	}

	return nil, fmt.Errorf("review %v not found", id)
}
func (s *PostgresStore) GetReviews() ([]*Review, error) {
	query := "select * from review"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	reviews := []*Review{}

	for rows.Next() {
		review, err := scanRowToReview(rows)

		if err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (s *PostgresStore) UpdateReview(review *Review) error {
	query := `UPDATE review SET 
		movie_name = $1,
		rating = $2,
		comment = $3
		WHERE id = $4`

	_, err := s.db.Query(query, review.MovieName, review.Rating, review.Comment, review.Id)

	if err != nil {
		return err
	}

	return nil
}

func scanRowToReview(rows *sql.Rows) (*Review, error) {
	review := Review{}

	err := rows.Scan(
		&review.Id,
		&review.MovieName,
		&review.Rating,
		&review.Comment,
		&review.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &review, nil
}
