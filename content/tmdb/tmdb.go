package tmdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"ocelot/content/types"
)

type Client struct {
	ApiKey string `json:"api_read_access_token"`
}

func (t Client) Fetch(params types.FetchParams, result any) error {
	if params.Method == "" {
		params.Method = "GET"
	}
	client := &http.Client{}
	req, err := http.NewRequest(params.Method, fmt.Sprintf("https://api.themoviedb.org/3%s?%s", params.Endpoint, strings.Join(params.Queries, "&")), nil)
	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.ApiKey))

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Status code not 200")
	}

	err = json.NewDecoder(resp.Body).Decode(result)

	return err
}

func (t Client) GetFromId(Id int) (types.Movie, error) {
	var result types.Movie

	err := t.Fetch(types.FetchParams{
		Endpoint: "/movie/11",
	}, &result)

	return result, err
}
