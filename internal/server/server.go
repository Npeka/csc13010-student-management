package server

import (
	"fmt"

	"github.com/csc13010-student-management/config"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	gin *gin.Engine
	cfg *config.Config
	lg  *logger.LoggerZap
	pg  *gorm.DB
	rd  *redis.Client
}

func NewServer(
	cfg *config.Config,
	lg *logger.LoggerZap,
	pg *gorm.DB,
	rd *redis.Client,
) *Server {
	return &Server{
		gin: newGinServer(cfg.Server),
		cfg: cfg,
		lg:  lg,
		pg:  pg,
		rd:  rd,
	}
}

func newGinServer(cfg config.ServerConfig) *gin.Engine {
	if cfg.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		return gin.Default()
	}

	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.gin); err != nil {
		return err
	}

	s.gin.Run(fmt.Sprintf(":%v", s.cfg.Server.Port))
	return nil
}
