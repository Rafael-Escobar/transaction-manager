package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/transaction-manager/docs"
	"github.com/transaction-manager/internal/controllers"
)

type Router struct {
	Router *gin.Engine
}

func NewRouter(router *gin.Engine) *Router {
	return &Router{
		Router: router,
	}
}

func (r *Router) RegisterRoutes(
	appInfo *controllers.AppInfo,
	accountController *controllers.Account,
	transactionController *controllers.Transaction,
) {
	v1 := r.Router.Group("/v1")
	// App Info
	//	@BasePath	/app-info
	v1.GET("/app-info", appInfo.Handler)
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.registerAccountRoutes(accountController)
	r.registerTransactionRoutes(transactionController)
}
