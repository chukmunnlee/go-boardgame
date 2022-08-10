package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

var Commit string

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

	bga := New(*clientId)
	var games *[]Boardgame
	games, err := bga.Search(*search, *limit, *skip)

	if nil != err {
		log.Fatalf("Error: %v\n", err)
	}

	bold := color.New(color.Bold).Add(color.FgGreen).SprintfFunc()

	if "" == Commit {
		Commit = "dev"
	}

	fmt.Printf("Version: %s\n", Commit)
	for i := range *games {
		fmt.Printf("%s: %s\n", bold("Title"), (*games)[i].Name)
		fmt.Printf("%s: %s\n", bold("Description"), (*games)[i].Description)
	}
}
