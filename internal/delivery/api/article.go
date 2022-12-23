package api

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetArticles(ctx echo.Context) (err error) {
	var (
		limit  int
		offset int
	)

	if ctx.FormValue("limit") == "" {
		limit = 100
	} else {
		limit, err = strconv.Atoi(ctx.FormValue("limit"))
		if err != nil {
			res := m.SetError(http.StatusBadRequest, "limit must be an integer")
			return ctx.JSON(http.StatusBadRequest, res)
		}
	}

	if ctx.FormValue("offset") == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(ctx.FormValue("offset"))
		if err != nil {
			res := m.SetError(http.StatusBadRequest, "offset must be an integer")
			return ctx.JSON(http.StatusBadRequest, res)
		}
	}

	datas, err := h.repository.GetArticles(ctx.Request().Context(), limit, offset)
	if err != nil {
		if err == utils.ErrNotFound {
			res := m.SetResponse(http.StatusOK, utils.ErrNotFound.Error(), []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		} else {
			log.Println("[Delivery][GetArticles] can't get list of articles, err:", err.Error())
			res := m.SetError(http.StatusInternalServerError, "failed to get list of articles")
			return ctx.JSON(http.StatusInternalServerError, res)
		}
	}

	articles := make([]interface{}, len(datas))
	for i, v := range datas {
		articles[i] = v
	}
	res := m.SetResponse(http.StatusOK, "success", articles)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) GetArticleDetails(ctx echo.Context) (err error) {
	var (
		id int
	)

	if !utils.IsValidNumeric(ctx.FormValue("id")) {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
	} else {
		id, _ = strconv.Atoi(ctx.FormValue("id"))
	}

	article, err := h.repository.GetArticleDetails(ctx.Request().Context(), id)
	if err != nil {
		if err == utils.ErrNotFound {
			res := m.SetResponse(http.StatusOK, utils.ErrNotFound.Error(), []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		} else {
			log.Println("[Delivery][GetArticleDetails] can't get article details, err:", err.Error())
			res := m.SetError(http.StatusInternalServerError, "failed to get article details")
			return ctx.JSON(http.StatusInternalServerError, res)
		}
	}

	var data []interface{}
	data = append(data, article)

	res := m.SetResponse(http.StatusOK, "success", data)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) InsertArticle(ctx echo.Context) (err error) {
	return
}
func (h *handler) UpdateArticle(ctx echo.Context) (err error) {
	return
}
func (h *handler) DeleteArticle(ctx echo.Context) (err error) {
	return
}
