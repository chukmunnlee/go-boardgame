package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	// Parse the parameter
	limit := flag.Int("limit", 20, "Limits the number of results returned")
	skip := flag.Int("skip", 0, "Skips the number of results provided")
	clientId := flag.String("client-id", os.Getenv("BGA_CLIENT_ID"), "Boardgame Atlas key")
	search := flag.String("search", "", "Boardgame name")

	flag.Parse()

	if "" == *clientId {
		log.Panicln("client-id is not set")
	}
	if "" == *search {
		log.Fatalln("search parameter is not set")
	}

	fmt.Printf("search=%s, limit=%d, offset=%d\n", *search, *limit, *skip)

	bga := BoardgameAtlas{ClientId: *clientId}
	bga.Search(*search, *limit, *skip)

}
