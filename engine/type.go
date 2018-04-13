package engine

type Request struct {
	Url       string
	ParseFunc ParseFunc
}

type ParseResult struct {
	Request []Request
	Items   []Item
}

type Item struct {
	Name string
	Url  string
}

type ParseFunc func(contents []byte) ParseResult
