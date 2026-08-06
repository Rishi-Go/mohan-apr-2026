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

type CommentHandler struct {
	Service service.CommentService
}

func InitCommentHandler(svc service.CommentService) *CommentHandler {
	return &CommentHandler{Service: svc}
}

func (h *CommentHandler) InsertComment(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	BlogId, err := uuid.FromString(uuidStr)

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	var res = dto.CommentRequest{}

	res.BlogID = BlogId

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)

	ID, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	res.UserID = ID

	if err := Ctx.Bind().Body(&res); err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	result, err := h.Service.InsertComment(res)
	if err != nil {
		logger.Log.Error("Failed to add Comment", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Comment added successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.CommentInsertResponse{
			Comments: result,
		},
	})
	logger.Log.Info("Comment added successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *CommentHandler) GetComment(Ctx fiber.Ctx) error {

	comment := Ctx.Query("comment")

	UserStr := Ctx.Query("user-id")

	userid := uuid.FromStringOrNil(UserStr)

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

	result, Page, err := h.Service.GetComment(page, limit, offset, comment, userid, blogid)
	if err != nil {
		logger.Log.Error("Failed to retreived comment details", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: &dto.CommentInsertResponses{
			Comments: result,
		}})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Comment details retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.CommentResponse{
			Comments:   result,
			Pagination: *Page,
		},
	})
	logger.Log.Info("Comment details retreived successfully")

	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	return nil

}

func (h *CommentHandler) SelectComment(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	CommentId, err := uuid.FromString(uuidStr)

	if err != nil {

		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	comment, err := h.Service.SelectComment(CommentId)
	if err != nil {
		logger.Log.Error("Failed to retreived comment detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorMessage{Message: err.Error(), StatusCode: http.StatusNotFound, ID: CommentId})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Comment detail retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.CommentInsertResponse{
			Comments: comment,
		},
	})
	logger.Log.Info("Comment detail retreived successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *CommentHandler) UpdateComment(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	CommentId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	var res = dto.CommentRequest{}

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)

	ID, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	res.UserID = ID

	if err := Ctx.Bind().Body(&res); err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.UpdateComment(res, CommentId)
	if err != nil {
		logger.Log.Error("Failed to update comment detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Update Comment Record", ID: CommentId}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Comment Updated successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Updated Record", ID: CommentId,
		}})
	logger.Log.Info("Comment Updated successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *CommentHandler) DeleteComment(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	CommentId, err := uuid.FromString(uuidStr)
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

	err = h.Service.DeleteComment(CommentId, LoginUser, role)
	if err != nil {
		logger.Log.Error("Failed to delete comment detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Delete Comment Record", ID: CommentId}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Comment Deleted successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Deleted Record", ID: CommentId,
		}})
	logger.Log.Info("Comment Deleted successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}
