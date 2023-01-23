package utils

import (
	"time"
)

var (
	STATUS_SUCCESS = "success"
	STATUS_FAILED  = "failed"

	PATH_ARTICLES   = "/articles"
	PATH_ARTICLE    = "/article"
	PATH_CATEGORIES = "/categories"
	PATH_CATEGORY   = "/category"

	jakartaLoc, _ = time.LoadLocation("Asia/Jakarta")
	TimeNow       = time.Now().In(jakartaLoc).Format("2006-01-02T15:04:05Z")
)
