package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// IPRateLimit ...
type IPRateLimit struct {
	IP        string
	RateCount int
	RateSec   time.Duration
}

// BasicKey 基礎的key
func (ip *IPRateLimit) BasicKey() string {
	return fmt.Sprintf("ratelimit:%s", ip.IP)
}

// GenerateKey 產生一組 redis 的 key
func (ip *IPRateLimit) GenerateKey() string {
	return fmt.Sprintf(ip.BasicKey() + ":" + uuid.New().String())
}

// SearchKey 搜尋 redis 當前 ip 的 key
func (ip *IPRateLimit) SearchKey() string {
	return fmt.Sprintf(ip.BasicKey() + ":*")
}
