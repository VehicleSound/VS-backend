package http

import (
	"github.com/gin-gonic/gin"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/delivery"
	"strings"
)

type HttpServer struct {
	router  *gin.Engine
	config  *config.Config
	running bool

	auth  delivery.AuthController
	user  delivery.UserController
	tag   delivery.TagController
	sound delivery.SoundController
	file  delivery.FileController
}

func NewHttpServer(
	config *config.Config,
	auth delivery.AuthController,
	user delivery.UserController,
	tag delivery.TagController,
	sound delivery.SoundController,
	file delivery.FileController) *HttpServer {

	s := &HttpServer{
		router: gin.Default(),
		config: config,
		auth:   auth,
		user:   user,
		tag:    tag,
		sound:  sound,
		file:   file,
	}

	s.configureRouter()
	return s
}

func (s *HttpServer) Run() error {
	err := s.router.Run("localhost:8080")
	if err != nil {
		return err
	}

	s.running = true
	return nil
}

func (s *HttpServer) authMiddleware(ctx *gin.Context) {
	if ctx.Request.RequestURI == "/api/v1/signin" {
		ctx.Next()
		return
	}
	hv := ctx.Request.Header.Get("Authorization")

	if hv != "" && len(strings.Split(hv, " ")) == 2 {
		token := strings.Split(hv, " ")[1]
		resp, err := s.auth.ValidateToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(401, &ErrorResponse{
				Code:    401,
				Message: err.Error(),
			})
			return
		}

		ctx.Set("user", resp)
		ctx.Next()
		return
	}

	ctx.AbortWithStatusJSON(401, &ErrorResponse{
		Code:    401,
		Message: "Wrong auth token",
	})
}

func (s *HttpServer) corsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}

func (s *HttpServer) configureRouter() {
	s.router.Use(s.corsMiddleware)
	s.router.Use(s.authMiddleware)

	s.router.Static("/assets/images", "./static/images")
	s.router.Static("/assets/sounds", "./static/sounds")

	s.router.POST("/api/v1/register", s.register)
	s.router.POST("/api/v1/signin", s.login)

	s.router.GET("/api/v1/users/:id", s.getUserById)
	s.router.POST("/api/v1/users/search", s.getUserByCredentials)

	s.router.POST("/api/v1/tags", s.createTag)
	s.router.GET("/api/v1/tags", s.getAllTags)
	s.router.GET("/api/v1/tags/:id", s.getTagById)

	s.router.GET("/api/v1/sounds", s.getAllSounds)
	s.router.GET("/api/v1/sounds/:id", s.getSoundById)
	s.router.POST("/api/v1/sounds", s.createSound)

	s.router.POST("/api/v1/upload_image", s.uploadImage)
	s.router.POST("/api/v1/upload_sound", s.uploadSound)
}
