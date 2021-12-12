package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/cyuliu/golang/httpserver/metrics"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	metrics.Register()
	mux := http.NewServeMux()
	mux.HandleFunc("/delay ", delayHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.Handle("/metrics", promhttp.Handler())
}

func delayHandler(writer http.ResponseWriter, request *http.Request) {
	glog.V(4).Info("Entering delay handler")
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()

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
