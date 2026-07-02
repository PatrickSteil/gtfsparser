// Copyright 2026 Patrick Steil
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

// A FareTransferRule defines how fare products are applied when transferring
// between legs that match From_leg_group_id/To_leg_group_id (fare_transfer_rules.txt)
type FareTransferRule struct {
	From_leg_group_id   string
	To_leg_group_id     string
	Transfer_count      int
	Duration_limit      int
	Duration_limit_type int8
	Fare_transfer_type  int8
	Fare_product_id     string
}
