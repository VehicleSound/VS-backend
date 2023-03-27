package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/timickb/transport-sound/internal/config"
	"github.com/timickb/transport-sound/internal/interfaces"
	"net/http"
	"regexp"
	"strings"
)

type Server struct {
	router *gin.Engine
	config *config.AppConfig

	auth   interfaces.AuthController
	user   interfaces.UserController
	tag    interfaces.TagController
	sound  interfaces.SoundController
	file   interfaces.FileController
	search interfaces.SearchController
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func NewHttpServer(
	config *config.AppConfig,
	auth interfaces.AuthController,
	user interfaces.UserController,
	tag interfaces.TagController,
	sound interfaces.SoundController,
	file interfaces.FileController,
	search interfaces.SearchController) *Server {

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
	err := s.router.Run(fmt.Sprintf(":%d", s.config.AppPort))

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) authMiddleware(ctx *gin.Context) {
	excludedURIs := []string{
		`^/assets`,
		`^/metrics`,
		`^/api/v1/random`,
		`^/api/v1/search`,
		`^/api/v1/sounds`,
		`^/api/v1/tags`,
		`^/api/v1/register`,
		`^/api/v1/signin`,
	}

	for _, uri := range excludedURIs {
		if matched, _ := regexp.MatchString(uri, ctx.Request.RequestURI); matched {
			ctx.Next()
			return
		}
	}

	tokenHeader := ctx.Request.Header.Get("Authorization")

	if tokenHeader != "" && len(strings.Split(tokenHeader, " ")) == 2 {
		token := strings.Split(tokenHeader, " ")[1]
		resp, err := s.auth.GetUserByToken(ctx, token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{
				Code:    401,
				Message: err.Error(),
			})
			return
		}

		ctx.Set("user", resp)
		ctx.Next()
		return
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{
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

	s.router.GET("/metrics", prometheusHandler())

	api := s.router.Group(fmt.Sprintf("/api/%s", ApiVersion))

	api.POST("/register", s.register)
	api.POST("/signin", s.login)
	api.GET("/me", s.me)

	api.GET("/users/:id", s.getUserById)
	api.POST("/users/search", s.getUserByCredentials)

	api.POST("/create_tag", s.createTag)
	api.GET("/tags", s.getAllTags)
	api.GET("/tags/:id", s.getTagById)

	api.GET("/sounds", s.getAllSounds)
	api.GET("/sounds/:id", s.getSoundById)
	api.POST("/create_sound", s.createSound)
	api.GET("/random", s.randomSounds)

	api.POST("/search", s.searchSounds)

	api.POST("/upload_image", s.uploadImage)
	api.POST("/upload_sound", s.uploadSound)

	api.POST("/add_favourite", s.addFavourite)
}
