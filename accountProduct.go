// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_types

type AccountProduct struct {
	IsGalaxyCompatible bool     `json:"isGalaxyCompatible"`
	Tags               []string `json:"tags"`
	Id                 int      `json:"id"`
	Availability       struct {
		IsAvailable          bool `json:"isAvailable"`
		IsAvailableInAccount bool `json:"isAvailableInAccount"`
	} `json:"availability"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Url     string `json:"url"`
	WorksOn struct {
		Windows bool `json:"Windows"`
		Mac     bool `json:"Mac"`
		Linux   bool `json:"Linux"`
	} `json:"worksOn"`
	Category     string `json:"category"`
	Rating       int    `json:"rating"`
	IsComingSoon bool   `json:"isComingSoon"`
	IsMovie      bool   `json:"isMovie"`
	IsGame       bool   `json:"isGame"`
	Slug         string `json:"slug"`
	Updates      int    `json:"updates"`
	IsNew        bool   `json:"isNew"`
	DlcCount     int    `json:"dlcCount"`
	ReleaseDate  struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"releaseDate"`
	IsBaseProductMissing bool `json:"isBaseProductMissing"`
	IsHidingDisabled     bool `json:"isHidingDisabled"`
	IsInDevelopment      bool `json:"isInDevelopment"`
	IsHidden             bool `json:"isHidden"`
}

func (ap *AccountProduct) GetId() int {
	return ap.Id
}

func (ap *AccountProduct) GetTitle() string {
	return ap.Title
}

func (ap *AccountProduct) GetImage() string {
	return ap.Image
}
