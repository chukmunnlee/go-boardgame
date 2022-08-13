package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BOARDGAME_ATLAS_SEARCH = "https://api.boardgameatlas.com/api/search"

type Payload struct {
	Games []Boardgame `json:"games"`
	Count int32       `json:"count"`
}

type Boardgame struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url"`
	ImageUrl    string `json:"image_url"`
}

func New(clientId string) *BoardgameAtlas {
	bga := new(BoardgameAtlas)
	bga.clientId = clientId
	return bga
}

type BoardgameAtlas struct {
	clientId string
}

func (bga *BoardgameAtlas) Search(search string, limit int, skip int) (*[]Boardgame, error) {

	req, err := http.NewRequest(http.MethodGet, BOARDGAME_ATLAS_SEARCH, nil)
	if nil != err {
		return nil, fmt.Errorf("cannot create URL: %v", err)
	}

	headers := req.Header
	headers.Add("Accept", "application/json")

	query := req.URL.Query()
	query.Set("name", search)
	query.Set("client_id", bga.clientId)
	query.Set("limit", fmt.Sprintf("%d", limit))
	query.Set("skip", fmt.Sprintf("%d", skip))
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		return nil, fmt.Errorf("request error: %v", err)
	}

	payload := Payload{}
	err = json.NewDecoder(resp.Body).Decode(&payload)
	if nil != err {
		return nil, fmt.Errorf("cannot unmarshall payload: %v", err)
	}

	return &payload.Games, nil
}
