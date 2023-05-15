package ratelimit

import (
	"sync"
	"time"
)

const (
	DEFAULT_QUANTUM = 1
)

type Bucket struct {
	// token的生产速率，每rate时间生产quantum个令牌
	rate time.Duration

	quantum int64

	// 桶容量
	cap int64

	// 桶中剩余令牌数
	tokens int64

	// 上次放令牌的时间
	latestTime time.Time

	mu sync.Mutex
}

// Allow 用来判断当前是否有令牌可取
func (b *Bucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.adjustAvailableToken()
	if b.tokens > 0 {
		b.tokens--
		return true
	} else {
		return false
	}
}

// adjustAvailableToken 用来调整token数量
func (b *Bucket) adjustAvailableToken() {
	now := time.Now()
	nums := int64(now.Sub(b.latestTime)/b.rate) * b.quantum
	if nums == 0 {
		return
	}

	b.latestTime = now
	b.tokens += nums
	if b.tokens > b.cap {
		b.tokens = b.cap
	}
	b.latestTime = time.Now()
}

func NewBucket(rate time.Duration, cap int64) *Bucket {
	return NewBucketWithQuantum(rate, cap, DEFAULT_QUANTUM)
}

func NewBucketWithQuantum(rate time.Duration, cap int64, quantum int64) *Bucket {
	return &Bucket{
		rate:       rate,
		quantum:    quantum,
		cap:        cap,
		tokens:     cap,
		latestTime: time.Now(),
	}
}
