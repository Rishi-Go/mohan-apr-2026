package handler

import (
	"blog_post/internals/dto"
	"blog_post/internals/service"
	"net/http"

	"github.com/gofiber/fiber"
)

type AuthHandler struct {
	Service service.AuthService
}

func InitHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{Service: svc}
}

func (h *AuthHandler) InsertSignUP(c *fiber.Ctx) {

	var res = dto.SignUpRequest{}

	if err := c.BodyParser(&res); err != nil {
		c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	err := h.Service.InsertSignUP(res)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	err = c.JSON(dto.ResponseMessage{Message: "INSERTED SUCCESSFULLY"})
	if err != nil {
		c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
		return
	}

}
