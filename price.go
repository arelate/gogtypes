package gogtypes

type Price struct {
	Amount                     string  `json:"amount"`
	BaseAmount                 string  `json:"baseAmount"`
	FinalAmount                string  `json:"finalAmount"`
	IsDiscounted               bool    `json:"isDiscounted"`
	DiscountPercentage         int     `json:"discountPercentage"`
	DiscountDifference         string  `json:"discountDifference"`
	Symbol                     string  `json:"symbol"`
	IsFree                     bool    `json:"isFree"`
	Discount                   int     `json:"discount"`
	IsBonusStoreCreditIncluded bool    `json:"isBonusStoreCreditIncluded"`
	BonusStoreCreditAmount     string  `json:"bonusStoreCreditAmount"`
	PromoId                    *string `json:"promoId"`
}
