package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cyuliu/golang/httpserver/metrics"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	metrics.Register()
	mux := http.NewServeMux()
	mux.HandleFunc("/delay", delayHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.Handle("/metrics", promhttp.Handler())
	srv := http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Fatalf("listen: %s\n", err)
		}
	}()
	glog.Info("Server Started")
	<-done
	glog.Info("Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Fatalf("Server Shutdown Failed:%+v", err)
	}
	glog.Info("Server Exited Properly")
}

func delayHandler(writer http.ResponseWriter, request *http.Request) {
	glog.V(4).Info("Entering delay handler")
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	query := request.URL.Query()
	for key, value := range query {
		io.WriteString(writer, fmt.Sprintf("%s=%s\n", key, value))
	}
	io.WriteString(writer, "===================Details of the http request header:============\n")
	for k, v := range request.Header {
		io.WriteString(writer, fmt.Sprintf("%s=%s\n", k, v))
	}

	// 随机等待
	delay := randInt(10, 2000)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	glog.V(4).Infof("Respond in %d ms", delay)
	glog.V(4).Info("Exiting delay handler")
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "ok\n")
}
