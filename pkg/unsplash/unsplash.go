package unsplash

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const baseURL = "https://api.unsplash.com"

func GetPhotos(page int, perPage int) ([]Photo, error) {
	url := fmt.Sprintf("%s/photos?page=%s&per_page=%s&client_id=%s", baseURL, strconv.Itoa(page), strconv.Itoa(perPage), os.Getenv("UNSPLASH_ACCESS_KEY"))
	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(body))
	}

	var photos []Photo
	err = json.NewDecoder(res.Body).Decode(&photos)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func SearchPhotos(query string, page int, perPage int) (SearchResponse, error) {
	url := fmt.Sprintf("%s/search/photos?query=%s&page=%s&per_page=%s&client_id=%s", baseURL, query, strconv.Itoa(page), strconv.Itoa(perPage), os.Getenv("UNSPLASH_ACCESS_KEY"))
	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return SearchResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return SearchResponse{}, errors.New(string(body))
	}

	var searchResponse SearchResponse
	err = json.NewDecoder(res.Body).Decode(&searchResponse)
	if err != nil {
		return SearchResponse{}, err
	}

	return searchResponse, nil
}
