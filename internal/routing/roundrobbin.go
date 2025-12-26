package routing

import (
	"fmt"
	"sync"

	"github.com/Suke2004/load-balancer/internal/core"
)

type RoundRobinRouter struct {
	Current int
	Mu sync.Mutex
}

func NewRoundRobinRouter() *RoundRobinRouter {
	return &RoundRobinRouter{
		Current: 0,
	}
}

func (rr *RoundRobinRouter) GetNextAvailableServer(
	backends []*core.Backend,
) *core.Backend {
	rr.Mu.Lock()
	defer rr.Mu.Unlock()

	n := len(backends)
	if n == 0 {
		fmt.Println("No Servers available in the pool")
		return nil
	}

	for i:=0;i<n;i++ {
		idx := (rr.Current + i) % n
		backend := backends[idx]

		backend.Mutex.Lock()
		isAlive := backend.Alive
		backend.Mutex.Unlock()

		if isAlive {
			rr.Current = (idx + 1) % n
			return backend
		}
	}
	return nil
}

func (rr *RoundRobinRouter) Name() string {
	return "RoundRobin"
}