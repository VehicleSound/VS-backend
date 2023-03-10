package http

import (
	"github.com/gin-gonic/gin"
	"github.com/timickb/transport-sound/internal/delivery"
)

type HttpServer struct {
	router  *gin.Engine
	running bool

	auth  delivery.AuthController
	user  delivery.UserController
	tag   delivery.TagController
	sound delivery.SoundController
	file  delivery.FileController
}

func NewHttpServer(
	auth delivery.AuthController,
	user delivery.UserController,
	tag delivery.TagController,
	sound delivery.SoundController,
	file delivery.FileController) *HttpServer {

	s := &HttpServer{
		router: gin.Default(),
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

func (s *HttpServer) configureRouter() {
	s.router.Static("/assets/images", "./static/images")
	s.router.Static("/assets/sounds", "./static/sounds")

	s.router.POST("/api/v1/register", s.register)
	s.router.POST("/api/v1/signin", s.login)

	s.router.GET("/user/:id", s.getUserById)
	s.router.POST("/user/search", s.getUserByCredentials)

	s.router.POST("/api/v1/tags", s.createTag)
	s.router.GET("/api/v1/tags", s.getAllTags)
	s.router.GET("/api/v1/tags/:id", s.getTagById)

	s.router.GET("/api/v1/sounds", s.getAllSounds)
	s.router.GET("/api/v1/sounds/:id", s.getSoundById)
	s.router.POST("/api/v1/sounds", s.createSound)

	s.router.POST("/api/v1/upload_image", s.uploadImage)
	s.router.POST("/api/v1/upload_sound", s.uploadSound)
}
