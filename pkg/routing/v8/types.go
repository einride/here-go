package v8

import (
	"fmt"
	"strconv"
)

type Route struct {
	ID string `json:"id"`
}

type TransportMode int

const (
	TransportModeCar TransportMode = iota
	TransportModeTruck
	TransportModePedestrian
)

func (t TransportMode) String() string {
	var out string
	switch t {
	case TransportModeCar:
		out = "car"
	case TransportModeTruck:
		out = "truck"
	case TransportModePedestrian:
		out = "pedestrian"
	}
	return out
}

type MatchSideOfStreet int

const (
	MatchSoSUnspecified MatchSideOfStreet = iota
	MatchSoSOnlyIfDivided
	MatchSoSAlways
)

func (h MatchSideOfStreet) String() string {
	var out string
	switch h {
	case MatchSoSAlways:
		out = "always"
	case MatchSoSOnlyIfDivided:
		out = "onlyIfDivided"
	}
	return out
}

type LatLong struct {
	Lat  float64
	Long float64
}

func (l LatLong) String() string {
	return strconv.FormatFloat(l.Lat, 'f', 8, 64) + "," + strconv.FormatFloat(l.Long, 'f', 8, 64)
}

type PlaceOptions struct {
	// Degrees clock-wise from north. Indicating desired direction at the place. E.g. 90 indicating east.
	Course int
	// SideOfStreetHint indicated the side of the street that should be used. E.g. if the location is to
	// the left of the street, the router will prefer using that side in case the street has dividers.
	// E.g. If Place is set to 52.511496,13.304140 and sideOfStreetHint=52.512149,13.304076 indicates that
	// the north side of the street should be preferred.
	// This option is required, if MatchSideOfStreet is set to always.
	SideOfStreetHint LatLong
	// MatchSideOfStreet specifies how the location set by sideOfStreetHint should be handled.
	// Requires SideOfStreetHint to be specified as well.
	//     always : 	   Always prefer the given side of street.
	//     onlyIfDivided: Only prefer using side of street set by SideOfStreetHint in case the street has dividers.
	//                    This is the default behavior.
	MatchSideOfStreet MatchSideOfStreet
	// NameHint causes the router to look for the place with the most similar name.
	// This can e.g. include things like:
	//     `North` being used to differentiate between interstates `I66 North` and `I66 South`.
	//     `Downtown Avenue` being used to correctly select a residential street.
	NameHint string
	// Radius asks the router to consider all places within the given radius as potential candidates for route calculation.
	// This can be either because it is not important which place is used, or because it is unknown.
	// Radius more than 200meter are not supported.
	// Specified in Meters.
	Radius int
	// Ask the routing service to try find a route that avoids actions for the indicated distance.
	// E.g. if the origin is determined by a moving vehicle, the user might not have time to react to early actions.
	// Specified in Meters.
	MinCourseDistance int
}

func (o *PlaceOptions) String() string {
	var out string

	if o.Course != 0 {
		out += fmt.Sprintf(";%d", o.Course)
	}
	if o.SideOfStreetHint.Lat != 0 || o.SideOfStreetHint.Long != 0 {
		out += ";" + o.SideOfStreetHint.String()
	}
	if o.MatchSideOfStreet != MatchSoSUnspecified {
		out += ";" + o.MatchSideOfStreet.String()
	}
	if o.NameHint != "" {
		out += ";" + o.NameHint
	}
	if o.Radius != 0 {
		out += fmt.Sprintf(";%d", o.Radius)
	}
	if o.MinCourseDistance != 0 {
		out += fmt.Sprintf(";%d", o.MinCourseDistance)
	}
	return out
}

type WaypointOptions struct {
	// Desired duration for the stop, in seconds.
	// The section arriving at this via waypoint will have a wait post action reflecting the stopping time.
	// The consecutive section will start at the arrival time of the former section + stop duration.
	StopDuration int64
}

func (o *WaypointOptions) String() string {
	var out string

	if o.StopDuration != 0 {
		out += fmt.Sprintf("!%d", o.StopDuration)
	}
	return out
}

type Waypoint struct {
	Place           LatLong
	PlaceOptions    PlaceOptions
	WaypointOptions WaypointOptions
}

func (w *Waypoint) String() string {
	var out string

	out += w.Place.String()
	out += w.PlaceOptions.String()
	out += w.WaypointOptions.String()
	return out
}
