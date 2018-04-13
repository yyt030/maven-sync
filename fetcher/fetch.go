package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"maven-sync/config"
)

func Fetch(url string) ([]byte, error) {
	<-config.RateLimiter
	log.Printf("fetch url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)

	return ioutil.ReadAll(bodyReader)
}
