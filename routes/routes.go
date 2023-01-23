package routes

import (
	"cms-admin/internal/delivery/api"
	"cms-admin/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes(handler api.DeliveryHandler) *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	g := e.Group("/admin/v1")
	g.GET(utils.PATH_CATEGORIES, handler.GetCategoryTree)
	g.GET(utils.PATH_CATEGORY, handler.GetCategoryDetails)
	g.POST(utils.PATH_CATEGORY, handler.InsertCategory)
	g.PATCH(utils.PATH_CATEGORY, handler.UpdateCategory)
	g.DELETE(utils.PATH_CATEGORY, handler.DeleteCategory)

	return e
}
func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))
}
