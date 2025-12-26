package health

import (
	"net"
	"time"

	"github.com/Suke2004/load-balancer/internal/core"
)

type HealthChecker struct {
	Pool *core.ServerPool
	Interval time.Duration
	Timeout time.Duration
}

func (c *HealthChecker) Start() {
	ticker := time.NewTicker(c.Timeout)

	go func() {
		for range ticker.C {
			for _, backend := range c.Pool.GetServers() {
				go c.CheckBackend(backend)
			}
		}
	}()
}

func (c *HealthChecker) CheckBackend(b *core.Backend) {
	start := time.Now()

	conn, err := net.DialTimeout("tcp", b.Address, c.Timeout)
	b.Mutex.Lock() // Locking the backend for safe concurrent access
	defer b.Mutex.Unlock()

	if err != nil {
		b.Alive = false
		b.ErrorCount++
		return
	}

	_ = conn.Close()
	b.Alive = true
	b.Latency = time.Since(start)
}

