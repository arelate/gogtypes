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
	// TODO: figure out a more sustainable plan for languages
	Languages struct {
		Cn string `json:"cn"`
		De string `json:"de"`
		En string `json:"en"`
		Es string `json:"es"`
		Jp string `json:"jp"`
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
	DLCs      interface{} `json:"dlcs"`
	Downloads struct {
		Installers []struct {
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
		} `json:"installers"`
		Patches       []interface{} `json:"patches"`
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
		Languages interface{} `json:"languages"`
		Links     struct {
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
		Downloads struct {
			Installers []struct {
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
			} `json:"installers"`
			Patches       []interface{} `json:"patches"`
			LanguagePacks []interface{} `json:"language_packs"`
			BonusContent  []struct {
				Id        int    `json:"id"`
				Name      string `json:"name"`
				Type      string `json:"type"`
				Count     int    `json:"count"`
				TotalSize int    `json:"total_size"`
				Files     []struct {
					Id       int    `json:"id"`
					Size     int    `json:"size"`
					Downlink string `json:"downlink"`
				} `json:"files"`
			} `json:"bonus_content"`
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
	Videos []struct {
		VideoUrl     string `json:"video_url"`
		ThumbnailUrl string `json:"thumbnail_url"`
		Provider     string `json:"provider"`
	} `json:"videos"`
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
			Cn string `json:"cn,omitempty"`
			De string `json:"de,omitempty"`
			Jp string `json:"jp,omitempty"`
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

func (apv1 *ApiProductV1) GetId() int {
	return apv1.Id
}

func (apv1 *ApiProductV1) GetTitle() string {
	return apv1.Title
}

func (apv1 *ApiProductV1) GetIcon() string {
	return apv1.Images.Icon
}

// ApiProducts doesn't contain logo.
// images.logo, images.logo2x are basically specially formatted DownloadType.Image
// func (apv1 *ApiProductV1) GetLogo() string { return "" }

func (apv1 *ApiProductV1) GetBackgroundImage() string {
	return apv1.Images.Background
}

func (apv1 *ApiProductV1) GetScreenshots() []string {
	screenshots := make([]string, 0)
	for _, screenshot := range apv1.Screenshots {
		screenshots = append(screenshots, screenshot.FormatterTemplateUrl)
	}
	return screenshots
}
