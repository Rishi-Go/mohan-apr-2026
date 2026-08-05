package handler

import (
	"blog_post/internals/dto"
	"blog_post/internals/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type LikeHandler struct {
	Service service.LikeService
}

func InitLikeHandler(svc service.LikeService) *LikeHandler {
	return &LikeHandler{Service: svc}
}

func (h *LikeHandler) InsertLike(Ctx fiber.Ctx) error {

	var res = dto.LikeRequest{}

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)

	userid, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	res.UserID = userid

	if err := Ctx.Bind().Body(&res); err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.InsertLike(res)
	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = Ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"Message": "Liked successfully", "StatusCode": fiber.StatusCreated})
	if err != nil {

		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *LikeHandler) GetLike(Ctx fiber.Ctx) error {

	like := Ctx.Query("like")
	islike, err := strconv.ParseBool(like)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: "Conversion error", StatusCode: http.StatusBadRequest})
	}

	userStr := Ctx.Query("user-id")

	userid := uuid.FromStringOrNil(userStr)

	blogStr := Ctx.Query("blog-id")

	blogid := uuid.FromStringOrNil(blogStr)

	page, err := strconv.Atoi(Ctx.Query("page"))

	if page < 1 {
		page = 1
	} else if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	limit, err := strconv.Atoi(Ctx.Query("limit"))

	if limit > 99 || limit < 0 {

		return Ctx.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{Message: "Limit should be with in 1 - 99", StatusCode: http.StatusInternalServerError})
	} else if limit == 0 {

		limit = 10

	} else if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	offset := (page - 1) * limit

	result, Page, err := h.Service.GetLike(page, limit, offset, islike, userid, blogid)
	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = Ctx.JSON(&dto.LikeResponse{
		Like:       result,
		Pagination: *Page,
	})

	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	return nil

}

func (h *LikeHandler) SelectLike(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	LikeId, err := uuid.FromString(uuidStr)

	if err != nil {

		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	ID, err := h.Service.SelectLike(LikeId)
	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = Ctx.JSON(ID)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *LikeHandler) DeleteLike(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	LikeId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)
	role := claims["role"].(string)

	LoginUser, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.DeleteLike(LikeId, LoginUser, role)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	err = Ctx.JSON(dto.Response{Message: "DELETED SUCCESSFULLY", ID: LikeId})
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}
