// Copyright 2015 geOps
// Authors: patrick.brosi@geops.de
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

import (
	url "net/url"
)

// A Stop object describes a single stop location
type Stop struct {
	Id                  string
	Code                string
	Name                string
	Desc                string
	Has_LatLon          bool
	Lat                 float32
	Lon                 float32
	Zone_id             string
	Url                 *url.URL
	Location_type       int8
	Parent_station      *Stop
	Timezone            Timezone
	Wheelchair_boarding int8
	Level               *Level
	Platform_code       string
}
