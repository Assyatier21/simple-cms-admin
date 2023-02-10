package models

import "errors"

var (
	ERROR_FORMAT_ID         = "id must be an integer"
	ERROR_FORMAT_LIMIT      = "limit must be an integer"
	ERROR_FORMAT_OFFSET     = "offset must be an integer"
	ERROR_FORMAT_SLUG       = "incorrect slug format"
	ERROR_FORMAT_CATEGORYID = "categoryid must be an integer"
	ERROR_FORMAT_METADATA   = "incorrect metadata format"

	ERROR_EMPTY_ID          = "id can't be empty"
	ERROR_EMPTY_LIMIT       = "limit can't be empty"
	ERROR_EMPTY_OFFSET      = "offset can't be empty"
	ERROR_EMPTY_TITLE       = "title can't be empty"
	ERROR_EMPTY_SLUG        = "slug can't be empty"
	ERROR_EMPTY_HTMLCONTENT = "htmlcontent can't be empty"
	ERROR_EMPTY_CATEGORYID  = "categoryid can't be empty"
	ERROR_EMPTY_METADATA    = "metadata can't be empty"
	ERROR_BAD_REQUEST       = "bad request"

	ERROR_FORMAT_EMPTY_ID         = "id must be an integer and can't be empty"
	ERROR_FORMAT_EMPTY_SLUG       = "incorrect slug format or slug can't be empty"
	ERROR_FORMAT_EMPTY_CATEGORYID = "incorrect categoryid format or categoryid can't be empty"
	ERROR_FORMAT_EMPTY_METADATA   = "incorrect metadata format or metadata can't be empty"

	ERROR_NOT_FOUND        = errors.New("data not found")
	ERROR_NO_ROWS_AFFECTED = errors.New("no rows affected")
	ERROR_NO_ROWS_RESULT   = errors.New("no rows in result set")
)
