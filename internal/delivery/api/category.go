package api

import (
	m "cms-admin/models"
	msg "cms-admin/models/lib"
	"cms-admin/utils"
	"cms-admin/utils/helper"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

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
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID, nil)
	}

	category, err := h.usecase.GetCategoryDetails(ctx.Request().Context(), id)
	if err != nil {
		log.Println("[Delivery][GetCategoryDetails] can't get category details, err:", err.Error())
		if strings.Contains(err.Error(), "Bad Request") {
			return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_BAD_REQUEST, nil)
		}
		return helper.WriteResponse(ctx, http.StatusInternalServerError, utils.STATUS_FAILED, err.Error(), nil)
	}

	if category == nil {
		return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "no category found", nil)
	}
	return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "category details returned successfully", category)
}
func (h *handler) InsertCategory(ctx echo.Context) (err error) {
	var (
		title string
		slug  string
	)

	title = ctx.FormValue("title")
	if title == "" {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_EMPTY_TITLE, nil)
	}

	slug = ctx.FormValue("slug")
	if slug == "" || !utils.IsValidSlug(slug) {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_SLUG, nil)
	}

	category, err := h.usecase.InsertCategory(ctx.Request().Context(), title, slug)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return helper.WriteResponse(ctx, http.StatusConflict, utils.STATUS_FAILED, "slug has been used in another category", nil)
			}
		}
		log.Println("[Delivery][InsertCategory] can't insert category, err:", err.Error())
		return helper.WriteResponse(ctx, http.StatusInternalServerError, utils.STATUS_FAILED, err.Error(), nil)
	}
	return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "category created successfully", category)
}
func (h *handler) UpdateCategory(ctx echo.Context) (err error) {
	var (
		id    int
		title string
		slug  string
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID, nil)
	}

	title = ctx.FormValue("title")

	slug = ctx.FormValue("slug")
	if slug != "" && !utils.IsValidSlug(slug) {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_SLUG, nil)
	}

	category, err := h.usecase.UpdateCategory(ctx.Request().Context(), id, title, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "no category updated", nil)
		}
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return helper.WriteResponse(ctx, http.StatusConflict, utils.STATUS_FAILED, "slug has been used in another category", nil)
			}
		}
		log.Println("[Delivery][UpdateCategory] can't update category, err:", err.Error())
		return helper.WriteResponse(ctx, http.StatusInternalServerError, utils.STATUS_FAILED, err.Error(), nil)
	}
	return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "category updated successfully", category)
}
func (h *handler) DeleteCategory(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID, nil)
	}

	err = h.usecase.DeleteCategory(ctx.Request().Context(), id)
	if err != nil {
		if err == msg.ERROR_NO_ROWS_AFFECTED {
			return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "no category deleted", nil)
		}
		log.Println("[Delivery][DeleteCategory] can't delete category, err:", err.Error())
		return helper.WriteResponse(ctx, http.StatusInternalServerError, utils.STATUS_FAILED, err.Error(), nil)
	}
	return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "category deleted successfully", nil)
}
