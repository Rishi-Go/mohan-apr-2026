package handler

import (
	"blog_post/internals/dto"
	"blog_post/internals/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofrs/uuid"
)

type BlogHandler struct {
	Service service.BlogService
}

func InitBlogHandler(svc service.BlogService) *BlogHandler {
	return &BlogHandler{Service: svc}
}

func (h *BlogHandler) InsertBlog(Ctx fiber.Ctx) error {

	var res = dto.BlogRequest{}

	if err := Ctx.Bind().Body(&res); err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err := h.Service.InsertBlog(res)
	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = Ctx.JSON(dto.ResponseMessage{Message: "INSERTED SUCCESSFULLY"})
	if err != nil {

		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *BlogHandler) GetBlog(Ctx fiber.Ctx) error {

	title := Ctx.Query("title")

	categoryStr := Ctx.Query("category-id")

	categoryId := uuid.FromStringOrNil(categoryStr)

	AuthorStr := Ctx.Query("author-id")

	authorId := uuid.FromStringOrNil(AuthorStr)

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

	result, Page, err := h.Service.GetBlog(page, limit, offset, title, categoryId, authorId)
	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = Ctx.JSON(&dto.BlogResponse{
		Blog:       result,
		Pagination: *Page,
	})

	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	return nil

}

func (h *BlogHandler) SelectBlog(Ctx fiber.Ctx) error {
	
	uuidStr := Ctx.Params("id")

	BlogId, err := uuid.FromString(uuidStr)

	if err != nil {

		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	ID, err := h.Service.SelectBlog(BlogId)
	if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = Ctx.JSON(ID)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *BlogHandler) UpdateBlog(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	BlogId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	var res = dto.BlogRequest{}

	if err := Ctx.Bind().Body(&res); err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.UpdateBlog(res, BlogId)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = Ctx.JSON(dto.Response{Message: "UPDATED SUCCESSFULLY", ID: BlogId})
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *BlogHandler) DeleteBlog(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	BlogId, err := uuid.FromString(uuidStr)
	if err != nil {
		return  Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	err = h.Service.DeleteBlog(BlogId)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	err = Ctx.JSON(dto.Response{Message: "DELETED SUCCESSFULLY", ID: BlogId })
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}