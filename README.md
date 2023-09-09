# Movie_rater_application

This application can register a new user, log in to an existing one, rate a new movie, and get the list of movies with their current ratings.

The default password is set to "qwerty"

To run the application, run the command *go run main.go*, Make sure the main.go file is on the cwd.

Make sure to change the enviornment variables according to your local configurations.

API Documentation: https://docs.google.com/document/d/1G2XkiBG61FWn67sFQ73y7Rh2hZdF9eUtVTlJyeLvhdw/edit?usp=sharing

This applications servers 5 kind of operations

Base URL - "http://localhost:8080"

Register User (POST): "/auth/register" Body (email, name)

Login User (POST): "/auth/login" Body (email, name, password)

RateMovie (POST): "/rate" Body (movieName, rating)

ListMovies (GET): "/listmovies" 

ListMovieRatings (GET): "/listmovieratings" 
