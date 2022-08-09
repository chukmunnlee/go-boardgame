package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BOARDGAME_ATLAS_SEARCH = "https://api.boardgameatlas.com/api/search"

type Boardgame struct {
	Name        string
	Price       float32
	Description string
	Url         string
	ImageUrl    string
}

type BoardgameAtlas struct {
	ClientId string
}

func (bga *BoardgameAtlas) Search(search string, limit int, skip int) {

	req, err := http.NewRequest(http.MethodGet, BOARDGAME_ATLAS_SEARCH, nil)
	if nil != err {
		log.Fatalf("Cannot create URL: %v\n", err)
	}

	query := req.URL.Query()
	query.Set("name", search)
	query.Set("client_id", bga.ClientId)
	query.Set("limit", fmt.Sprintf("%d", limit))
	query.Set("skip", fmt.Sprintf("%d", skip))
	req.URL.RawQuery = query.Encode()

	fmt.Printf("url: %s\n", req.URL.String())

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		log.Panicf("Request error: %v\n", err)
	}

	payload, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Panicf("Cannot read payload: %v\n", err)
	}

	fmt.Printf("payload: %s\n", string(payload))
}
