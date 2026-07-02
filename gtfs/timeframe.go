// Copyright 2026 Patrick Steil
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

// A Timeframe defines a time range that, together with a service,
// is referenced by a timeframe_group_id from fare_leg_rules.txt
type Timeframe struct {
	Group_id   string
	Start_time Time
	End_time   Time
	Service    *Service
}
