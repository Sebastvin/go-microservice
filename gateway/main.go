package main

import (
	"context"
	"log"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
	common "github.com/sebastvin/commons"
	"github.com/sebastvin/commons/discovery"
	"github.com/sebastvin/commons/discovery/consul"
	"github.com/sebastvin/omsv-gateway/gateway"
)

var (
	serviceName = "gateway"
	httpAddr    = common.EnvString("HTTP_ADDR", ":3000")
	consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
	jaegerAddr  = common.EnvString("JAEGER_ADDR", "localhost:4318")
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	err := common.SetGlobalTracer(context.TODO(), serviceName, jaegerAddr)
	if err != nil {
		log.Fatal("failed to set global tracer")
	}

	registry, err := consul.NewRegistry(consulAddr, serviceName)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, httpAddr); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Fatal("failed to health check")
			}
			time.Sleep(time.Second * 1)
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)

	mux := http.NewServeMux()

	ordersGateway := gateway.NewGRPCGateway(registry)

	handler := NewHandler(ordersGateway)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, corsMiddleware(mux)); err != nil {
		log.Fatal("Failed to start http server")
	}
}
