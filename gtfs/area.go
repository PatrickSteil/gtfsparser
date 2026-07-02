// Copyright 2026 Patrick Steil
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

// An Area is a grouping of stops (areas.txt), used to define
// origin/destination conditions in fare_leg_rules.txt
type Area struct {
	Id   string
	Name string
}
