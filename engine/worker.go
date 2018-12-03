package engine

import (
	"crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher : error "+"fetcher url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parser(body, r.Url), nil
}
