package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestPackage(t *testing.T) {
	file, err := os.Open("package_test_data.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	contents, _ := ioutil.ReadAll(file)
	r := ParsePackage(contents, "http://repo.maven.apache.org/maven2/commons-daemon/commons-daemon/")

	for _, i := range r.Request {
		fmt.Printf("%+v\n", i.Url)
	}

}
