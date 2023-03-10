package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/timickb/transport-sound/internal/controller/dto"
)

type HttpServer struct {
	router  *gin.Engine
	running bool

	auth  AuthController
	user  UserController
	tag   TagController
	sound SoundController
	file  FileController
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewHttpServer(
	auth AuthController,
	user UserController,
	tag TagController,
	sound SoundController,
	file FileController) *HttpServer {

	return &HttpServer{
		router: gin.Default(),
		auth:   auth,
		user:   user,
		tag:    tag,
		sound:  sound,
		file:   file,
	}
}

func (s *HttpServer) Run() error {
	s.router.Static("/assets/images", "./static/images")
	s.router.Static("/assets/sounds", "./static/sounds")

	s.router.POST("/api/v1/register", s.register)
	s.router.POST("/api/v1/signin", s.login)

	s.router.POST("/api/v1/tags", s.createTag)
	s.router.GET("/api/v1/tags", s.getAllTags)
	s.router.GET("/api/v1/tags/:id", s.getTagById)

	s.router.GET("/api/v1/sounds", s.getAllSounds)
	s.router.GET("/api/v1/sounds/:id", s.getSoundById)
	s.router.POST("/api/v1/sounds", s.createSound)

	s.router.POST("/api/v1/upload_image", s.uploadImage)
	s.router.POST("/api/v1/upload_sound", s.uploadSound)

	err := s.router.Run("localhost:8080")
	if err != nil {
		return err
	}

	s.running = true
	return nil
}

func (s *HttpServer) login(ctx *gin.Context) {
	req := dto.AuthRequest{}

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.auth.SignIn(&req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (s *HttpServer) register(ctx *gin.Context) {
	req := dto.RegisterRequest{}

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.user.Register(&req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (s *HttpServer) createTag(ctx *gin.Context) {
	req := dto.CreateTagRequest{}

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.tag.CreateTag(&req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (s *HttpServer) getAllTags(ctx *gin.Context) {
	resp, err := s.tag.GetAllTags()

	if err != nil {
		ctx.IndentedJSON(500, &ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (s *HttpServer) getTagById(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := s.tag.GetTagById(id)

	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (s *HttpServer) getAllSounds(ctx *gin.Context) {
	resp, err := s.sound.GetAllSounds()

	if err != nil {
		ctx.IndentedJSON(500, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (s *HttpServer) getSoundById(ctx *gin.Context) {
	id := ctx.Param("id")

	sound, err := s.sound.GetSoundById(id)
	if err != nil {
		ctx.IndentedJSON(500, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, sound)
}

func (s *HttpServer) uploadImage(ctx *gin.Context) {
	req := &dto.UploadFileRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.file.UploadImage(req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (s *HttpServer) uploadSound(ctx *gin.Context) {
	req := &dto.UploadFileRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.file.UploadSound(req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	ctx.IndentedJSON(200, resp)
}

func (s *HttpServer) createSound(ctx *gin.Context) {
	req := &dto.CreateSoundRequest{}
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
	}

	resp, err := s.sound.CreateSound(req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	ctx.IndentedJSON(200, resp)
}
