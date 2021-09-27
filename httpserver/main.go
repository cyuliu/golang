package main
import (
  "io"
	"log"
	"net/http"
)

func main() {
  mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}

// 健康检查
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}
