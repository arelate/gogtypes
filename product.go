// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_types

type Product struct {
	CustomAttributes          interface{}     `json:"customAttributes"`
	Developer                 string          `json:"developer"`
	Publisher                 string          `json:"publisher"`
	Gallery                   []string        `json:"gallery"`
	Video                     Video           `json:"video"`
	SupportedOperatingSystems []string        `json:"supportedOperatingSystems"`
	Genres                    []string        `json:"genres"`
	GlobalReleaseDate         int             `json:"globalReleaseDate"`
	IsTBA                     bool            `json:"isTBA"`
	Price                     Price           `json:"price"`
	IsDiscounted              bool            `json:"isDiscounted"`
	IsInDevelopment           bool            `json:"isInDevelopment"`
	Id                        int             `json:"id"`
	ReleaseDate               int             `json:"releaseDate"`
	Availability              Availability    `json:"availability"`
	SalesVisibility           SalesVisibility `json:"salesVisibility"`
	Buyable                   bool            `json:"buyable"`
	Title                     string          `json:"title"`
	Image                     string          `json:"image"`
	Url                       string          `json:"url"`
	SupportUrl                string          `json:"supportUrl"`
	ForumUrl                  string          `json:"forumUrl"`
	WorksOn                   WorksOn         `json:"worksOn"`
	Category                  string          `json:"category"`
	OriginalCategory          string          `json:"originalCategory"`
	Rating                    int             `json:"rating"`
	Type                      int             `json:"type"`
	IsComingSoon              bool            `json:"isComingSoon"`
	IsPriceVisible            bool            `json:"isPriceVisible"`
	IsMovie                   bool            `json:"isMovie"`
	IsGame                    bool            `json:"isGame"`
	Slug                      string          `json:"slug"`
	IsWishlistable            bool            `json:"isWishlistable"`
}
