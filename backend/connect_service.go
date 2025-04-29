package main

import (
	"net/http"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// 设置CORS并启动服务器
func setupAndServeAPI(mux *http.ServeMux, addr string) error {
	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
		AllowedOrigins: []string{"http://localhost:3000"}, // 前端地址
		AllowedHeaders: []string{
			"Accept",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"X-CSRF-Token",
			"Connect-Protocol-Version",
		},
	})

	// 使用h2c启用HTTP/2服务
	h2s := &http2.Server{}
	handler := h2c.NewHandler(corsHandler.Handler(mux), h2s)

	return http.ListenAndServe(addr, handler)
}
