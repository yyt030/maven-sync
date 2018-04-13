package main

import (
	"maven-sync/engine"
	"maven-sync/parser"
)

func main() {
	const startUrl = "http://repo.maven.apache.org/maven2/"
	e := engine.SimpleEngine{}

	e.Run(engine.Request{
		Url: startUrl,
		ParseFunc: func(c []byte) engine.ParseResult {
			return parser.ParsePackage(c, startUrl)
		},
	})
}