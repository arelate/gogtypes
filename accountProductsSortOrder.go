package gog_types

type AccountProductsSortOrder string

const (
	AccountProductsSortTitle        AccountProductsSortOrder = "title"
	AccountProductsSortPurchaseDate                          = "date_purchased"
	AccountProductsSortTagsOrder                             = "tags"
	AccountProductsSortOldestFirst                           = "release_date_asc"
	AccountProductsSortNewestFirst                           = "release_date_asc"
)
