package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const (
	port = ":80"
)

// 用于记录每个路径访问次数的map。带互斥锁确保并发安全
var (
	stats = make(map[string]int) // 例如： /hi:7
	mu    sync.Mutex
)

// 中间件：统计访问次数，并调用原始handler,难点！其实只做了计数的统计，剩下的还是交给原来的函数做
func countAndHandle(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		stats[path]++
		mu.Unlock()
		handler(w, r)
	}
}

// 根路径处理器
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

// /hi 路径处理器
func HiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi there!")
}

// /stats 路径处理器：输出访问统计（JSON 格式）
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// /reset 路径处理器：重置所有统计
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	for k := range stats {
		stats[k] = 0
	}
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Stats reset.")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", countAndHandle("/", HelloWorld))
	mux.HandleFunc("/hi", countAndHandle("/hi", HiHandler))
	mux.HandleFunc("/stats", countAndHandle("/stats", StatsHandler))
	mux.HandleFunc("/reset", countAndHandle("/reset", ResetHandler))

	fmt.Printf("🚀 Server started at http://localhost%v\n", port)
	http.ListenAndServe(port, mux)
}
