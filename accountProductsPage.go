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
	ContentSystemCompatibility interface{} `json:"contentSystemCompatibility"`
	MoviesCount                int         `json:"moviesCount"`
	Tags                       []struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		ProductCount string `json:"productCount"`
	} `json:"tags"`
	Products                   []AccountProduct `json:"products"`
	UpdatedProductsCount       int              `json:"updatedProductsCount"`
	HiddenUpdatedProductsCount int              `json:"hiddenUpdatedProductsCount"`
	AppliedFilters             struct {
		Tags []struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			ProductCount string `json:"productCount"`
		} `json:"tags"`
	} `json:"appliedFilters"`
	HasHiddenProducts bool `json:"hasHiddenProducts"`
}

func (app *AccountProductsPage) GetProducts() []IdGetter {
	idGetters := make([]IdGetter, 0)
	for _, ap := range app.Products {
		idGetters = append(idGetters, &ap)
	}
	return idGetters
}
