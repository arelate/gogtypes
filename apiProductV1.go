package gog_types

type ApiProductV1 struct {
	Id                         int    `json:"id"`
	Title                      string `json:"title"`
	PurchaseLink               string `json:"purchase_link"`
	Slug                       string `json:"slug"`
	ContentSystemCompatibility struct {
		Windows bool `json:"windows"`
		Osx     bool `json:"osx"`
		Linux   bool `json:"linux"`
	} `json:"content_system_compatibility"`
	Languages struct {
		Br string `json:"br"`
		Cn string `json:"cn"`
		Cz string `json:"cz"`
		De string `json:"de"`
		En string `json:"en"`
		Es string `json:"es"`
		Fr string `json:"fr"`
		It string `json:"it"`
		Jp string `json:"jp"`
		Pl string `json:"pl"`
		Ru string `json:"ru"`
	} `json:"languages"`
	Links struct {
		PurchaseLink string `json:"purchase_link"`
		ProductCard  string `json:"product_card"`
		Support      string `json:"support"`
		Forum        string `json:"forum"`
	} `json:"links"`
	InDevelopment struct {
		Active bool        `json:"active"`
		Until  interface{} `json:"until"`
	} `json:"in_development"`
	IsSecret      bool   `json:"is_secret"`
	IsInstallable bool   `json:"is_installable"`
	GameType      string `json:"game_type"`
	IsPreOrder    bool   `json:"is_pre_order"`
	ReleaseDate   string `json:"release_date"`
	Images        struct {
		Background          string `json:"background"`
		Logo                string `json:"logo"`
		Logo2X              string `json:"logo2x"`
		Icon                string `json:"icon"`
		SidebarIcon         string `json:"sidebarIcon"`
		SidebarIcon2X       string `json:"sidebarIcon2x"`
		MenuNotificationAv  string `json:"menuNotificationAv"`
		MenuNotificationAv2 string `json:"menuNotificationAv2"`
	} `json:"images"`
	Dlcs struct {
		Products []struct {
			Id           int    `json:"id"`
			Link         string `json:"link"`
			ExpandedLink string `json:"expanded_link"`
		} `json:"products"`
		AllProductsUrl         string `json:"all_products_url"`
		ExpandedAllProductsUrl string `json:"expanded_all_products_url"`
	} `json:"dlcs"`
	Downloads struct {
		Installers []struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			Os           string `json:"os"`
			Language     string `json:"language"`
			LanguageFull string `json:"language_full"`
			Version      string `json:"version"`
			TotalSize    int64  `json:"total_size"`
			Files        []struct {
				Id       string `json:"id"`
				Size     int64  `json:"size"`
				Downlink string `json:"downlink"`
			} `json:"files"`
		} `json:"installers"`
		Patches []struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			Os           string `json:"os"`
			Language     string `json:"language"`
			LanguageFull string `json:"language_full"`
			Version      string `json:"version"`
			TotalSize    int64  `json:"total_size"`
			Files        []struct {
				Id       string `json:"id"`
				Size     int64  `json:"size"`
				Downlink string `json:"downlink"`
			} `json:"files"`
		} `json:"patches"`
		LanguagePacks []interface{} `json:"language_packs"`
		BonusContent  []interface{} `json:"bonus_content"`
	} `json:"downloads"`
	ExpandedDlcs []struct {
		Id                         int    `json:"id"`
		Title                      string `json:"title"`
		PurchaseLink               string `json:"purchase_link"`
		Slug                       string `json:"slug"`
		ContentSystemCompatibility struct {
			Windows bool `json:"windows"`
			Osx     bool `json:"osx"`
			Linux   bool `json:"linux"`
		} `json:"content_system_compatibility"`
		Languages struct {
			En string `json:"en"`
			Br string `json:"br,omitempty"`
			Cn string `json:"cn,omitempty"`
			Cz string `json:"cz,omitempty"`
			De string `json:"de,omitempty"`
			Es string `json:"es,omitempty"`
			Fr string `json:"fr,omitempty"`
			Gk string `json:"gk,omitempty"`
			It string `json:"it,omitempty"`
			Jp string `json:"jp,omitempty"`
			Nl string `json:"nl,omitempty"`
			Ru string `json:"ru,omitempty"`
			Pl string `json:"pl,omitempty"`
		} `json:"languages"`
		Links struct {
			PurchaseLink string `json:"purchase_link"`
			ProductCard  string `json:"product_card"`
			Support      string `json:"support"`
			Forum        string `json:"forum"`
		} `json:"links"`
		InDevelopment struct {
			Active bool        `json:"active"`
			Until  interface{} `json:"until"`
		} `json:"in_development"`
		IsSecret      bool    `json:"is_secret"`
		IsInstallable bool    `json:"is_installable"`
		GameType      string  `json:"game_type"`
		IsPreOrder    bool    `json:"is_pre_order"`
		ReleaseDate   *string `json:"release_date"`
		Images        struct {
			Background          *string `json:"background"`
			Logo                *string `json:"logo"`
			Logo2X              *string `json:"logo2x"`
			Icon                string  `json:"icon"`
			SidebarIcon         *string `json:"sidebarIcon"`
			SidebarIcon2X       *string `json:"sidebarIcon2x"`
			MenuNotificationAv  *string `json:"menuNotificationAv"`
			MenuNotificationAv2 *string `json:"menuNotificationAv2"`
		} `json:"images"`
		Downloads struct {
			Installers []struct {
				Id           string `json:"id"`
				Name         string `json:"name"`
				Os           string `json:"os"`
				Language     string `json:"language"`
				LanguageFull string `json:"language_full"`
				Version      string `json:"version"`
				TotalSize    int64  `json:"total_size"`
				Files        []struct {
					Id       string `json:"id"`
					Size     int64  `json:"size"`
					Downlink string `json:"downlink"`
				} `json:"files"`
			} `json:"installers"`
			Patches []struct {
				Id           string `json:"id"`
				Name         string `json:"name"`
				Os           string `json:"os"`
				Language     string `json:"language"`
				LanguageFull string `json:"language_full"`
				Version      string `json:"version"`
				TotalSize    int    `json:"total_size"`
				Files        []struct {
					Id       string `json:"id"`
					Size     int    `json:"size"`
					Downlink string `json:"downlink"`
				} `json:"files"`
			} `json:"patches"`
			LanguagePacks []interface{} `json:"language_packs"`
			BonusContent  []interface{} `json:"bonus_content"`
		} `json:"downloads"`
	} `json:"expanded_dlcs"`
	Description struct {
		Lead             string `json:"lead"`
		Full             string `json:"full"`
		WhatsCoolAboutIt string `json:"whats_cool_about_it"`
	} `json:"description"`
	Screenshots []struct {
		ImageId              string `json:"image_id"`
		FormatterTemplateUrl string `json:"formatter_template_url"`
		FormattedImages      []struct {
			FormatterName string `json:"formatter_name"`
			ImageUrl      string `json:"image_url"`
		} `json:"formatted_images"`
	} `json:"screenshots"`
	Videos          []interface{} `json:"videos"`
	RelatedProducts []struct {
		Id                         int    `json:"id"`
		Title                      string `json:"title"`
		PurchaseLink               string `json:"purchase_link"`
		Slug                       string `json:"slug"`
		ContentSystemCompatibility struct {
			Windows bool `json:"windows"`
			Osx     bool `json:"osx"`
			Linux   bool `json:"linux"`
		} `json:"content_system_compatibility"`
		Languages struct {
			En string `json:"en"`
			Br string `json:"br,omitempty"`
			Cn string `json:"cn,omitempty"`
			Cz string `json:"cz,omitempty"`
			De string `json:"de,omitempty"`
			Es string `json:"es,omitempty"`
			Fr string `json:"fr,omitempty"`
			Gk string `json:"gk,omitempty"`
			It string `json:"it,omitempty"`
			Jp string `json:"jp,omitempty"`
			Nl string `json:"nl,omitempty"`
			Ru string `json:"ru,omitempty"`
			Pl string `json:"pl,omitempty"`
		} `json:"languages"`
		Links struct {
			PurchaseLink string `json:"purchase_link"`
			ProductCard  string `json:"product_card"`
			Support      string `json:"support"`
			Forum        string `json:"forum"`
		} `json:"links"`
		InDevelopment struct {
			Active bool        `json:"active"`
			Until  interface{} `json:"until"`
		} `json:"in_development"`
		IsSecret      bool   `json:"is_secret"`
		IsInstallable bool   `json:"is_installable"`
		GameType      string `json:"game_type"`
		IsPreOrder    bool   `json:"is_pre_order"`
		ReleaseDate   string `json:"release_date"`
		Images        struct {
			Background          string `json:"background"`
			Logo                string `json:"logo"`
			Logo2X              string `json:"logo2x"`
			Icon                string `json:"icon"`
			SidebarIcon         string `json:"sidebarIcon"`
			SidebarIcon2X       string `json:"sidebarIcon2x"`
			MenuNotificationAv  string `json:"menuNotificationAv"`
			MenuNotificationAv2 string `json:"menuNotificationAv2"`
		} `json:"images"`
	} `json:"related_products"`
	Changelog string `json:"changelog"`
}
