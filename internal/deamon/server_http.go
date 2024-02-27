package deamon

import (
	"github.com/gin-gonic/gin"
	"github.com/transaction-manager/internal/clients/postgrees"
	"github.com/transaction-manager/internal/config"
	"github.com/transaction-manager/internal/controllers"
	"github.com/transaction-manager/internal/routes"
	"github.com/transaction-manager/internal/usecases"
)

var (
	version   string
	gitCommit string
	buildID   string
)

func RunHttpServer(cfg *config.Config) {

	dbClient, err := postgrees.NewClient(cfg.RelationalDB.DSN(), cfg.RelationalDBConnection)
	if err != nil {
		panic(err)
	}
	// Repositories
	accountRepository := postgrees.NewAccountRepository(dbClient)
	transactionRepository := postgrees.NewTransactionRepository(dbClient)

	createAccount := usecases.NewCreateAccountUseCase(accountRepository)
	getAccount := usecases.NewGetAccountUseCase()
	createTransaction := usecases.NewCreateTransactionUseCase(transactionRepository, accountRepository)

	appInfoController := controllers.NewAppInfoHandler(version, gitCommit, buildID)
	accountController := controllers.NewAccountHandler(createAccount, getAccount)
	transactionController := controllers.NewTransactionHandler(createTransaction)

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	router := routes.NewRouter(r)
	router.RegisterRoutes(appInfoController, accountController, transactionController)
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
