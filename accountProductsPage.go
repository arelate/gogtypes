// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_types

type AccountProductsPage struct {
	Page
	SortBy          string `json:"sortBy"`
	TotalProducts   int    `json:"totalProducts"`
	ProductsPerPage int    `json:"productsPerPage"`
	// TODO: find data examples where ContentSystemCompatibility is not empty and create a type from that
	ContentSystemCompatibility interface{}      `json:"contentSystemCompatibility"`
	MoviesCount                int              `json:"moviesCount"`
	Tags                       []Tag            `json:"tags"`
	Products                   []AccountProduct `json:"products"`
	UpdatedProductsCount       int              `json:"updatedProductsCount"`
	HiddenUpdatedProductsCount int              `json:"hiddenUpdatedProductsCount"`
	AppliedFilters             AppliedFilter    `json:"appliedFilters"`
	HasHiddenProducts          bool             `json:"hasHiddenProducts"`
}
