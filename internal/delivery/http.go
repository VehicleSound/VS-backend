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
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewHttpServer(auth AuthController, user UserController) *HttpServer {
	router := gin.Default()

	return &HttpServer{router: router, auth: auth, user: user}
}

func (s *HttpServer) Run() error {
	s.router.POST("/register", s.register)
	s.router.POST("/signin", s.login)

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
	}

	resp, err := s.auth.SignIn(&req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
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
	}

	ctx.IndentedJSON(200, resp)
}
