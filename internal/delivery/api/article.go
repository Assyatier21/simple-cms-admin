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

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) GetArticles(ctx echo.Context) (err error) {
	var (
		limit    int
		offset   int
		sort_by  string
		order_by string
	)

	limit = 100
	if ctx.FormValue("limit") != "" {
		limit, err = strconv.Atoi(ctx.FormValue("limit"))
		if err != nil {
			return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_ID, nil)
		}
	}

	offset = 0
	if ctx.FormValue("offset") != "" {
		offset, err = strconv.Atoi(ctx.FormValue("offset"))
		if err != nil {
			return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_OFFSET, nil)
		}
	}

	sort_by = ctx.FormValue("sort_by")
	order_by = ctx.FormValue("order_by")

	articles, err := h.usecase.GetArticles(ctx.Request().Context(), limit, offset, sort_by, order_by)
	if err != nil {
		log.Println("[Delivery][GetArticles] failed to get list of articles, err: ", err)
		return helper.WriteResponse(ctx, http.StatusInternalServerError, utils.STATUS_FAILED, err.Error(), nil)

	}
	return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "list of articles returned successfully", articles)
}
func (h *handler) GetArticleDetails(ctx echo.Context) (err error) {
	var (
		article []interface{}
		id      string
	)

	id = ctx.FormValue("id")
	article, err = h.usecase.GetArticleDetails(ctx.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "no article found", nil)
		}
		log.Println("[Delivery][GetArticleDetails] failed to get article details, err: ", err)
		return helper.WriteResponse(ctx, http.StatusInternalServerError, utils.STATUS_FAILED, err.Error(), nil)
	}
	return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "article details returned successfully", article)
}
func (h *handler) InsertArticle(ctx echo.Context) (err error) {
	var (
		title       string
		slug        string
		htmlcontent string
		categoryid  int
		metadata    string
	)

	title = ctx.FormValue("title")
	if title == "" {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_EMPTY_TITLE, nil)
	}

	slug = ctx.FormValue("slug")
	if slug == "" || !utils.IsValidSlug(slug) {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_SLUG, nil)
	}

	htmlcontent = ctx.FormValue("html_content")
	if htmlcontent == "" {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_EMPTY_HTMLCONTENT, nil)
	}

	categoryid, err = strconv.Atoi(ctx.FormValue("category_id"))
	if err != nil {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_CATEGORYID, nil)
	}

	metadata = ctx.FormValue("metadata")
	if metadata == "" {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_EMPTY_METADATA)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	article, err := h.usecase.InsertArticle(ctx.Request().Context(), title, slug, htmlcontent, categoryid, metadata)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return helper.WriteResponse(ctx, http.StatusConflict, utils.STATUS_FAILED, "slug has been used in another article", nil)
			}
		}
		log.Println("[Delivery][InsertArticle] failed to insert article, err: ", err)
		return helper.WriteResponse(ctx, http.StatusInternalServerError, utils.STATUS_FAILED, err.Error(), nil)
	}
	return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "article created successfully", article)
}
func (h *handler) UpdateArticle(ctx echo.Context) (err error) {
	var (
		id          string
		title       string
		slug        string
		htmlcontent string
		categoryid  int
		metadata    string
	)

	id = ctx.FormValue("id")
	title = ctx.FormValue("title")

	slug = ctx.FormValue("slug")
	if slug != "" && !utils.IsValidSlug(slug) {
		return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_SLUG, nil)
	}

	htmlcontent = ctx.FormValue("html_content")

	categoryid = 0
	if ctx.FormValue("category_id") != "" {
		categoryid, err = strconv.Atoi(ctx.FormValue("category_id"))
		if err != nil {
			return helper.WriteResponse(ctx, http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_CATEGORYID, nil)
		}
	}

	metadata = ctx.FormValue("metadata")

	article, err := h.usecase.UpdateArticle(ctx.Request().Context(), id, title, slug, htmlcontent, categoryid, metadata)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "no article updated", nil)
		}
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return helper.WriteResponse(ctx, http.StatusConflict, utils.STATUS_FAILED, "slug has been used in another article", nil)
			}
		}
		log.Println("[Delivery][UpdateArticle] failed to update article, err: ", err)
		return helper.WriteResponse(ctx, http.StatusInternalServerError, utils.STATUS_FAILED, err.Error(), nil)
	}
	return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "article updated successfully", article)
}
func (h *handler) DeleteArticle(ctx echo.Context) (err error) {
	id := ctx.FormValue("id")

	err = h.usecase.DeleteArticle(ctx.Request().Context(), id)
	if err != nil {
		if err == msg.ERROR_NO_ROWS_AFFECTED {
			return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "no article deleted", nil)
		}
		log.Println("[Delivery][DeleteArticle] failed to delete article, err: ", err)
		return helper.WriteResponse(ctx, http.StatusInternalServerError, utils.STATUS_FAILED, err.Error(), nil)
	}
	return helper.WriteResponse(ctx, http.StatusOK, utils.STATUS_SUCCESS, "article deleted successfully", nil)
}
