package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"automobile-parts-backend/config"
	"automobile-parts-backend/controller"
	"automobile-parts-backend/middleware"
	"automobile-parts-backend/service"
	"automobile-parts-backend/utils"
)

func init() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("[INFO] 未找到 .env 文件，将使用环境变量或默认值")
	} else {
		log.Println("[INFO] 已加载 .env 文件")
	}
}

// setupRouter 配置所有API路由和中间件
// 参数：
//   - cfg: 应用配置信息
//   - 各控制器实例：处理不同业务模块的API请求
//
// 返回：
//   - *gin.Engine: 配置完成的Gin引擎实例
func setupRouter(
	cfg config.Config,
	authController *controller.AuthController,
	partController *controller.PartController,
	bomController *controller.BOMController,
	productionController *controller.ProductionController,
	qualityController *controller.QualityController,
	supplyChainController *controller.SupplyChainController,
	aftersaleController *controller.AftersaleController,
	fabricController *controller.FabricController,
	aiController *controller.AIController,
) *gin.Engine {
	// 创建Gin引擎实例
	router := gin.New()

	// 注册全局中间件
	router.Use(middleware.Logger())       // 日志中间件：记录所有API请求
	router.Use(gin.Recovery())            // 恢复中间件：捕获panic并返回500错误
	router.Use(middleware.CORS())         // CORS中间件：处理跨域请求
	router.Use(middleware.ErrorHandler()) // 错误处理中间件：统一处理API错误

	// 健康检查API：用于监控系统运行状态
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.Success(map[string]string{"status": "ok"}, "ok"))
	})

	// 认证API组：不需要身份验证
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", authController.Login)       // 用户登录
		auth.POST("/register", authController.Register) // 用户注册
	}

	// 零部件API组：需要身份验证
	parts := router.Group("/api/parts")
	parts.Use(middleware.Auth(cfg)) // 应用JWT认证中间件
	{
		parts.POST("", partController.CreatePart)                    // 创建零部件
		parts.GET("", partController.ListParts)                      // 列出零部件（按批次号或VIN）
		parts.GET("/my", partController.ListMyParts)                 // 列出我的零部件
		parts.GET("/:id", partController.GetPart)                    // 获取单个零部件
		parts.GET("/:id/lifecycle", partController.GetPartLifecycle) // 获取零部件生命周期
		parts.DELETE("/:id", partController.DeletePart)              // 删除零部件
	}

	// BOM API组：需要身份验证
	boms := router.Group("/api/boms")
	boms.Use(middleware.Auth(cfg))
	{
		boms.POST("", bomController.CreateBOM)       // 创建BOM
		boms.GET("", bomController.ListBOMs)         // 列出BOM列表
		boms.GET("/:id", bomController.GetBOM)       // 获取单个BOM
		boms.PUT("/:id", bomController.UpdateBOM)    // 更新BOM
		boms.DELETE("/:id", bomController.DeleteBOM) // 删除BOM
	}

	// 生产数据API组：需要身份验证
	production := router.Group("/api/production")
	production.Use(middleware.Auth(cfg))
	{
		production.POST("", productionController.CreateProductionData) // 创建生产数据
		production.GET("", productionController.ListProductionData)    // 列出生产数据
	}

	// 质检数据API组：需要身份验证
	quality := router.Group("/api/quality")
	quality.Use(middleware.Auth(cfg))
	{
		quality.POST("", qualityController.CreateQualityInspection) // 创建质检数据
		quality.GET("", qualityController.ListQualityInspections)   // 列出质检数据
	}

	// 采购订单API组：需要身份验证
	orders := router.Group("/api/orders")
	orders.Use(middleware.Auth(cfg))
	{
		orders.POST("", supplyChainController.CreateSupplyOrder) // 创建采购订单
		orders.GET("", supplyChainController.ListSupplyOrders)   // 列出采购订单
	}

	// 物流数据API组：需要身份验证
	logistics := router.Group("/api/logistics")
	logistics.Use(middleware.Auth(cfg))
	{
		logistics.POST("", supplyChainController.CreateLogisticsData) // 创建物流数据
		logistics.GET("", supplyChainController.ListLogisticsData)    // 列出物流数据
	}

	// 故障报告API组：需要身份验证
	faults := router.Group("/api/faults")
	faults.Use(middleware.Auth(cfg))
	{
		faults.POST("", aftersaleController.CreateFaultReport) // 创建故障报告
		faults.GET("", aftersaleController.ListFaultReports)   // 列出故障报告
	}

	// 召回记录API组：需要身份验证
	recalls := router.Group("/api/recalls")
	recalls.Use(middleware.Auth(cfg))
	{
		recalls.POST("", aftersaleController.CreateRecallRecord) // 创建召回记录
		recalls.GET("", aftersaleController.ListRecallRecords)   // 列出召回记录
	}

	// 售后记录API组：需要身份验证
	aftersaleRecords := router.Group("/api/aftersale-records")
	aftersaleRecords.Use(middleware.Auth(cfg))
	{
		aftersaleRecords.POST("", aftersaleController.CreateAftersaleRecord) // 创建售后记录
		aftersaleRecords.GET("", aftersaleController.ListAftersaleRecords)   // 列出售后记录
	}

	// 制造商API组：需要身份验证和制造商角色权限
	manufacturer := router.Group("/api/manufacturer")
	manufacturer.Use(middleware.Auth(cfg))                  // JWT认证
	manufacturer.Use(middleware.Permission("manufacturer")) // 制造商权限控制
	{
		// 零部件管理
		manufacturer.POST("/parts", partController.CreatePart)
		manufacturer.GET("/parts", partController.ListParts)
		manufacturer.GET("/parts/:id", partController.GetPart)
		// BOM管理
		manufacturer.POST("/boms", bomController.CreateBOM)
		manufacturer.GET("/boms/:id", bomController.GetBOM)
		// 生产数据管理
		manufacturer.POST("/production", productionController.CreateProductionData)
		// 质检数据管理
		manufacturer.POST("/quality", qualityController.CreateQualityInspection)
	}

	// 整车车企API组：需要身份验证和车企角色权限
	automaker := router.Group("/api/automaker")
	automaker.Use(middleware.Auth(cfg))               // JWT认证
	automaker.Use(middleware.Permission("automaker")) // 车企权限控制
	{
		// 供应链管理
		automaker.POST("/orders", supplyChainController.CreateSupplyOrder)      // 创建采购订单
		automaker.POST("/logistics", supplyChainController.CreateLogisticsData) // 创建物流数据
	}

	// 售后中心API组：需要身份验证和售后角色权限
	aftersale := router.Group("/api/aftersale")
	aftersale.Use(middleware.Auth(cfg))               // JWT认证
	aftersale.Use(middleware.Permission("aftersale")) // 售后权限控制
	{
		// 售后记录管理
		aftersale.POST("/faults", aftersaleController.CreateFaultReport)      // 创建故障报告
		aftersale.POST("/recalls", aftersaleController.CreateRecallRecord)    // 创建召回记录
		aftersale.POST("/records", aftersaleController.CreateAftersaleRecord) // 创建售后记录
	}

	// Fabric区块链API组：提供直接的Fabric区块链访问接口
	if fabricController != nil {
		fabricController.RegisterRoutes(router) // 注册Fabric相关路由
	}

	// AI智能问答API组：所有角色都可以访问
	ai := router.Group("/api/ai")
	{
		ai.POST("/question", aiController.AskQuestion)             // 提问
		ai.GET("/conversation", aiController.GetConversation)      // 获取对话历史
		ai.DELETE("/conversation", aiController.ClearConversation) // 清空对话
		ai.GET("/stats", aiController.GetStats)                    // 获取统计信息
		ai.GET("/health", aiController.HealthCheck)                // 健康检查
	}

	return router
}

