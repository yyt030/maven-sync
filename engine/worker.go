package engine

import (
	"log"

	"maven-sync/fetcher"
)

func Work(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher: error fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParseFunc(body), nil
}
