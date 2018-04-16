package main

import (
	"flag"
	"fmt"

	"maven-sync/engine"
	"maven-sync/parser"

	"github.com/go-redis/redis"
)

var process = flag.Int("process", 10, "Number of multiple requests to download packages")
var startUrl = flag.String("startUrl", "http://repo.maven.apache.org/maven2/", "Start download url")
var redisUrl = flag.String("redisUrl", "localhost:6379", "Redis url for save downloaded filename")

func main() {
	flag.Parse()
	//const startUrl = "http://repo.maven.apache.org/maven2/"
	//e := engine.SimpleEngine{}

	// Create redis client
	//var client *redis.Client
	client := redis.NewClient(&redis.Options{
		Addr:     *redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
		client = nil
	}

	//get := client.Get("key")
	//fmt.Println(get.String())

	e := engine.ConcurrentEngine{
		RequestChan:     make(chan engine.Request),
		ItemChan:        make(chan engine.Item),
		ParseResultChan: make(chan engine.ParseResult),
		WorkerCount:     *process,
		RedisClient:     client,
	}

	e.Run(engine.Request{
		Url: *startUrl,
		ParseFunc: func(c []byte) engine.ParseResult {
			return parser.ParsePackage(c, *startUrl)
		},
	})
}
