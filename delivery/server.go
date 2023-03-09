package delivery

import (
	"fmt"
	"log"

	"github.com/rizkyfazri23/dripay/config"
	"github.com/rizkyfazri23/dripay/controller"
	"github.com/rizkyfazri23/dripay/manager"
	"github.com/gin-gonic/gin"
)

type AppServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (p *AppServer) v1() {
	v1Routes := p.engine.Group("/v1")
	p.gatewayController(v1Routes)
}

func (p *AppServer) gatewayController(rg *gin.RouterGroup) {
	controller.NewGatewayController(rg, p.usecaseManager.GatewayUsecase())
}

func (p *AppServer) Run() {
	p.v1()
	err := p.engine.Run(p.host)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application failed to run", err)
		}
	}()
	if err != nil {
		panic(err)
	}
}

func Server() *AppServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf(":%s", c.ApiPort)
	return &AppServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
