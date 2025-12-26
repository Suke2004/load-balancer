package api

import (
	"net/http"
	"encoding/json"

	"github.com/sbirmecha99/load_balancer/internal/core"
)


func MetricsHandler(pool *core.Backend) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode(pool)
	}
}