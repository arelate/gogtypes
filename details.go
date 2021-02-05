// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogtypes

type Details struct {
	Title           string `json:"title"`
	BackgroundImage string `json:"backgroundImage"`
	CdKey           string `json:"cdKey"`
	TextInformation string `json:"textInformation"`
	// TODO: Create a helper func to process Downloads
	Downloads [][]interface{} `json:"downloads"`
	// TODO: find data examples where GalaxyDownloads is not empty and create a type from that
	GalaxyDownloads  []interface{} `json:"galaxyDownloads"`
	Extras           []Extra       `json:"extras"`
	DLCs             []Details     `json:"dlcs"`
	Tags             []Tag         `json:"tags"`
	IsPreOrder       bool          `json:"isPreOrder"`
	ReleaseTimestamp int           `json:"releaseTimestamp"`
	// TODO: find data examples where Messages is not empty and create a type from that
	Messages             []interface{} `json:"messages"`
	Changelog            string        `json:"changelog"`
	ForumLink            string        `json:"forumLink"`
	IsBaseProductMissing bool          `json:"isBaseProductMissing"`
	// TODO: find data examples where MissingBaseProduct is not empty and create a type from that
	MissingBaseProduct     interface{}             `json:"missingBaseProduct"`
	Features               []string                `json:"features"`
	SimpleGalaxyInstallers []SimpleGalaxyInstaller `json:"simpleGalaxyInstallers"`
}
