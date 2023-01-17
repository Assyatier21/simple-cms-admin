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
	g.GET(utils.PathCategories, handler.GetCategoryTree)
	g.GET(utils.PathCategory, handler.GetCategoryDetails)
	g.POST(utils.PathCategory, handler.InsertCategory)
	g.PATCH(utils.PathCategory, handler.UpdateCategory)
	g.DELETE(utils.PathCategory, handler.DeleteCategory)

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
