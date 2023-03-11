package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/delivery"
	"strings"
)

type Server struct {
	router *gin.Engine
	config *config.Config

	auth   delivery.AuthController
	user   delivery.UserController
	tag    delivery.TagController
	sound  delivery.SoundController
	file   delivery.FileController
	search delivery.SearchController
}

func NewHttpServer(
	config *config.Config,
	auth delivery.AuthController,
	user delivery.UserController,
	tag delivery.TagController,
	sound delivery.SoundController,
	file delivery.FileController,
	search delivery.SearchController) *Server {

	s := &Server{
		router: gin.Default(),
		config: config,
		auth:   auth,
		user:   user,
		tag:    tag,
		sound:  sound,
		file:   file,
		search: search,
	}

	s.configureRouter()
	return s
}

func (s *Server) Run() error {
	err := s.router.Run(fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) authMiddleware(ctx *gin.Context) {
	if strings.HasSuffix(ctx.Request.RequestURI, "signin") ||
		strings.HasSuffix(ctx.Request.RequestURI, "register") {
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

func (s *Server) corsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	ctx.Next()
}

func (s *Server) configureRouter() {
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

	s.router.POST("/api/v1/search", s.searchSounds)

	s.router.POST("/api/v1/upload_image", s.uploadImage)
	s.router.POST("/api/v1/upload_sound", s.uploadSound)
}
