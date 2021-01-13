# Image Repository

## How to Use
### Install Required Dependencies
To install the necessary dependencies for the frontend run `yarn install` in the `/frontend` directory. To install necessary dependencies for the backend run `go mod download` in the `/backend` directory.
### Run the Application
To run the frontend, run `yarn start` in the `/frontend` folder. To run the backend run `go run main.go` in the `/backend` folder. To view the application, open [localhost:3000](localhost:3000) in the browser.

## Features
- upload one image
- upload multiple images
- delete images
- search for images from text (description and location)
- JWT authorization implementation
- image visibility (public or private)
- access control

## Technologies
The frontend is built with React. The backend is written in Go using the gin-gonic web framework. The application uses an SQLite database, and the backend uses GORM to query the database.
