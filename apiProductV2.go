package gog_types

import (
	"strings"
	"time"
)

const formatterTemplate = "_{formatter}"

type ApiProductV2 struct {
	InDevelopment struct {
		Active bool `json:"active"`
	} `json:"inDevelopment"`
	Copyrights    string `json:"copyrights"`
	IsUsingDosBox bool   `json:"isUsingDosBox"`
	Description   string `json:"description"`
	Size          int    `json:"size"`
	Overview      string `json:"overview"`
	Links         struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Store struct {
			Href string `json:"href"`
		} `json:"store"`
		Support struct {
			Href string `json:"href"`
		} `json:"support"`
		Forum struct {
			Href string `json:"href"`
		} `json:"forum"`
		Icon struct {
			Href string `json:"href"`
		} `json:"icon"`
		IconSquare struct {
			Href string `json:"href"`
		} `json:"iconSquare"`
		Logo struct {
			Href string `json:"href"`
		} `json:"logo"`
		BoxArtImage struct {
			Href string `json:"href"`
		} `json:"boxArtImage"`
		BackgroundImage struct {
			Href string `json:"href"`
		} `json:"backgroundImage"`
		GalaxyBackgroundImage struct {
			Href string `json:"href"`
		} `json:"galaxyBackgroundImage"`
		IncludesGames []struct {
			Href string `json:"href"`
		} `json:"includesGames"`
	} `json:"_links"`
	Embedded struct {
		Product struct {
			IsAvailableForSale bool      `json:"isAvailableForSale"`
			IsVisibleInCatalog bool      `json:"isVisibleInCatalog"`
			IsPreorder         bool      `json:"isPreorder"`
			IsVisibleInAccount bool      `json:"isVisibleInAccount"`
			Id                 int       `json:"id"`
			Title              string    `json:"title"`
			IsInstallable      bool      `json:"isInstallable"`
			GlobalReleaseDate  time.Time `json:"globalReleaseDate"`
			HasProductCard     bool      `json:"hasProductCard"`
			GogReleaseDate     time.Time `json:"gogReleaseDate"`
			IsSecret           bool      `json:"isSecret"`
			Links              struct {
				Image struct {
					Href       string   `json:"href"`
					Templated  bool     `json:"templated"`
					Formatters []string `json:"formatters"`
				} `json:"image"`
				Checkout struct {
					Href string `json:"href"`
				} `json:"checkout"`
				Prices struct {
					Href      string `json:"href"`
					Templated bool   `json:"templated"`
				} `json:"prices"`
			} `json:"_links"`
		} `json:"product"`
		ProductType   string `json:"productType"`
		Localizations []struct {
			Embedded struct {
				Language struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"language"`
				LocalizationScope struct {
					Type string `json:"type"`
				} `json:"localizationScope"`
			} `json:"_embedded"`
		} `json:"localizations"`
		Videos []struct {
			Provider    string `json:"provider"`
			VideoId     string `json:"videoId"`
			ThumbnailId string `json:"thumbnailId"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Thumbnail struct {
					Href string `json:"href"`
				} `json:"thumbnail"`
			} `json:"_links"`
		} `json:"videos"`
		// TODO: find examples where this is not nil to infer the type
		Bonuses     []interface{} `json:"bonuses"`
		Screenshots []struct {
			Links struct {
				Self struct {
					Href       string   `json:"href"`
					Templated  bool     `json:"templated"`
					Formatters []string `json:"formatters"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"screenshots"`
		Publisher struct {
			Name string `json:"name"`
		} `json:"publisher"`
		Developers []struct {
			Name string `json:"name"`
		} `json:"developers"`
		SupportedOperatingSystems []struct {
			OperatingSystem struct {
				Name     string `json:"name"`
				Versions string `json:"versions"`
			} `json:"operatingSystem"`
			SystemRequirements []struct {
				Type         string `json:"type"`
				Description  string `json:"description"`
				Requirements []struct {
					Id          string `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
				} `json:"requirements"`
			} `json:"systemRequirements"`
		} `json:"supportedOperatingSystems"`
		Tags []struct {
			Id    int    `json:"id"`
			Name  string `json:"name"`
			Level int    `json:"level"`
			Slug  string `json:"slug"`
		} `json:"tags"`
		// TODO: find examples where this is not nil to infer the type
		EsrbRating interface{} `json:"esrbRating"`
		PegiRating struct {
			AgeRating int `json:"ageRating"`
			// TODO: find examples where this is not nil to infer the type
			ContentDescriptors []interface{} `json:"contentDescriptors"`
		} `json:"pegiRating"`
		// TODO: find examples where this is not nil to infer the type
		UskRating interface{} `json:"uskRating"`
		// TODO: find examples where this is not nil to infer the type
		BrRating interface{} `json:"brRating"`
		Features []struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"features"`
		// TODO: find examples where this is not nil to infer the type
		Editions []interface{} `json:"editions"`
		// TODO: find examples where this is not nil to infer the type
		Series interface{} `json:"series"`
	} `json:"_embedded"`
}

func (apv2 *ApiProductV2) GetId() int {
	return apv2.Embedded.Product.Id
}

func (apv2 *ApiProductV2) GetTitle() string {
	return apv2.Embedded.Product.Title
}

func (apv2 *ApiProductV2) GetDeveloper() string {
	devs := make([]string, 0, len(apv2.Embedded.Developers))
	for _, dev := range apv2.Embedded.Developers {
		devs = append(devs, dev.Name)
	}
	return strings.Join(devs, ",")
}

func (apv2 *ApiProductV2) GetPublisher() string {
	return apv2.Embedded.Publisher.Name
}

func (apv2 *ApiProductV2) GetImage() string {
	return strings.Replace(apv2.Embedded.Product.Links.Image.Href, formatterTemplate, "", 1)
}
