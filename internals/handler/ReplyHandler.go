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

type ReplyHandler struct {
	Service service.ReplyService
}

func InitReplyHandler(svc service.ReplyService) *ReplyHandler {
	return &ReplyHandler{Service: svc}
}

func (h *ReplyHandler) InsertReply(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	commentid, err := uuid.FromString(uuidStr)

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	var res = dto.ReplyRequest{}

	res.CommentID = commentid

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)

	authorid, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	res.UserID = authorid

	if err := Ctx.Bind().Body(&res); err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	result, err := h.Service.InsertReply(res)
	if err != nil {
		logger.Log.Error("Failed to add reply", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Reply added successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.ReplyInsertResponse{
			Reply: result,
		},
	})
	logger.Log.Info("Reply added successfully")

	if err != nil {

		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *ReplyHandler) GetReply(Ctx fiber.Ctx) error {

	reply := Ctx.Query("reply")

	userStr := Ctx.Query("user-id")

	userid := uuid.FromStringOrNil(userStr)

	commentStr := Ctx.Query("comment-id")

	commentid := uuid.FromStringOrNil(commentStr)

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

	result, Page, err := h.Service.GetReply(page, limit, offset, reply, commentid, userid)
	if err != nil {
		logger.Log.Error("Failed to retreive reply details", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: &dto.ReplyInsertResponses{
			Reply: result,
		}})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Reply details retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.ReplyResponse{
			Reply:      result,
			Pagination: *Page,
		},
	})
	logger.Log.Info("Reply details retreived successfully")

	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	return nil

}

func (h *ReplyHandler) SelectReply(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	replyId, err := uuid.FromString(uuidStr)

	if err != nil {

		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	ID, err := h.Service.SelectReply(replyId)
	if err != nil {
		logger.Log.Error("Failed to retreive reply detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorMessage{Message: err.Error(), StatusCode: http.StatusNotFound, ID: replyId})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Reply detail retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.ReplyInsertResponse{
			Reply: ID,
		},
	})
	logger.Log.Info("Reply detail retreived successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *ReplyHandler) UpdateReply(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	replyId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	var res = dto.ReplyRequest{}

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

	err = h.Service.UpdateReply(res, replyId)
	if err != nil {
		logger.Log.Error("Failed to update reply detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Update Reply Record", ID: ID}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Reply Updated successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Updated Record", ID: ID,
		}})
	logger.Log.Info("Reply Updated successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *ReplyHandler) DeleteReply(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	ReplyId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)

	role := claims["role"].(string)

	UserId, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.DeleteReply(ReplyId, UserId, role)
	if err != nil {
		logger.Log.Error("Failed to Delete reply detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Deleted Reply Record", ID: UserId}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Reply Deleted successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Deleted Record", ID: UserId,
		},
	})
	logger.Log.Info("Reply Updated successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}
