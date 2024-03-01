package routes

import "github.com/transaction-manager/internal/controllers"

func (r *Router) registerTransactionRoutes(
	controller *controllers.TransactionHandler,
) {
	v1 := r.Router.Group("/v1")
	group := v1.Group("/transactions")

	// Create Transaction
	//	@BasePath	/transactions
	group.POST("", controller.CreateTransactionHandler)

}
