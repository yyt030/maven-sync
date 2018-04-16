package engine

import (
	"maven-sync/fetcher"
)

type SimpleEngine struct{}

var visitedUrls = make(map[string]bool)

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// Pop first request
		r := requests[0]
		requests = requests[1:]

		// Fetch data from url and parser
		parseResult, err := Work(r)
		if err != nil {
			continue
		}

		// Add request
		requests = append(requests, parseResult.Request...)

		// Download file item
		for _, item := range parseResult.Items {
			if isDuplicate(item.Url) {
				continue
			}
			fetcher.Download(item.Name, item.Url)
		}
	}
}

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
