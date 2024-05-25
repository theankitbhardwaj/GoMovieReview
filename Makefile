build:
	@go build -o bin/GoMovieReview

run: build
	@./bin/GoMovieReview

test:
	@go test -v ./ ...