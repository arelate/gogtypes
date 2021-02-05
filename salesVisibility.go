package gogtypes

type SalesVisibility struct {
	IsActive   bool `json:"isActive"`
	FromObject Date `json:"fromObject"`
	From       int  `json:"from"`
	ToObject   Date `json:"toObject"`
	To         int  `json:"to"`
}
