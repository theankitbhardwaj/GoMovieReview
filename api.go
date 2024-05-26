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

	router.HandleFunc("GET /review", s.handleGetMovieReview)
	router.HandleFunc("GET /review/{id}", s.handleGetMovieReviewByID)
	router.HandleFunc("POST /review", s.handlePostMovieReview)
	router.HandleFunc("PUT /review/{id}", s.handlePostMovieReview)
	router.HandleFunc("DELETE /review/{id}", s.handleDeleteMovieReview)

	log.Print("Server started ", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)

	if err != nil {
		log.Fatal("Server couldn't be started ", err)
	}
}

func (s *APIServer) handleGetMovieReview(w http.ResponseWriter, r *http.Request) {
	reviews, err := s.store.GetReviews()
	if err != nil {
		WriteJSON(w, 500, err)
	}
	WriteJSON(w, 200, reviews)
}

func (s *APIServer) handleGetMovieReviewByID(w http.ResponseWriter, r *http.Request) {
	reviewIdStr := r.PathValue("id")
	reviewId, err := uuid.Parse(reviewIdStr)
	if err != nil {
		WriteJSON(w, 400, &MyError{Error: fmt.Sprintf("Invalid id provided %v", reviewIdStr)})
		return
	}
	review, error := s.store.GetReviewByID(reviewId)
	if error != nil {
		fmt.Print(error)
		WriteJSON(w, 404, error)
		return
	}

	WriteJSON(w, 200, review)
}

func (s *APIServer) handlePostMovieReview(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, 200, "New Movie review posted!")
}

func (s *APIServer) handleDeleteMovieReview(w http.ResponseWriter, r *http.Request) {
	reviewId := r.PathValue("id")
	WriteJSON(w, 200, fmt.Sprintf("Movie Review deleted for id: %v", reviewId))
}

func (s *APIServer) handleUpdateMovieReview(w http.ResponseWriter, r *http.Request) {
	reviewId := r.PathValue("id")
	WriteJSON(w, 200, fmt.Sprintf("Movie Review updated for id: %v", reviewId))
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}
