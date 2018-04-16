package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"testing"
)

func TestDownload(t *testing.T) {
	//u := "http://repo.maven.apache.org/maven2/acegisecurity/acegi-security/0.6.1/acegi-security-0.6.1.jar.sha1"
	u := "http://repo.maven.apache.org/maven2"
	parse, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("path:%s\n", parse.Path)

	dirName := "pkg" + path.Dir(parse.Path)

	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Printf("create path failed: %v, error:%v\n", dirName, err)
	}

	file, err := os.Create(path.Join(dirName, "index.html"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(file, resp.Body)

}
