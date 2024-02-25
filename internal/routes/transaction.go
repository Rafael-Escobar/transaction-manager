package routes

import "github.com/transaction-manager/internal/controllers"

func (r *Router) registerTransactionRoutes(
	controller *controllers.Transaction,
) {
	group := r.Router.Group("/transactions")

	// Create Transaction
	//	@BasePath	/transactions
	group.POST("", controller.CreateTransactionHandler)

}
