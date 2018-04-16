package engine

import (
	"fmt"
	"time"

	"maven-sync/fetcher"

	"github.com/go-redis/redis"
)

type ConcurrentEngine struct {
	RequestChan     chan Request
	ItemChan        chan Item
	ParseResultChan chan ParseResult
	WorkerCount     int
	RedisClient     *redis.Client
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	// Create worker
	// read chan from RequestChan and do work
	// read chan from ItemChan and do work
	c.CreateWorker()
	// Get one request
	for _, r := range seeds {
		go func() { c.RequestChan <- r }()
	}

	fmt.Println(">>> sleeping ...")

	for {
		p := <-c.ParseResultChan
		go func() {
			for _, rr := range p.Request {
				c.RequestChan <- rr
			}
		}()
		go func() {
			for _, ii := range p.Items {
				c.ItemChan <- ii
			}
		}()
	}
}

func (c *ConcurrentEngine) CreateWorker() {
	for i := 0; i < c.WorkerCount; i++ {
		go func(i int) {
			for {
				select {
				case item := <-c.ItemChan:
					if c.RedisClient != nil {
						val, err := c.RedisClient.Get("maven:" + item.Url).Result()
						if err == nil && string(val) == item.Name {
							continue
						}
					}

					if err := fetcher.Download(item.Name, item.Url); err == nil && c.RedisClient != nil {
						c.RedisClient.Set("maven:"+item.Url, item.Name, 0)
					}
				case r := <-c.RequestChan:
					parseResult, err := Work(r)
					if err != nil {
						continue
					}
					go func() { c.ParseResultChan <- parseResult }()
				default:
					time.Sleep(time.Second)
				}

			}
		}(i)
	}
}
