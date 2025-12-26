package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/Suke2004/load_balancer/internal/core"
	"github.com/Suke2004/load_balancer/internal/api"
	"github.com/Suke2004/load_balancer/internal/routing"
)
func startDummyBackend(port string) {
	go func() {
		mux := http.NewServeMux()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello Backend %s", port)
		})

		log.Printf("[DUMMY] starting backend on :%s\n", port)

		if err := http.ListenAndServe(":"+port, mux); err != nil {
			log.Printf("[DUMMY %s] crashed: %v", port, err)
		}
	}()
}

func main(){
	startDummyBackend("9001")
	startDummyBackend("9002")
	startDummyBackend("9003")
	select {}
}