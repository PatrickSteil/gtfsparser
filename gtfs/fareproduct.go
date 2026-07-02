// Copyright 2026 Patrick Steil
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

// A FareProduct is a fare that can be purchased by a rider (fare_products.txt).
// Note that fare_product_id is not necessarily unique on its own -- the same
// id may appear multiple times in fare_products.txt with different
// fare_media_id/amount combinations. Feed.FareProducts therefore stores
// FareProduct slices keyed by fare_product_id.
type FareProduct struct {
	Id       string
	Name     string
	Media    *FareMedium
	Amount   string
	Currency string
}
