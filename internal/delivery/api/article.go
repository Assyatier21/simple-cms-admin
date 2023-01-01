package api

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"encoding/json"
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

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
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
	var (
		returnArticle m.Article
	)

	if ctx.FormValue("title") == "" {
		res := m.SetError(http.StatusBadRequest, "title can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if ctx.FormValue("slug") == "" || !utils.IsValidSlug(ctx.FormValue("slug")) {
		res := m.SetError(http.StatusBadRequest, "slug format is wrong")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if ctx.FormValue("html_content") == "" {
		res := m.SetError(http.StatusBadRequest, "html_content can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if ctx.FormValue("category_id") == "" || !utils.IsValidNumeric(ctx.FormValue("category_id")) {
		res := m.SetError(http.StatusBadRequest, "category_id must be an integer and can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if ctx.FormValue("metadata") == "" {
		res := m.SetError(http.StatusBadRequest, "metadata can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
	} else {
		metadataJSON := ctx.FormValue("metadata")
		err := json.Unmarshal([]byte(metadataJSON), &returnArticle.MetaData)
		if err != nil {
			res := m.SetError(http.StatusBadRequest, "error unmarshalling metadata")
			return ctx.JSON(http.StatusBadRequest, res)
		}
	}

	ctx.Bind(&returnArticle)
	utils.SetArticleCreatedUpdatedTimeNow(&returnArticle)

	article, err := h.repository.InsertArticle(ctx.Request().Context(), returnArticle)
	if err != nil {
		log.Println("[Delivery][InsertArticle] can't insert article, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	data = append(data, article)

	res := m.SetResponse(http.StatusOK, "success", data)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) UpdateArticle(ctx echo.Context) (err error) {
	var (
		updatedArticle m.Article
	)

	_, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if ctx.FormValue("title") == "" {
		updatedArticle.Title = ""
	}

	if ctx.FormValue("slug") == "" {
		updatedArticle.Slug = ""
	} else if !utils.IsValidSlug(ctx.FormValue("slug")) {
		res := m.SetError(http.StatusBadRequest, "slug format wrong")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if ctx.FormValue("html_content") == "" {
		updatedArticle.HtmlContent = ""
	}

	if ctx.FormValue("category_id") == "" {
		updatedArticle.CategoryID = 0
	} else if _, err = strconv.Atoi(ctx.FormValue("category_id")); err != nil {
		res := m.SetError(http.StatusBadRequest, "category_id must be an integer")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if ctx.FormValue("metadata") == "" {
		updatedArticle.MetaData = m.MetaData{}
	} else {
		metadataJSON := ctx.FormValue("metadata")
		err := json.Unmarshal([]byte(metadataJSON), &updatedArticle.MetaData)
		if err != nil {
			res := m.SetError(http.StatusBadRequest, "error unmarshalling metadata")
			return ctx.JSON(http.StatusBadRequest, res)
		}
	}

	ctx.Bind(&updatedArticle)
	utils.SetArticleUpdatedTimeNow(&updatedArticle)

	article, err := h.repository.UpdateArticle(ctx.Request().Context(), updatedArticle)
	if err != nil {
		log.Println("[Delivery][UpdateArticle] can't update article, err:", err.Error())
		if err == utils.NoRowsAffected {
			res := m.SetError(http.StatusOK, utils.NoRowsAffected.Error())
			return ctx.JSON(http.StatusOK, res)
		} else if err == utils.ErrNotFound {
			res := m.SetError(http.StatusNotFound, utils.ErrNotFound.Error())
			return ctx.JSON(http.StatusNotFound, res)
		} else {
			res := m.SetError(http.StatusInternalServerError, err.Error())
			return ctx.JSON(http.StatusInternalServerError, res)
		}
	}

	var data []interface{}
	data = append(data, article)

	res := m.SetResponse(http.StatusOK, "success", data)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) DeleteArticle(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	err = h.repository.DeleteArticle(ctx.Request().Context(), id)
	if err != nil {
		if err == utils.NoRowsAffected {
			log.Println("[Delivery][DeleteArticle] can't delete article, err:", err.Error())
			res := m.SetError(http.StatusOK, utils.NoRowsAffected.Error())
			return ctx.JSON(http.StatusOK, res)
		} else {
			log.Println("[Delivery][DeleteArticle] can't delete article, err:", err.Error())
			res := m.SetError(http.StatusInternalServerError, err.Error())
			return ctx.JSON(http.StatusInternalServerError, res)
		}
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
