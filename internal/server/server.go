package server

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/csc13010-student-management/config"
	"github.com/csc13010-student-management/internal/auth"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	cfg *config.Config
	lg  *logger.LoggerZap
	pg  *gorm.DB
	rd  *redis.Client
	r   *gin.Engine
	e   *casbin.Enforcer
	w   *ServerWorker
}

type ServerWorker struct {
	authWorker auth.IAuthWorker
	stdWorker  student.IStudentWorker
}

func NewServer(
	cfg *config.Config,
	lg *logger.LoggerZap,
	pg *gorm.DB,
	rd *redis.Client,
	e *casbin.Enforcer,
) *Server {
	return &Server{
		cfg: cfg,
		lg:  lg,
		pg:  pg,
		rd:  rd,
		r:   newGinServer(cfg.Server),
		e:   e,
		w:   &ServerWorker{},
	}
}

func newGinServer(cfg config.ServerConfig) *gin.Engine {
	if cfg.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
	}

	r := gin.Default()

	// Thêm CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://example.com"},     // Các domain được phép
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}, // Các phương thức HTTP được phép
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},          // Các header được phép
		ExposeHeaders:    []string{"Content-Length"},                                   // Header được expose
		AllowCredentials: true,                                                         // Cho phép gửi cookie
		// MaxAge:           12 * time.Hour,                                          // Cache CORS trong bao lâu
	}))

	return r
}

func (s *Server) Run() error {
	if err := s.MapHandlers(); err != nil {
		return err
	}

	s.StartWorker()

	s.r.Run(fmt.Sprintf(":%v", s.cfg.Server.Port))
	return nil
}
