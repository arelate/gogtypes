package gog_types

type IdGetter interface {
	GetId() int
}

type ProductsGetter interface {
	GetProducts() []IdGetter
}

type TitleGetter interface {
	GetTitle() string
}
