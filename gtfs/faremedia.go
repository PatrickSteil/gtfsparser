// Copyright 2026 Patrick Steil
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

// A FareMedium is a fare media that can be employed to use fare products,
// for example a transit card or a contactless bank card (fare_media.txt)
type FareMedium struct {
	Id   string
	Name string
	Type int8
}
