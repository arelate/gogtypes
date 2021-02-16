package gog_types

type ProductsSortOrder string

const (
	ProductsSortByBestselling        ProductsSortOrder = "popularity"
	ProductsSortByAlphabetically                       = "title"
	ProductsSortByUserRating                           = "rating"
	ProductsSortByDateAdded                            = "date"
	ProductsSortByBestsellingAllTime                   = "bestselling"
	ProductsSortByOldestFirst                          = "release_asc"
	ProductsSortByNewestFirst                          = "release_desc"
)