// main 应用程序入口函数
func main() {
	log.Println("========================================")
	log.Println("  汽车零件管理系统 - 后端服务")
	log.Println("========================================\n")

	log.Println("步骤1: 加载应用配置...")
	cfg := config.Load()
	log.Println("  ✅ 配置加载成功")

	log.Println("\n步骤2: 初始化Fabric服务...")
	var fabricService *service.FabricService
	var err error

	if cfg.FabricEnabled {
		log.Println("  Fabric功能已启用，正在连接...")
		fabricService, err = service.NewFabricService(cfg)
		if err != nil {
			log.Printf("  ⚠️  Fabric服务初始化失败: %v", err)
			log.Println("  应用将以不包含Fabric功能的方式启动")
			fabricService = nil
		} else {
			log.Println("  ✅ Fabric服务初始化成功")
		}
	} else {
		log.Println("  Fabric功能未启用")
	}

	log.Println("\n步骤3: 初始化服务层...")
	authService := service.NewAuthService(cfg)
	partService := service.NewPartService(fabricService)
	bomService := service.NewBOMService(fabricService)
	productionService := service.NewProductionService(fabricService)
	qualityService := service.NewQualityService(fabricService)
	supplyChainService := service.NewSupplyChainService(fabricService)
	aftersaleService := service.NewAftersaleService(fabricService)
	aiService := service.NewAIService(cfg.AIServiceURL)
	log.Println("  ✅ 服务层初始化成功")

	log.Println("\n步骤4: 初始化控制器层...")
	authController := controller.NewAuthController(authService)
	partController := controller.NewPartController(partService)
	bomController := controller.NewBOMController(bomService)
	productionController := controller.NewProductionController(productionService)
	qualityController := controller.NewQualityController(qualityService)
	supplyChainController := controller.NewSupplyChainController(supplyChainService)
	aftersaleController := controller.NewAftersaleController(aftersaleService)
	fabricController := controller.NewFabricController(fabricService)
	aiController := controller.NewAIController(aiService)
	log.Println("  ✅ 控制器层初始化成功")

	log.Println("\n步骤5: 设置路由...")
	router := setupRouter(
		cfg,
		authController,
		partController,
		bomController,
		productionController,
		qualityController,
		supplyChainController,
		aftersaleController,
		fabricController,
		aiController,
	)
	log.Println("  ✅ 路由设置成功")

	log.Println("\n步骤6: 启动HTTP服务器...")
	addr := ":" + cfg.ServerPort
	if addr == ":" {
		addr = ":8080"
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		log.Printf("  🚀 服务器启动成功，监听地址: %s\n", addr)
		if fabricService != nil {
			log.Println("  Fabric功能已启用")
		} else {
			log.Println("  Fabric功能未启用")
		}
		log.Println("  按 Ctrl+C 停止服务器\n")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("\n步骤7: 正在关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("服务器关闭失败: %v", err)
	}

	if fabricService != nil {
		log.Println("正在关闭Fabric服务...")
		if err := fabricService.Close(); err != nil {
			log.Printf("关闭Fabric服务时出错: %v", err)
		}
	}

	log.Println("✅ 服务器已停止")
}
