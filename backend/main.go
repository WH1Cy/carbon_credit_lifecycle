package main

import (
	"log"
	"net/http"

	"go-carbon-backend/blockchain"
	"go-carbon-backend/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Println("========================================")
	log.Println("CarbonTrustConnect Backend Server")
	log.Println("========================================")

	// 1. 初始化区块链连接器
	connector, err := blockchain.NewConnector()
	if err != nil {
		log.Fatalf("Failed to initialize blockchain connector: %v", err)
	}

	// 2. 创建 handler 实例
	appHandler := &handlers.AppHandler{
		Connector: connector,
	}

	// 3. 创建路由
	r := mux.NewRouter()

	// =================================================================================
	// 数据查询 API (读操作)
	// =================================================================================
	// 平台概览
	r.HandleFunc("/api/summary", appHandler.GetSummaryHandler).Methods("GET")

	// 用户信息
	r.HandleFunc("/api/users", appHandler.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/api/users/{address}", appHandler.GetUserHistoryHandler).Methods("GET")

	// 项目信息
	r.HandleFunc("/api/projects", appHandler.GetAllProjectsHandler).Methods("GET")
	r.HandleFunc("/api/projects/{id}", appHandler.GetProjectHistoryHandler).Methods("GET")
	r.HandleFunc("/api/projects/{id}/reviews", appHandler.GetProjectReviewHistoryHandler).Methods("GET")

	// 碳信用信息
	r.HandleFunc("/api/credits", appHandler.GetAllCreditsHandler).Methods("GET")
	r.HandleFunc("/api/credits/{id}", appHandler.GetCreditHistoryHandler).Methods("GET")

	// 证书信息
	r.HandleFunc("/api/certificates", appHandler.GetAllCertificatesHandler).Methods("GET")
	r.HandleFunc("/api/certificates/{id}", appHandler.GetCertificateHistoryHandler).Methods("GET")

	// 4. 配置 CORS (跨域资源共享)
	// 这允许我们的前端 (通常在不同的源上) 调用后端 API
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // 在生产环境中应设为您的前端域名
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})
	handler := c.Handler(r)

	// 5. 启动服务器
	port := "8080"
	log.Printf("Starting Go backend server on port %s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
