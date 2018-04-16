package main

import (
	"maven-sync/engine"
	"maven-sync/parser"
)

func main() {
	const startUrl = "http://repo.maven.apache.org/maven2/"
	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		RequestChan:     make(chan engine.Request),
		ItemChan:        make(chan engine.Item),
		ParseResultChan: make(chan engine.ParseResult),
		WorkerCount:     20,
	}

	e.Run(engine.Request{
		Url: startUrl,
		ParseFunc: func(c []byte) engine.ParseResult {
			return parser.ParsePackage(c, startUrl)
		},
	})
}
