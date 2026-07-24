package handler

import (
	"blog_post/internals/dto"
	"blog_post/internals/service"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type LikeHandler struct {
	Service service.LikeService
}

func InitLikrHandler(svc service.LikeService) *LikeHandler {
	return &LikeHandler{Service: svc}
}

func (h *LikeHandler) InsertLike(Ctx fiber.Ctx) error {

	var res = dto.LikeRequest{}

	if err := Ctx.Bind().Body(&res); err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err := h.Service.InsertLike(res)
	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = Ctx.JSON(dto.ResponseMessage{Message: "INSERTED SUCCESSFULLY"})
	if err != nil {

		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}
