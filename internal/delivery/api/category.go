package api

import (
	m "cms-admin/models"
	msg "cms-admin/models/lib"
	"cms-admin/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) GetCategoryTree(ctx echo.Context) (err error) {
	categories, err := h.usecase.GetCategoryTree(ctx.Request().Context())
	if err != nil {
		log.Println("[Delivery][GetCategoryTree] can't get list of categories, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "list of categories returned successfully", categories)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) GetCategoryDetails(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	category, err := h.usecase.GetCategoryDetails(ctx.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "no category found", []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		}
		log.Println("[Delivery][GetCategoryDetails] can't get category details, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "category details returned successfully", category)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) InsertCategory(ctx echo.Context) (err error) {
	var (
		title string
		slug  string
	)

	title = ctx.FormValue("title")
	if title == "" {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_EMPTY_TITLE)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	slug = ctx.FormValue("slug")
	if slug == "" || !utils.IsValidSlug(slug) {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_SLUG)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	category, err := h.usecase.InsertCategory(ctx.Request().Context(), title, slug)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				res := m.SetError(http.StatusConflict, utils.STATUS_FAILED, "slug has been used in another category")
				return ctx.JSON(http.StatusOK, res)
			}
		}
		log.Println("[Delivery][InsertCategory] can't insert category, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusCreated, utils.STATUS_SUCCESS, "category created successfully", category)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) UpdateCategory(ctx echo.Context) (err error) {
	var (
		id    int
		title string
		slug  string
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	title = ctx.FormValue("title")

	slug = ctx.FormValue("slug")
	if slug != "" && !utils.IsValidSlug(slug) {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_SLUG)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	category, err := h.usecase.UpdateCategory(ctx.Request().Context(), id, title, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "no category updated", []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		}
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				res := m.SetError(http.StatusConflict, utils.STATUS_FAILED, "slug has been used in another category")
				return ctx.JSON(http.StatusOK, res)
			}
		}
		log.Println("[Delivery][UpdateCategory] can't update category, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "category updated successfully", category)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) DeleteCategory(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	err = h.usecase.DeleteCategory(ctx.Request().Context(), id)
	if err != nil {
		if err == msg.ERROR_NO_ROWS_AFFECTED {
			res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "no category deleted", nil)
			return ctx.JSON(http.StatusOK, res)
		}
		log.Println("[Delivery][DeleteCategory] can't delete category, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)

	}

	res := m.SetResponse(http.StatusNoContent, utils.STATUS_SUCCESS, "category deleted successfully", []interface{}{})
	return ctx.JSON(http.StatusOK, res)
}
