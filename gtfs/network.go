// Copyright 2026 Patrick Steil
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

// A Network is a grouping of routes (networks.txt / routes.txt#network_id),
// used to define fare-relevant route groupings in fare_leg_rules.txt
type Network struct {
	Id   string
	Name string
}
