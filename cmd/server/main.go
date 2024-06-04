package main

import (
	"IntraCat/cmd/server/env"
	"IntraCat/cmd/server/handers"
	"flag"
	"net/http"
	"time"
)


func main() {
    addr := flag.String("addr", ":3000", "Port to serve from")
    envVars := env.NewEnv()

    router := http.NewServeMux()
    router.HandleFunc("GET /", handers.IndexPageGet(&envVars))

    server := http.Server {
        Addr: *addr,
        Handler: router,
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout: 1 * time.Minute,
    }

    if err := server.ListenAndServe(); err != nil {
        envVars.Logger.Error(err.Error())
    }
}
