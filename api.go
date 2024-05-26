package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /review", makeHttpHandlerFunc(s.handleGetMovieReview))
	router.HandleFunc("GET /review/{id}", makeHttpHandlerFunc(s.handleGetMovieReviewByID))
	router.HandleFunc("POST /review", makeHttpHandlerFunc(s.handlePostMovieReview))
	router.HandleFunc("PUT /review/{id}", makeHttpHandlerFunc(s.handleUpdateMovieReview))
	router.HandleFunc("DELETE /review/{id}", makeHttpHandlerFunc(s.handleDeleteMovieReview))

	log.Print("Server started ", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)

	if err != nil {
		log.Fatal("Server couldn't be started ", err)
	}
}

func (s *APIServer) handleGetMovieReview(w http.ResponseWriter, r *http.Request) error {
	reviews, err := s.store.GetReviews()
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, reviews)
}

func (s *APIServer) handleGetMovieReviewByID(w http.ResponseWriter, r *http.Request) error {
	reviewIdStr := r.PathValue("id")
	reviewId, err := uuid.Parse(reviewIdStr)
	if err != nil {
		return fmt.Errorf("invalid id provided %v", reviewIdStr)
	}
	review, err := s.store.GetReviewByID(reviewId)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, err.Error())
	}

	return WriteJSON(w, http.StatusOK, review)
}

func (s *APIServer) handlePostMovieReview(w http.ResponseWriter, r *http.Request) error {
	createReview := &CreateReview{}

	if err := json.NewDecoder(r.Body).Decode(createReview); err != nil {
		return err
	}

	review := NewReview(createReview.MovieName, createReview.Rating, createReview.Comment)

	if err := s.store.CreateReview(review); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}

	return WriteJSON(w, http.StatusCreated, review)
}

func (s *APIServer) handleDeleteMovieReview(w http.ResponseWriter, r *http.Request) error {
	reviewIdStr := r.PathValue("id")

	reviewId, err := uuid.Parse(reviewIdStr)

	review, _ := s.store.GetReviewByID(reviewId)
	if err != nil {
		return fmt.Errorf("invalid id provided %v", reviewIdStr)
	}
	if err := s.store.DeleteReview(reviewId); err != nil {
		return WriteJSON(w, http.StatusNotFound, err.Error())
	}

	return WriteJSON(w, http.StatusOK, review)
}

func (s *APIServer) handleUpdateMovieReview(w http.ResponseWriter, r *http.Request) error {
	reviewIdStr := r.PathValue("id")
	reviewId, err := uuid.Parse(reviewIdStr)

	if err != nil {
		return fmt.Errorf("invalid id provided %v", reviewIdStr)
	}

	review, err := s.store.GetReviewByID(reviewId)

	if err != nil {
		return WriteJSON(w, http.StatusNotFound, err.Error())
	}

	updatedReview := &CreateReview{}

	if err := json.NewDecoder(r.Body).Decode(updatedReview); err != nil {
		return WriteJSON(w, http.StatusBadRequest, err.Error())
	}

	review.Comment = updatedReview.Comment
	review.MovieName = updatedReview.MovieName
	review.Rating = updatedReview.Rating

	if err := s.store.UpdateReview(review); err != nil {
		return WriteJSON(w, http.StatusNotFound, err.Error())
	}

	return WriteJSON(w, http.StatusCreated, review)
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

func makeHttpHandlerFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, MyError{Error: err.Error()})
		}
	}
}
