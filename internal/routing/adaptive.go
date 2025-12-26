package routing

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Suke2004/load-balancer/internal/core"
)

type AdaptiveRouter struct {
	Pool *core.ServerPool
	rr *RoundRobinRouter
	lc *LeastConnectionsRouter
	rn *RandomConnectionsRouter
	CurrentAlgo string
	Reason string
	LastPicked string
}
//decsion to switch algorithm can be based on various factors like server load, response time, etc.
type Decision struct {
	Time time.Time `json:"time"`
	Algo string `json:"algo"`
	Reason string `json:"reason"`
	Backend string `json:"backend"`
}
var (
	DecisionLog []Decision
	DecisionMu sync.Mutex
)

func NewAdaptiveRouter(pool *core.ServerPool) *AdaptiveRouter {
	return &AdaptiveRouter{
		Pool: pool,
		rr: NewRoundRobinRouter(),
		lc: NewLeastConnectionsRouter(),
		rn: NewRandomConnectionsRouter(),
		CurrentAlgo: "RoundRobin",
		Reason: "normal_conditions",
	}
}

func (ar *AdaptiveRouter) Pick() *core.Backend {
	backends := ar.Pool.GetServers()
	if len(backends) == 0 {
		log.Println("[adaptive] no backends in pool")
		return nil
	}
	var totalConns int64
	var totalLatency int64
	var totalErrors int64
	var maxConns int64
	aliveCount:=0

	for _, backend := range backends {
		backend.Mutex.Lock()

		if backend.Alive {
			aliveCount++
			totalConns += backend.ActiveConns
			totalLatency += int64(backend.Latency)
			totalErrors += backend.ErrorCount

			if backend.ActiveConns > maxConns {
				maxConns = backend.ActiveConns
			}
		}
		backend.Mutex.Unlock()
	}

	if aliveCount == 0 {
		log.Println("[adaptive] no alive backends")
	}

	avgConns := totalConns/int64(aliveCount)
	avgLatency := totalLatency/int64(aliveCount)
	errorRate := float64(totalErrors)/float64(totalConns+1)
	avgLatencyMS := avgLatency/int64(time.Millisecond)

	log.Printf("Adaptive algo=%s picked=%s avgConns=%d maxConns=%d avgLatencyMS=%d errorRate=%.2f", ar.CurrentAlgo, ar.LastPicked, avgConns, maxConns, avgLatencyMS, errorRate)

	var selectedBackend *core.Backend

	//decision logic to switch algorithms
	if errorRate>0.3{
		ar.CurrentAlgo="random"
		ar.Reason = fmt.Sprintf("high_error_rate{%.2f}", errorRate)
		selectedBackend = ar.rn.GetNextAvailableServer(backends)
	} else if maxConns > 3 {
		ar.CurrentAlgo = "least_connections"
		ar.Reason = "high concurrency"
		selectedBackend = ar.lc.GetNextAvailableServer(backends)
	} else {
		ar.CurrentAlgo = "round_robin"
		ar.Reason = "normal_conditon"
		selectedBackend =  ar.rr.GetNextAvailableServer(backends)
	}

	if selectedBackend != nil {
		ar.LastPicked = selectedBackend.Address
		DecisionMu.Lock()
		DecisionLog = append(DecisionLog, Decision{
			Time: time.Now(),
			Algo: ar.CurrentAlgo,
			Reason : ar.Reason,
			Backend: selectedBackend.Address,
		})
		DecisionMu.Unlock()
	}
	return selectedBackend
}

func (ar *AdaptiveRouter) GetNextAvailableServer(_ []*core.Backend) *core.Backend {
	return ar.Pick()
}

func (ar *AdaptiveRouter) Name() string {
	return "Adaptive"
}

func (ar *AdaptiveRouter) CurrentAlgorithm() string {
	return ar.CurrentAlgo
}

func (ar *AdaptiveRouter) Reasons() string {
	return ar.Reason
}

func (ar *AdaptiveRouter) LastPickedBackend() string {
	return ar.LastPicked
}

func (ar *AdaptiveRouter) GetDecisionLog() []Decision {
	return DecisionLog
}


