package models

import "errors"

var (
	ERROR_FORMAT_ID         = "id must be an integer"
	ERROR_FORMAT_LIMIT      = "limit must be an integer"
	ERROR_FORMAT_OFFSET     = "offset must be an integer"
	ERROR_FORMAT_ORDER_BY   = "order by must be a boolean"
	ERROR_FORMAT_SLUG       = "incorrect slug format"
	ERROR_FORMAT_CATEGORYID = "categoryid must be an integer"
	ERROR_FORMAT_METADATA   = "incorrect metadata format"

	ERROR_EMPTY_ID          = "id failed to be empty"
	ERROR_EMPTY_LIMIT       = "limit failed to be empty"
	ERROR_EMPTY_OFFSET      = "offset failed to be empty"
	ERROR_EMPTY_TITLE       = "title failed to be empty"
	ERROR_EMPTY_SLUG        = "slug failed to be empty"
	ERROR_EMPTY_HTMLCONTENT = "htmlcontent failed to be empty"
	ERROR_EMPTY_CATEGORYID  = "categoryid failed to be empty"
	ERROR_EMPTY_METADATA    = "metadata failed to be empty"
	ERROR_BAD_REQUEST       = "bad request"

	ERROR_FORMAT_EMPTY_ID         = "id must be an integer and failed to be empty"
	ERROR_FORMAT_EMPTY_SLUG       = "incorrect slug format or slug failed to be empty"
	ERROR_FORMAT_EMPTY_CATEGORYID = "incorrect categoryid format or categoryid failed to be empty"
	ERROR_FORMAT_EMPTY_METADATA   = "incorrect metadata format or metadata failed to be empty"

	ERROR_NOT_FOUND        = errors.New("data not found")
	ERROR_NO_ROWS_AFFECTED = errors.New("no rows affected")
	ERROR_NO_ROWS_RESULT   = errors.New("no rows in result set")
)
