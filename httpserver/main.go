package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", requestHandle)
	mux.HandleFunc("/header", requestHandle)
	mux.HandleFunc("/env", requestHandle)
	// 设置环境变量
	setEnv()
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func setEnv() {
	os.Setenv("VERSION", "1.0.0")
}

// 请求处理
func requestHandle(w http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/healthz":
		healthCheck(w, r)
	case "/header":
		injectResquestHeaders(w, r)
	case "/env":
		getEnv(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// 获取环境变量
func getEnv(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	w.Header().Set("X-VERSION", version)
}

// 处理请求头
func injectResquestHeaders(w http.ResponseWriter, r *http.Request) {
	for name, values := range r.Header {
		for _, value := range values {
			w.Header().Set(name, value)
		}
	}
	w.WriteHeader(http.StatusOK)
}

// 健康检查
func healthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
	w.WriteHeader(http.StatusOK)
}
