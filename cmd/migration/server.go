package migration

import (
	"log"
	"template/config"
	"template/db"
	"template/internal/app"
	"template/internal/repository"
	"template/internal/usecase"
	"template/internal/usecase/scheduler"
)

func StartServer(cfg config.Config) {
	dbs, err := db.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(dbs)
	transactionRepo := repository.NewTransactionRepository(dbs)
	campaignRepo := repository.NewCampaignRepository(dbs)
	productRepo := repository.NewProductRepository(dbs)
	userUsecase := usecase.NewUserUsecase(userRepo, transactionRepo, campaignRepo, productRepo)
	transactionUcase := usecase.NewTransactionsUsecase(transactionRepo, userRepo, productRepo, campaignRepo)

	// run scheduler
	scheduler.RunCron(userUsecase, transactionUcase)

	// run server api
	app.Run(userUsecase, transactionUcase)
}
