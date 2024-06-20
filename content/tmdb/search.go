package tmdb

import (
	"fmt"
	"strings"

	"ocelot/content/types"
)

func (t Client) SearchMovies(param types.SearchParam) (types.MovieSearchResponse, error) {

	var response types.MovieSearchResponse
	err := t.Fetch(types.FetchParams{
		Endpoint: "/search/movie",
		Queries:  []string{"query=" + strings.ReplaceAll(param.Query, " ", "%20"), fmt.Sprintf("first_air_date_year=%d", param.Year)},
	}, &response)
	if err != nil {
		return types.MovieSearchResponse{}, err
	}

	return response, nil
}

func (t Client) SearchShows(param types.SearchParam) (types.ShowSearchResponse, error) {

	var response types.ShowSearchResponse
	err := t.Fetch(types.FetchParams{
		Endpoint: "/search/tv",
		Queries:  []string{"query=" + strings.ReplaceAll(param.Query, " ", "%20"), fmt.Sprintf("first_air_date_year=%d", param.Year)},
	}, &response)

	if err != nil {
		return types.ShowSearchResponse{}, err
	}

	return response, nil
}
