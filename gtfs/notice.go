// Copyright 2026 Patrick Steil
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

// A Notice is a short, human-readable, non-structured message that can be
// shown to riders. Notices are defined in notices.txt and attached to
// routes, trips, trip segments or stops via NoticeAssignment.
type Notice struct {
	Id           string
	GroupId      string
	DisplayText  string
	Translations []*Translation
}

// A NoticeAssignment attaches a Notice (or all Notices sharing a
// notice_group_id) to a route, trip, trip segment or stop. Exactly one of
// Notice or GroupId is set.
type NoticeAssignment struct {
	// Notice is set if this assignment references a single notice via
	// notice_id. Mutually exclusive with GroupId.
	Notice *Notice

	// GroupId is set if this assignment references a notice group via
	// notice_group_id. Mutually exclusive with Notice.
	GroupId string
}

// A TripSegment defines a contiguous part of a Trip, specified by an
// inclusive range of stop_sequence values from stop_times.txt. Notices can
// be assigned to a TripSegment via NoticeAssignment, which is useful if a
// notice does not apply to an entire trip.
type TripSegment struct {
	Id                 string
	Trip               *Trip
	From_stop_sequence int
	To_stop_sequence   int
	NoticeAssignments  []*NoticeAssignment
}
