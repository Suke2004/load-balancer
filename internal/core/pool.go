package core

import (
	"sync"
)
type ServerPool struct {
	server []*Backend
	Mu sync.RWMutex
}

func NewServerPool() *ServerPool {
	return &ServerPool{}
}

func (p *ServerPool) AddServer(b *Backend) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.server = append(p.server, b)
}

func (p *ServerPool) GetServers() []*Backend{
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.server
}
