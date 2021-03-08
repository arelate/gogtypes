// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_types

type Details struct {
	Title           string `json:"title"`
	BackgroundImage string `json:"backgroundImage"`
	CdKey           string `json:"cdKey"`
	TextInformation string `json:"textInformation"`
	// TODO: Create a helper func to process Downloads
	Downloads []interface{} `json:"downloads"`
	// TODO: find data examples where GalaxyDownloads is not empty and create a type from that
	GalaxyDownloads []interface{} `json:"galaxyDownloads"`
	Extras          []struct {
		ManualUrl string `json:"manualUrl"`
		Name      string `json:"name"`
		Type      string `json:"type"`
		Info      int    `json:"info"`
		Size      string `json:"size"`
	} `json:"extras"`
	DLCs []Details `json:"dlcs"`
	Tags []struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		ProductCount string `json:"productCount"`
	} `json:"tags"`
	IsPreOrder           bool     `json:"isPreOrder"`
	ReleaseTimestamp     int      `json:"releaseTimestamp"`
	Messages             []string `json:"messages"`
	Changelog            string   `json:"changelog"`
	ForumLink            string   `json:"forumLink"`
	IsBaseProductMissing bool     `json:"isBaseProductMissing"`
	// TODO: find data examples where MissingBaseProduct is not empty and create a type from that
	MissingBaseProduct     interface{} `json:"missingBaseProduct"`
	Features               []string    `json:"features"`
	SimpleGalaxyInstallers []struct {
		Path string `json:"path"`
		Os   string `json:"os"`
	} `json:"simpleGalaxyInstallers"`
}

func (det *Details) GetTitle() string {
	return det.Title
}

func (det *Details) GetBackgroundImage() string {
	return det.BackgroundImage
}
