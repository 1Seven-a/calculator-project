package main

import (
	"fmt"
	"log"
	"net/http"

	"calculator/gen/proto/calculatorv1connect"
)

func main() {
	calculatorServer := &CalculatorServer{}
	mux := http.NewServeMux()

	// 注册计算器服务
	path, handler := calculatorv1connect.NewCalculatorServiceHandler(calculatorServer)
	mux.Handle(path, handler)

	fmt.Println("计算器服务器启动在 :8080...")
	err := setupAndServeAPI(mux, ":8080")
	if err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
