package usecase

import "github.com/labstack/echo/v4"

type Usecase interface {
	GetCategoryTree(ctx echo.Context) ([]interface{}, error)
	GetCategoryDetails(ctx echo.Context, id int) ([]interface{}, error)
	InsertCategory(ctx echo.Context, title string, slug string) ([]interface{}, error)
	UpdateCategory(ctx echo.Context, id int, title string, slug string) ([]interface{}, error)
	DeleteCategory(ctx echo.Context, id int) error
}

type handler struct {
	usecase Usecase
}

func New(usecase Usecase) Usecase {
	return &handler{
		usecase: usecase,
	}
}
