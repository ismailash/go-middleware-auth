package delivery

import (
	"fmt"
	"log"

	"enigmacamp.com/be-enigma-laundry/config"
	"enigmacamp.com/be-enigma-laundry/delivery/controller"
	"enigmacamp.com/be-enigma-laundry/delivery/middleware"
	"enigmacamp.com/be-enigma-laundry/manager"
	"enigmacamp.com/be-enigma-laundry/utils/common"
	"github.com/gin-gonic/gin"
)

type Server struct {
	uc         manager.UseCaseManager
	engine     *gin.Engine
	host       string
	logService common.MyLogger
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
	rg := s.engine.Group("/api/v1")
	controller.NewBillController(s.uc.BillUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infra, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := manager.NewRepoManager(infra)
	uc := manager.NewUseCaseManager(repo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.LogFileConfig)
	return &Server{
		uc:         uc,
		engine:     engine,
		host:       host,
		logService: logService,
	}
}
