package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":80"
)

var calls = 0
var cal1 = 0

// 定义一个 HelloWorld 函数，打印出访问的次数，需要传入两个参数
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return // 不处理非根路径的请求（如 /favicon.ico）
	}
	calls++
	// 访问递增，并显示出相关信息
	fmt.Fprintf(w, "Hello, world! You have called me %d times.\n", calls)
}

func HiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hi" {
		return // 不处理非根路径的请求（如 /favicon.ico）
	}
	cal1++
	fmt.Fprintf(w, "Hello, world! You have called me %d times.\n", cal1)
}

// 初始化函数，先执行
func init() {
	// 打印相关信息
	fmt.Printf("Started server at http://localhost%v.\n", port)

	// 显示定义路由
	mux := http.NewServeMux()

	// 调用http包的 HandleFunc 函数，需要一个字符串和一个函数;访问/就回调用一次函数
	mux.HandleFunc("/", HelloWorld)
	mux.HandleFunc("/hi", HiHandler)

	// 启动监听;调用http包的 ListenAndServe 函数 ，需要传入两个参数，字符串和 http.Handler，这是http包下的一个接口，不明白了....
	http.ListenAndServe(port, mux)
}

func main() {}
