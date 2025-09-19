package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Response 定义返回的 JSON 结构
type Response struct {
	Message string `json:"message"`
	Version string `json:"version"`
	Host    string `json:"host"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 获取当前主机名（Pod 名称）
	host, _ := os.Hostname()

	resp := Response{
		Message: "Hello from Go! 云效部署成功 ✅",
		Version: "1.0.0",
		Host:    host,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认端口
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusMovedPermanently)
	})

	log.Printf("Server is listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
