package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	dto2 "github.com/timickb/transport-sound/internal/infrastructure/controller/dto"
)

func (s *Server) login(ctx *gin.Context) {
	req := dto2.AuthRequest{}

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

func (s *Server) register(ctx *gin.Context) {
	req := dto2.RegisterRequest{}

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

func (s *Server) createTag(ctx *gin.Context) {
	req := dto2.CreateTagRequest{}

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

func (s *Server) getAllTags(ctx *gin.Context) {
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

func (s *Server) getTagById(ctx *gin.Context) {
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

func (s *Server) getAllSounds(ctx *gin.Context) {
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

func (s *Server) getSoundById(ctx *gin.Context) {
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

func (s *Server) uploadImage(ctx *gin.Context) {
	req := &dto2.UploadFileRequest{}
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

func (s *Server) uploadSound(ctx *gin.Context) {
	req := &dto2.UploadFileRequest{}
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

func (s *Server) createSound(ctx *gin.Context) {
	req := &dto2.CreateSoundRequest{}
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
	}

	resp, err := s.sound.CreateSound(nil, req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	ctx.IndentedJSON(200, resp)
}

func (s *Server) getUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := s.user.GetUserById(id)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, resp)
}

func (s *Server) getUserByCredentials(ctx *gin.Context) {
	req := &dto2.GetUserRequest{}

	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
	}

	resp, err := s.user.GetUser(req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	ctx.IndentedJSON(200, resp)
}

func (s *Server) searchSounds(ctx *gin.Context) {
	req := &dto2.SearchRequest{}

	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
	}

	resp, err := s.search.Search(req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	ctx.IndentedJSON(200, resp)
}

func (s *Server) randomSounds(ctx *gin.Context) {
	resp, err := s.sound.GetRandomSounds(20)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	ctx.IndentedJSON(200, resp)
}

func (s *Server) me(ctx *gin.Context) {
	resp, ok := ctx.Get("user")
	if !ok {
		ctx.IndentedJSON(401, &ErrorResponse{
			Code:    401,
			Message: "Unauthorized",
		})
		return
	}

	ctx.IndentedJSON(200, resp.(*dto2.TokenResponse))
}

func (s *Server) addFavourite(ctx *gin.Context) {
	req := &dto2.AddToFavRequest{}

	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: "Invalid body",
		})
	}

	err := s.user.AddToFav(req)
	if err != nil {
		ctx.IndentedJSON(400, &ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(200, nil)
}
