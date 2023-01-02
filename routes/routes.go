package routes

import (
	"cms-admin/internal/delivery/api"
	"cms-admin/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes(handler api.Handler) *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	g := e.Group("/admin/v1")
	g.GET(utils.PathArticles, handler.GetArticles)
	g.GET(utils.PathArticle, handler.GetArticleDetails)
	g.POST(utils.PathArticle, handler.InsertArticle)
	g.PATCH(utils.PathArticle, handler.UpdateArticle)
	g.DELETE(utils.PathArticle, handler.DeleteArticle)

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
