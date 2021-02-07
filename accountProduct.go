// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogtypes

type AccountProduct struct {
	IsGalaxyCompatible   bool         `json:"isGalaxyCompatible"`
	Tags                 []string     `json:"tags"`
	Id                   int          `json:"id"`
	Availability         Availability `json:"availability"`
	Title                string       `json:"title"`
	Image                string       `json:"image"`
	Url                  string       `json:"url"`
	WorksOn              WorksOn      `json:"worksOn"`
	Category             string       `json:"category"`
	Rating               int          `json:"rating"`
	IsComingSoon         bool         `json:"isComingSoon"`
	IsMovie              bool         `json:"isMovie"`
	IsGame               bool         `json:"isGame"`
	Slug                 string       `json:"slug"`
	Updates              int          `json:"updates"`
	IsNew                bool         `json:"isNew"`
	DlcCount             int          `json:"dlcCount"`
	ReleaseDate          Date         `json:"releaseDate"`
	IsBaseProductMissing bool         `json:"isBaseProductMissing"`
	IsHidingDisabled     bool         `json:"isHidingDisabled"`
	IsInDevelopment      bool         `json:"isInDevelopment"`
	IsHidden             bool         `json:"isHidden"`
}
