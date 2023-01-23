package models

import "errors"

var (
	ERROR_FORMAT_ID          = "id must be an integer"
	ERROR_FORMAT_LIMIT       = "limit must be an integer"
	ERROR_FORMAT_OFFSET      = "offset must be an integer"
	ERROR_FORMAT_SLUG        = "incorrect slug format"
	ERROR_FORMAT_CATEGORY_ID = "category_id must be an integer"
	ERROR_FORMAT_METADATA    = "incorrect metadata format"

	ERROR_EMPTY_ID           = "id can't be empty"
	ERROR_EMPTY_LIMIT        = "limit can't be empty"
	ERROR_EMPTY_OFFSET       = "offset can't be empty"
	ERROR_EMPTY_TITLE        = "title can't be empty"
	ERROR_EMPTY_SLUG         = "slug can't be empty"
	ERROR_EMPTY_HTML_CONTENT = "html_content can't be empty"
	ERROR_EMPTY_CATEGORY_ID  = "category_id can't be empty"
	ERROR_EMPTY_METADATA     = "metadata can't be empty"

	ERROR_FORMAT_EMPTY_ID          = "id must be an integer and can't be empty"
	ERROR_FORMAT_EMPTY_SLUG        = "incorrect slug format or slug can't be empty"
	ERROR_FORMAT_EMPTY_CATEGORY_ID = "incorrect category_id format or category_id can't be empty"
	ERROR_FORMAT_EMPTY_METADATA    = "incorrect metadata format or metadata can't be empty"

	ERROR_NOT_FOUND        = errors.New("data not found")
	ERROR_NO_ROWS_AFFECTED = errors.New("no rows affected")
)
