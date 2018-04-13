package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"maven-sync/config"
)

func Download(filename string, u string) error {
	<-config.RateLimiter
	log.Printf("fetch url: %s [%s]", u, filename)
	parse, err := url.Parse(u)
	if err != nil {
		panic(err)
	}

	dirName := "pkg" + path.Dir(parse.Path)
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Printf("create path %v failed: %v\n", dirName, err)
	}

	file, err := os.Create(path.Join(dirName, filename))
	if err != nil {
		log.Printf("create file error: %v, filename: %s\n", err, filename)
		return err
	}
	defer file.Close()

	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil

}
