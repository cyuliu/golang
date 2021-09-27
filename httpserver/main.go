package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	path := []string{"/healthz", "/header", "/env", "/log"}
	for _, p := range path {
		mux.HandleFunc(p, requestHandle)
	}
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
	case "/log":
		localLog(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

// 处理日志
func localLog(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	fmt.Printf("IP is %v, Code is %v", r.Host, status)
	w.WriteHeader(status)
}

// 获取环境变量
func getEnv(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	w.Header().Set("X-VERSION", version)
	w.WriteHeader(http.StatusOK)
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
