package main

import (
	balanceService "get-block/application/balance"
	"get-block/config"
	_ "get-block/docs"
	"get-block/domain/balance"
	"get-block/infrastructure/clients/getBlock"
	"get-block/ui"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
	"log"
)

type Application struct {
	router Router
}

type Router struct {
}

func (application *Application) Run(cnf *config.Config) {
	container := application.BuildContainer(*cnf)
	var err error

	err = application.runWepApp(container)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (application *Application) runWepApp(container *dig.Container) error {
	ginEngine := gin.Default()

	err := container.Invoke(
		func(
			GitLabJobController *ui.BalanceController,
		) {
			application.router.initHandlers(
				ginEngine,
				GitLabJobController,
			)
		})

	if err != nil {
		return err
	}
	log.Println("listen 0.0.0.0:8091")

	return ginEngine.Run("0.0.0.0:8091") // listen and serve on 0.0.0.0:8091
}

func (r *Router) initHandlers(
	ginEngine *gin.Engine,
	BalanceController *ui.BalanceController,
) {
	ginEngine.GET("/most-changed-balance", BalanceController.GetMostChangedBalanceAddress)
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (application *Application) BuildContainer(cnf config.Config) *dig.Container {
	container := dig.New()

	err := container.Provide(func() config.Config { return cnf })

	// Application
	err = container.Provide(balanceService.NewService)

	// Infrastructure
	err = container.Provide(func(cnf config.Config) balance.GetBlockClientInterface {
		return getBlock.NewClient(getBlock.NewConfig(cnf.GetBlock.BaseUrl, cnf.GetBlock.Key))
	})

	// UI
	err = container.Provide(ui.NewBalanceController)

	if err != nil {
		panic(err)
	}

	return container
}
