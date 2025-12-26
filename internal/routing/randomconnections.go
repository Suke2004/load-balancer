package routing

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/Suke2004/load_balancer/internal/core"
)

type RandomConnectionsRouter struct {
	Mu sync.Mutex
	rnd *rand.Rand
}

func NewRandomConnectionsRouter() *RandomConnectionsRouter {
	return &RandomConnectionsRouter{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (rc *RandomConnectionsRouter) GetNextAvailableServer(
	backends []*core.Backend,
) *core.Backend {
	rc.Mu.Lock()
	defer rc.Mu.Unlock()

	n := len(backends)
	if n == 0 {
		fmt.Println("No Servers available in the pool")
		return nil
	}
	for i:=0;i<n;i++ {
		idx := rc.rnd.Intn(n)
		backend := backends[idx]

		backend.Mutex.Lock()
		isAlive := backend.Alive
		backend.Mutex.Unlock()

		if isAlive {
			return backend
		}
	}
	return nil
}

func (rc *RandomConnectionsRouter) Name() string {
	return "Random Connections"
}