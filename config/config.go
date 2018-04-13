package config

import (
	"time"
)

const (
	Qps      = 20
	LimitNum = 3
)

var RateLimiter = time.Tick(time.Second / Qps)
