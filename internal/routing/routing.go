package routing

import (
	"github.com/Suke2004/load-balancer/internal/core"
)

type Router interface {
	GetNextAvailableServer(backends []*core.Backend) *core.Backend
}