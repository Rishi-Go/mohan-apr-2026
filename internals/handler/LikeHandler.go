package handler

import (
	"blog_post/internals/dto"
	"blog_post/internals/service"
	"blog_post/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type LikeHandler struct {
	Service service.LikeService
}

func InitLikeHandler(svc service.LikeService) *LikeHandler {
	return &LikeHandler{Service: svc}
}

func (h *LikeHandler) InsertLike(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	blogid, err := uuid.FromString(uuidStr)

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	var res = dto.LikeRequest{}

	res.BlogID = blogid

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)

	userid, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	res.UserID = userid

	if err := Ctx.Bind().Body(&res); err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	result, err := h.Service.InsertLike(res)
	if err != nil {
		logger.Log.Error("Failed to insert like", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Like added successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.LikeInsertResponse{
			Like: result,
		},
	})
	logger.Log.Info("Like added successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *LikeHandler) GetLike(Ctx fiber.Ctx) error {

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

	result, Page, err := h.Service.GetLike(page, limit, offset, userid, blogid)
	if err != nil {
		logger.Log.Error("Failed to retreived like details", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: &dto.LikeInsertResponses{
			Like: result,
		}})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Like details retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.LikeResponse{
			Like:       result,
			Pagination: *Page,
		},
	})
	logger.Log.Info("Like details retreived successfully")

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
		logger.Log.Error("Failed to retreived like detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorMessage{Message: err.Error(), StatusCode: http.StatusNotFound, ID: LikeId})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Like detail retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.LikeInsertResponse{
			Like: ID,
		},
	})
	logger.Log.Info("Like detail retreived successfully")

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

	UserId, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.DeleteLike(LikeId, UserId)
	if err != nil {
		logger.Log.Error("Failed to Delete like detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Delete Like Record", ID: UserId}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Like Removed successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Delete Record", ID: UserId,
		}})
	logger.Log.Info("Like Removed successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}
