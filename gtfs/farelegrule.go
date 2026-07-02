// Copyright 2026 Patrick Steil
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

// A FareLegRule assigns fare products to "legs" -- itineraries with a
// single fare-relevant network -- that match a set of conditions
// (fare_leg_rules.txt)
type FareLegRule struct {
	Leg_group_id            string
	Network                 *Network
	From_area               *Area
	To_area                 *Area
	From_timeframe_group_id string
	To_timeframe_group_id   string
	Fare_product_id         string
	Rule_priority           int
}
