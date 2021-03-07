package gog_types

type IdGetter interface {
	GetId() int
}

type TitleGetter interface {
	GetTitle() string
}
