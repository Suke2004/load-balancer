package l4

import (
	"io"
	"net"
	"time"

	"github.com/Suke2004/load_balancer/internal/core"
	"github.com/Suke2004/load_balancer/internal/routing"
)

type TCPProxy struct {
	Pool []*core.Backend
	Router routing.Router
}

func (p *TCPProxy) Start(listenAddr string) error {
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	for {
		clineConn, err := listen.Accept()
		if err != nil {
			continue
		}

		go p.handleConnection(clineConn)
	}
}

func (p *TCPProxy) handleConnection(clineConn net.Conn) {
	start := time.Now()
	backend := p.Router.GetNextAvailableServer(p.Pool)

	if backend == nil {
		clineConn.Close()
		return
	}

	backend.Mutex.Lock()
	backend.ActiveConns++
	backend.Mutex.Unlock()

	serverConn, err := net.Dial("tcp", backend.Address)
	if err != nil{
		backend.Mutex.Lock()
		backend.ActiveConns--
		backend.Mutex.Unlock()
		clineConn.Close()
		return
	}

	go io.Copy(serverConn, clineConn)
	io.Copy(clineConn, serverConn)

	clineConn.Close()
	serverConn.Close()

	backend.Mutex.Lock()
	backend.ActiveConns--
	lat := time.Since(start)
	backend.Latency = (backend.Latency+lat)/2
	backend.Mutex.Unlock()


}