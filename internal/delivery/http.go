package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/timickb/transport-sound/internal/controller"
)

type HttpServer struct {
	router  *gin.Engine
	running bool
	auth    AuthController
	user    UserController
	tag     TagController
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewHttpServer(auth AuthController, user UserController, tag TagController) *HttpServer {
	router := gin.Default()

	return &HttpServer{
		router: router,
		auth:   auth,
		user:   user,
		tag:    tag,
	}
}

func (s *HttpServer) Run() error {
	s.router.POST("/register", s.register)
	s.router.POST("/signin", s.login)
	s.router.POST("/tags", s.createTag)
	s.router.GET("/tags", s.getAllTags)
	s.router.GET("/tags/:id", s.getTagById)

	err := s.router.Run("localhost:8080")
	if err != nil {
		return err
	}

	s.running = true
	return nil
}

func (s *HttpServer) login(ctx *gin.Context) {
	req := controller.AuthRequest{}

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
	req := controller.RegisterRequest{}

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
	req := controller.CreateTagRequest{}

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
