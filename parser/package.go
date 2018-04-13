package parser

import (
	"regexp"

	"maven-sync/config"
	"maven-sync/engine"
)

var pkgDirRe = regexp.MustCompile(`<a href="([a-zA-Z0-9.-]+/)"[^>]*>([a-zA-Z0-9.-]+/)</a>`)
var pkgFileRe = regexp.MustCompile(`<a href="([a-zA-Z0-9.-]+)"[^>]*>([a-zA-Z0-9.-]+)</a>`)

func ParsePackage(contents []byte, url string) engine.ParseResult {
	dirMatches := pkgDirRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	num := 0
	for _, m := range dirMatches {
		u := string(m[1])
		if u == "../" {
			continue
		}

		newUrl := url + string(m[1])
		result.Request = append(result.Request, engine.Request{
			Url: newUrl,
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParsePackage(c, newUrl)
			},
		})

		num++
		if num >= config.LimitNum {
			break
		}
	}

	pkgMatches := pkgFileRe.FindAllSubmatch(contents, -1)
	for _, m := range pkgMatches {
		name := string(m[1])
		loc := regexp.MustCompile(`.md5|.sha1|.asc`).FindIndex(m[1])
		if loc != nil {
			continue
		}

		newUrl := url + string(m[1])
		result.Items = append(result.Items, engine.Item{
			Name: name,
			Url:  newUrl,
		})
	}

	return result
}
