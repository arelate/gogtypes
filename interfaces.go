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

type DeveloperGetter interface {
	GetDeveloper() string
}

type PublisherGetter interface {
	GetPublisher() string
}
