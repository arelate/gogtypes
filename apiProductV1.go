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
		En string `json:"en"`
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
	Dlcs      []interface{} `json:"dlcs"`
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
	ExpandedDlcs []interface{} `json:"expanded_dlcs"`
	Description  struct {
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
	RelatedProducts []interface{} `json:"related_products"`
	Changelog       string        `json:"changelog"`
}
