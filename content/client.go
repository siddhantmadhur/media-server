package content

import "ocelot/content/types"

type Client interface {
	Fetch(types.FetchParams, any) error
	SearchMovies(types.SearchParam) (types.MovieSearchResponse, error)
	SearchShows(types.SearchParam) (types.ShowSearchResponse, error)
	Authenticate() bool
}

func NewClient(client Client) (Client, error) {
	return client, nil
}
