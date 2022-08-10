package main

import (
	"encoding/json"
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
	output := flag.String("output", "text", "Output format")

	flag.Parse()

	if "" == *clientId {
		log.Panicln("client-id is not set")
	}
	if "" == *search {
		log.Fatalln("search parameter is not set")
	}

	bga := New(*clientId)
	var games *[]Boardgame
	games, err := bga.Search(*search, *limit, *skip)

	if nil != err {
		log.Fatalf("Error: %v\n", err)
	}

	if "" == Commit {
		Commit = "dev"
	}

	switch *output {
	case "json":
		printJson(games)

	case "text":
		printText(games)

	default:
		printText(games)
	}
}

func printJson(games *[]Boardgame) {
	jsonStr, err := json.Marshal(games)
	if nil != err {
		log.Fatalf("Cannot marshal games to Json: %v", err)
	}
	fmt.Println(string(jsonStr))
}

func printText(games *[]Boardgame) {
	bold := color.New(color.Bold).Add(color.FgGreen).SprintfFunc()
	fmt.Printf("Version: %s\n", Commit)
	for i := range *games {
		fmt.Printf("%s: %s\n", bold("Title"), (*games)[i].Name)
		fmt.Printf("%s: %s\n\n", bold("Description"), (*games)[i].Description)
	}
}
