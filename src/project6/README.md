# Go Movie API

This is a simple HTTP based Movie API developed using Go. It allows users to create, fetch and get details of movies. The application also includes basic authentication and authorization using JWT tokens.

## Structure of the Project

````
src/project6
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── main.go
├── middlewares.go
├── models.go
├── routes.go
├── routes_auth.go
├── routes_movie.go
├── routes_movie_test.go
├── server.go
├── store.go
└── templates
    ├── base.html
    └── login.html

1 directory, 15 files
````

## Models

The application uses the following data model:

> Movie: A data structure representing a movie with the following fields: ID, Title, ReleaseDate, Duration, and TrailerURL.

## Endpoints

- **/login (GET)** - User login page.
- **/ (GET)** - Welcome page.
- **/api/token (POST)** - Endpoint for token generation. Users can provide their username and password, and receive a JWT token if their credentials are valid.
- **/api/movies/{id:[0-9]+} (GET)** - Fetches a specific movie by its ID. This endpoint requires authentication.
- **/api/movies/ (GET)** - Lists all movies. This endpoint requires authentication.
- **/api/movies/ (POST)** - Creates a new movie. This endpoint requires authentication.

## Setup and Run

Before running the application, please ensure you have Go installed (version 1.16+ recommended).

1. Navigate into the src/project6 directory:

````bash
cd src/project6
````

To run the server:

````go
go run main.go
````

The server will start at localhost:8080.

## Tests

This project includes some basic integration tests. To run the tests:

````bash
go test -v
````