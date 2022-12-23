package api

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetCategoryTree(ctx echo.Context) (err error) {
	var (
		datas []m.Category
	)

	datas, err = h.repository.GetCategoryTree(ctx.Request().Context())
	if err != nil {
		if err == utils.ErrNotFound {
			res := m.SetResponse(http.StatusOK, utils.ErrNotFound.Error(), []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		} else {
			log.Println("[Delivery][GetCategoryTree] can't get list of categories, err:", err.Error())
			res := m.SetError(http.StatusInternalServerError, "failed to get list of categories")
			return ctx.JSON(http.StatusInternalServerError, res)
		}
	}

	categories := make([]interface{}, len(datas))
	for i, v := range datas {
		categories[i] = v
	}
	res := m.SetResponse(http.StatusOK, "success", categories)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) GetCategoryDetails(ctx echo.Context) (err error) {
	var (
		id int
	)

	if !utils.IsValidNumeric(ctx.FormValue("id")) {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
	} else {
		id, _ = strconv.Atoi(ctx.FormValue("id"))
	}

	category, err := h.repository.GetCategoryDetails(ctx.Request().Context(), id)
	if err != nil {
		if err == utils.ErrNotFound {
			res := m.SetResponse(http.StatusOK, utils.ErrNotFound.Error(), []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		} else {
			log.Println("[Delivery][GetCategoryDetails] can't get category details, err:", err.Error())
			res := m.SetError(http.StatusInternalServerError, "failed to get category details")
			return ctx.JSON(http.StatusInternalServerError, res)
		}
	}

	var data []interface{}
	data = append(data, category)

	res := m.SetResponse(http.StatusOK, "success", data)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) InsertCategory(ctx echo.Context) (err error) {
	return
}
func (h *handler) UpdateCategory(ctx echo.Context) (err error) {
	return
}
func (h *handler) DeleteCategory(ctx echo.Context) (err error) {
	return
}
