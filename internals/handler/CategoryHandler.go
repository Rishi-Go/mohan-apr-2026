package handler

import (
	"blog_post/internals/service"

	"github.com/gofiber/fiber"
)

type CategoryHandler struct {
	Service service.CategoryService
}

func (h CategoryHandler) InsertCategory(c *fiber.Ctx) error {

	// c.Writer.Header().Add("Content-Type", "application/json")

	// var res = dto.CategoryRequest{}

	// if err := c.ShouldBind(&res); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
	// 	return
	// }

	// err := h.Service.InsertCategory(res)
	// if err != nil {
	// 	c.Writer.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(c.Writer).Encode(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	// 	return
	// }

	// err = json.NewEncoder(c.Writer).Encode(dto.ResponseMessage{Message: "INSERTED SUCCESSFULLY"})
	// if err != nil {
	// 	c.Writer.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(c.Writer).Encode(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	// 	return
	// }
	return nil

}
