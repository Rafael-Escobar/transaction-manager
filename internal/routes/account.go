package routes

import "github.com/transaction-manager/internal/controllers"

func (r *Router) registerAccountRoutes(
	controller *controllers.AccountHandler,
) {
	v1 := r.Router.Group("/v1")
	group := v1.Group("/accounts")

	// Get Account
	//	@BasePath	/accounts/{id}
	group.GET("/:id", controller.GetAccountHandler)

	// Create Account
	//	@BasePath	/accounts
	group.POST("", controller.CreateAccountHandler)

}
