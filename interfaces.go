package gog_types

type ProductsGetter interface {
	GetProducts() []IdGetter
}

type IdGetter interface {
	GetId() int
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

type ImageGetter interface {
	GetImage() string
}

type BoxArtGetter interface {
	GetBoxArt() string
}
