package routes

import (
	"cms-admin/internal/delivery/api"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes(handler api.Handler) *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	g := e.Group("/admin/v1")
	g.GET("/articles", handler.GetArticles)
	g.GET("/article", handler.GetArticleDetails)
	// g.POST("/article", handler.InsertArticle)
	// g.PUT("/article", handler.UpdateArticle)
	// g.DELETE("/article", handler.DeleteArticle)

	g.GET("/categories", handler.GetCategoryTree)
	g.GET("/category", handler.GetCategoryByID)
	// g.POST("/category", handler.InsertCategory)
	// g.PUT("/category", handler.Updatecategory)
	// g.DELETE("/category", handler.DeleteCategory)

	return e
}
func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
	}))
}
