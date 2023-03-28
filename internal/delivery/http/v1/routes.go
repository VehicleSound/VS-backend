package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/timickb/transport-sound/internal/infrastructure/controller/dto"
	"net/http"
)

func (s *Server) login(ctx *gin.Context) {
	req := dto.AuthRequest{}

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.auth.SignIn(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) register(ctx *gin.Context) {
	req := dto.RegisterRequest{}

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.user.Register(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) createTag(ctx *gin.Context) {
	req := dto.CreateTagRequest{}

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.tag.CreateTag(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) getAllTags(ctx *gin.Context) {
	resp, err := s.tag.GetAllTags(ctx)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) getTagById(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := s.tag.GetTagById(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) getAllSounds(ctx *gin.Context) {
	resp, err := s.sound.GetAllSounds(ctx)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) getSoundById(ctx *gin.Context) {
	id := ctx.Param("id")

	sound, err := s.sound.GetSoundById(ctx, id)
	if err != nil {
		ctx.IndentedJSON(500, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    sound,
	})
}

func (s *Server) uploadImage(ctx *gin.Context) {
	req := &dto.UploadFileRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.file.UploadImage(ctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) uploadSound(ctx *gin.Context) {
	req := &dto.UploadFileRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
		return
	}

	resp, err := s.file.UploadSound(ctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) createSound(ctx *gin.Context) {
	req := &dto.CreateSoundRequest{}
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
	}

	resp, err := s.sound.CreateSound(ctx, nil, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) getUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := s.user.GetUserById(ctx, id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) getUserByCredentials(ctx *gin.Context) {
	req := &dto.GetUserRequest{}

	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
	}

	resp, err := s.user.GetUser(ctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) searchSounds(ctx *gin.Context) {
	req := &dto.SearchRequest{}

	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
	}

	resp, err := s.search.Search(ctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) randomSounds(ctx *gin.Context) {
	resp, err := s.sound.GetRandomSounds(ctx, 20)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp,
	})
}

func (s *Server) me(ctx *gin.Context) {
	resp, ok := ctx.Get("user")
	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, Response{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
		Data:    resp.(*dto.TokenResponse),
	})
}

func (s *Server) addFavourite(ctx *gin.Context) {
	req := &dto.AddToFavRequest{}

	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
	}

	err := s.user.AddToFav(ctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: SuccessMessage,
	})
}
