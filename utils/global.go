package utils

import (
	"errors"
	"time"
)

var (
	PathArticles   = "/articles"
	PathArticle    = "/article"
	PathCategories = "/categories"
	PathCategory   = "/category"
	ErrNotFound    = errors.New("data not found")
	NoRowsAffected = errors.New("no rows affected")
	jakartaLoc, _  = time.LoadLocation("Asia/Jakarta")
	TimeNow        = time.Now().In(jakartaLoc).Format("2006-01-02T15:04:05Z")
)
