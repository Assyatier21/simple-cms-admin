package api

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetCategoryTree(ctx echo.Context) (err error) {
	var (
		categories []interface{}
	)

	categories, err = h.usecase.GetCategoryTree(ctx)
	if err != nil {
		log.Println("[Delivery][GetCategoryTree] can't get list of categories, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, "failed to get list of categories")
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, "success", categories)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) GetCategoryDetails(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.ErrorFormatIDStr)
		return ctx.JSON(http.StatusBadRequest, res)
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
	var (
		returnCategory m.Category
	)

	if ctx.FormValue("title") == "" {
		res := m.SetError(http.StatusBadRequest, utils.ErrorTitleEmptyStr)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if ctx.FormValue("slug") == "" || !utils.IsValidSlug(ctx.FormValue("slug")) {
		res := m.SetError(http.StatusBadRequest, utils.ErrorFormatSlugStr)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	ctx.Bind(&returnCategory)
	utils.SetCategoryCreatedUpdatedTimeNow(&returnCategory)

	category, err := h.repository.InsertCategory(ctx.Request().Context(), returnCategory)
	if err != nil {
		log.Println("[Delivery][InsertCategory] can't insert category, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	data = append(data, category)

	res := m.SetResponse(http.StatusOK, "success", data)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) UpdateCategory(ctx echo.Context) (err error) {
	var (
		updatedCategory m.Category
	)

	updatedCategory.Id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.ErrorFormatIDStr)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if ctx.FormValue("title") == "" {
		updatedCategory.Title = ""
	} else {
		updatedCategory.Title = ctx.FormValue("title")
	}

	if ctx.FormValue("slug") == "" {
		updatedCategory.Slug = ""
	} else if !utils.IsValidSlug(ctx.FormValue("slug")) {
		res := m.SetError(http.StatusBadRequest, utils.ErrorFormatSlugStr)
		return ctx.JSON(http.StatusBadRequest, res)
	} else {
		updatedCategory.Slug = ctx.FormValue("slug")
	}

	ctx.Bind(&updatedCategory)
	utils.SetCategoryUpdatedTimeNow(&updatedCategory)

	fmt.Println(updatedCategory)

	category, err := h.repository.UpdateCategory(ctx.Request().Context(), updatedCategory)
	if err != nil {
		log.Println("[Delivery][UpdateCategory] can't update category, err:", err.Error())
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
	data = append(data, category)

	res := m.SetResponse(http.StatusOK, "success", data)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) DeleteCategory(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.ErrorFormatIDStr)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	err = h.repository.DeleteCategory(ctx.Request().Context(), id)
	if err != nil {
		log.Println("[Delivery][DeleteCategory] can't delete category, err:", err.Error())
		if err == utils.NoRowsAffected {
			res := m.SetError(http.StatusOK, utils.NoRowsAffected.Error())
			return ctx.JSON(http.StatusOK, res)
		} else {
			res := m.SetError(http.StatusInternalServerError, err.Error())
			return ctx.JSON(http.StatusInternalServerError, res)
		}
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "OK"})
}

// func (h *handler) GetCategoryTree(ctx echo.Context) (err error) {
// 	var (
// 		datas []m.Category
// 	)

// 	datas, err = h.repository.GetCategoryTree(ctx.Request().Context())
// 	if err != nil {
// 		if err == utils.ErrNotFound {
// 			res := m.SetResponse(http.StatusOK, utils.ErrNotFound.Error(), []interface{}{})
// 			return ctx.JSON(http.StatusOK, res)
// 		} else {
// 			log.Println("[Delivery][GetCategoryTree] can't get list of categories, err:", err.Error())
// 			res := m.SetError(http.StatusInternalServerError, "failed to get list of categories")
// 			return ctx.JSON(http.StatusInternalServerError, res)
// 		}
// 	}

// 	categories := make([]interface{}, len(datas))
// 	for i, v := range datas {
// 		categories[i] = v
// 	}
// 	res := m.SetResponse(http.StatusOK, "success", categories)
// 	return ctx.JSON(http.StatusOK, res)
// }
