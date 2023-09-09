# Grow-Simplee

To run the application, run the commant *go run main.go*, Make sure the main.go file is on the cwd.

Make sure to change the enviornment variables according to your local configurations.

API Documentation: https://docs.google.com/document/d/1G2XkiBG61FWn67sFQ73y7Rh2hZdF9eUtVTlJyeLvhdw/edit?usp=sharing

This applications servers 5 kind of operations

Base URL - "http://localhost:8080"

Register User (POST): "/auth/register" Body (email, name)

Login User (POST): "/auth/login" Body (email, name, password)

RateMovie (POST): "/rate" Body (movieName, rating)

ListMovies (GET): "/listmovies" 

ListMovieRatings (GET): "/listmovieratings" 