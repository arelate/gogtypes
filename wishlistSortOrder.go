package gogtypes

type WishlistSortOrder string

const (
	WishlistSortByTitle       WishlistSortOrder = "title"
	WishlistSortByDateAdded                     = "date_added"
	WishlistSortByUserReviews                   = "user_reviews"
)
