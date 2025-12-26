package routing

import (
	"sync"

	"github.com/Suke2004/load-balancer/internal/core"
)

type LeastConnectionsRouter struct {
	Mu sync.Mutex
}

func NewLeastConnectionsRouter() *LeastConnectionsRouter {
	return &LeastConnectionsRouter{}
}

func (lc *LeastConnectionsRouter) GetNextAvailableServer(
	backend []*core.Backend,
) *core.Backend {
	lc.Mu.Lock()
	defer lc.Mu.Unlock()

	n:=len(backend)

	if n==0{
		return nil
	}

	var selectedBackend *core.Backend
	minConns := int(^uint(0) >> 1) // Max int

	for _, backend := range backend {
		backend.Mutex.Lock()
		isAlive := backend.Alive
		ActiveConnections := backend.ActiveConns
		backend.Mutex.Unlock()

		if isAlive && ActiveConnections < int64(minConns) {
			minConns = int(ActiveConnections)
			selectedBackend = backend
		}
	}
	if selectedBackend != nil {
		selectedBackend.Mutex.Lock()
		selectedBackend.ActiveConns++
		selectedBackend.Mutex.Unlock()
	}
	return selectedBackend
}

func (lc *LeastConnectionsRouter) Name() string {
	return "Least Connections"
}