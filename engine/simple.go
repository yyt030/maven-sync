package engine

import (
	"fmt"
	"time"

	"maven-sync/downloader"
)

type SimpleEngine struct{}

var visitedUrls = make(map[string]bool)

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for {
		for len(requests) > 0 {
			r := requests[0]
			requests = requests[1:]

			parseResult, err := Work(r)
			if err != nil {
				continue
			}

			requests = append(requests, parseResult.Request...)
			for _, item := range parseResult.Items {
				if isDuplicate(item.Url) {
					continue
				}

				downloader.Download(item.Name, item.Url)
			}
		}
		fmt.Println("sleeping... you can shutdown me when all package download")
		time.Sleep(time.Minute)
	}
}

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
