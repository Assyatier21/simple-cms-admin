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

func (h *handler) GetArticles(ctx echo.Context) (err error) {
	var (
		limit  int
		offset int
	)

	limit = 100
	if ctx.FormValue("limit") != "" {
		limit, err = strconv.Atoi(ctx.FormValue("limit"))
		if err != nil {
			res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_ID)
			return ctx.JSON(http.StatusBadRequest, res)
		}
	}

	offset = 0
	if ctx.FormValue("offset") != "" {
		offset, err = strconv.Atoi(ctx.FormValue("offset"))
		if err != nil {
			res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_OFFSET)
			return ctx.JSON(http.StatusBadRequest, res)
		}
	}

	articles, err := h.usecase.GetArticles(ctx, limit, offset)
	if err != nil {
		log.Println("[Delivery][GetArticles] can't get list of articles, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "list of articles returned successfully", articles)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) GetArticleDetails(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	article, err := h.usecase.GetArticleDetails(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "no article found", []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		}
		log.Println("[Delivery][GetArticleDetails] can't get article details, err:", err.Error())
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "article details returned successfully", article)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) InsertArticle(ctx echo.Context) (err error) {
	var (
		title        string
		slug         string
		html_content string
		category_id  int
		metadata     string
	)

	title = ctx.FormValue("title")
	if title == "" {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_EMPTY_TITLE)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	slug = ctx.FormValue("slug")
	if slug == "" || !utils.IsValidSlug(slug) {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_EMPTY_SLUG)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	html_content = ctx.FormValue("html_content")
	if html_content == "" {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_EMPTY_HTML_CONTENT)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	category_id, err = strconv.Atoi(ctx.FormValue("category_id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_CATEGORY_ID)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	metadata = ctx.FormValue("metadata")
	if metadata == "" {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_EMPTY_METADATA)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	article, err := h.usecase.InsertArticle(ctx, title, slug, html_content, category_id, metadata)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				res := m.SetError(http.StatusConflict, utils.STATUS_FAILED, pqErr.Error())
				return ctx.JSON(http.StatusOK, res)
			}
		}
		log.Println("[Delivery][InsertArticle] can't insert article, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusCreated, utils.STATUS_SUCCESS, "article created successfully", article)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) UpdateArticle(ctx echo.Context) (err error) {
	var (
		id           int
		title        string
		slug         string
		html_content string
		category_id  int
		metadata     string
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

	html_content = ctx.FormValue("html_content")

	category_id = 0
	if ctx.FormValue("category_id") != "" {
		category_id, err = strconv.Atoi(ctx.FormValue("category_id"))
		if err != nil {
			res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_CATEGORY_ID)
			return ctx.JSON(http.StatusBadRequest, res)
		}
	}

	metadata = ctx.FormValue("metadata")

	article, err := h.usecase.UpdateArticle(ctx, id, title, slug, html_content, category_id, metadata)
	if err != nil {
		if err == sql.ErrNoRows {
			res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "no article updated", []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		}
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				res := m.SetError(http.StatusConflict, utils.STATUS_FAILED, "slug has been used in another article")
				return ctx.JSON(http.StatusOK, res)
			}
		}
		log.Println("[Delivery][UpdateArticle] can't update article, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "article updated successfully", article)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) DeleteArticle(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	err = h.usecase.DeleteArticle(ctx, id)
	if err != nil {
		if err == msg.ERROR_NO_ROWS_AFFECTED {
			res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "no article deleted", nil)
			return ctx.JSON(http.StatusOK, res)
		}
		log.Println("[Delivery][DeleteArticle] can't delete article, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)

	}

	res := m.SetResponse(http.StatusNoContent, utils.STATUS_SUCCESS, "article deleted successfully", []interface{}{})
	return ctx.JSON(http.StatusOK, res)
}
