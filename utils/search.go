package utils

import (
	"backend-assignment/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func SearchMovie(query string) string {
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%s", query)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJmZTVjOTBkNDE4MzZlZGRlYWY5ZTE5OTMwMTE1NmE5OSIsInN1YiI6IjY0YzM2OGRhZDg2MWFmMDBmZmY5NTJhMCIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.xXgIgp5b98nd6182E5i3o6L5-NsHzbmUEuItl8LyQLI")

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var movieData responses.Movie
	err = json.Unmarshal([]byte(body), &movieData)
	if err != nil {
		log.Err(err).Msg("Error unmarshaling JSON:")
		return ""
	}
	if movieData.TotalResults == 0 {
		return ""
	}
	return movieData.Results[0].Title
}
