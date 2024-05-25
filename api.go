package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /review", handleGetMovieReview)
	router.HandleFunc("GET /review/{id}", handleGetMovieReviewByID)
	router.HandleFunc("POST /review", handlePostMovieReview)
	router.HandleFunc("PUT /review/{id}", handleUpdateMovieReview)
	router.HandleFunc("DELETE /review/{id}", handleDeleteMovieReview)

	log.Print("Server started ", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)

	if err != nil {
		log.Fatal("Server couldn't be started ", err)
	}
}

func handleGetMovieReview(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, 200, "All Movie Reviews")
}

func handleGetMovieReviewByID(w http.ResponseWriter, r *http.Request) {
	reviewId := r.PathValue("id")
	WriteJSON(w, 200, fmt.Sprintf("Movie Review for id: %v", reviewId))
}

func handlePostMovieReview(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, 200, "New Movie review posted!")
}

func handleDeleteMovieReview(w http.ResponseWriter, r *http.Request) {
	reviewId := r.PathValue("id")
	WriteJSON(w, 200, fmt.Sprintf("Movie Review deleted for id: %v", reviewId))
}

func handleUpdateMovieReview(w http.ResponseWriter, r *http.Request) {
	reviewId := r.PathValue("id")
	WriteJSON(w, 200, fmt.Sprintf("Movie Review updated for id: %v", reviewId))
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}
