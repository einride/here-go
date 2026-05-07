package routingv8

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	invalid     = "invalid"
	unspecified = "unspecified"
	none        = "None"
)

type CalculateMatrixBody struct {
	// Origins defining start points of the routes in the matrix.
	// See https://developer.here.com/documentation/matrix-routing-api/8.6.0/dev_guide/topics/modes/modes.html
	// for guidance on the matrix limitations.
	Origins []*GeoWaypoint `json:"origins"`
	// Destinations defining destinations of the routes in the matrix.
	// See https://developer.here.com/documentation/matrix-routing-api/8.6.0/dev_guide/topics/modes/modes.html
	// for guidance on the matrix limitations.
	Destinations []*GeoWaypoint `json:"destinations"`
	// DepartureTime of departure for all origins. Default to now.
	DepartureTime string `json:"departureTime,omitempty"`
	// RegionDefinition of where the matrix should be calculated.
	RegionDefinition RegionDefinition `json:"regionDefinition"`
	// Profile to use for route calculation in the matrix.
	Profile Profile `json:"profile,omitempty"`
	// RoutingMode optimization.
	RoutingMode RoutingMode `json:"routingMode,omitempty"`
	// TransportMode to use.
	TransportMode TransportMode `json:"transportMode,omitempty"`
	// MatrixAttributes to receive back in the response.
	MatrixAttributes *MatrixAttributes `json:"matrixAttributes,omitempty"`
	// Truck configuration
	Truck *Truck `json:"truck,omitempty"`
}

type CalculateMatrixRequest struct {
	// Async flag requires the Client to poll the calculation results and finally requesting to download
	// the calculation results.
	Async Async
	// Body to pass to request to Here Maps API
	Body *CalculateMatrixBody
}

type RoutesRequest struct {
	Origin        GeoWaypoint
	Destination   GeoWaypoint
	TransportMode TransportMode
	AvoidAreas    []AreaFeature
	// Which attributes to return in the response.
	// If not specified defaults to SummaryReturnAttribute.
	Return []ReturnAttribute
	// The time of departure.
	// If not specified the current time is used.
	// To not take time into account use DepartureTimeAny.
	DepartureTime string
	// Spans define which content attributes that are included in the response spans
	Spans []SpanAttribute
	// Vehicle-specific parameters.
	Vehicle *Vehicle
}

type ReturnAttribute string

const (
	PolylineReturnAttribute  ReturnAttribute = "polyline"
	SummaryReturnAttribute   ReturnAttribute = "summary"
	ElevationReturnAttribute ReturnAttribute = "elevation"
)

type GeoWaypoint struct {
	// Elevation is available in responses when using the ElevationReturnAttribute
	// TODO(odsod): Consider refactoring this to another message.
	//              We get 400 BAD REQUEST if we provide elv on outgoing messages.
	Elevation float64 `json:"elv,omitempty"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"lng"`
	// Radius instructs the router to consider all places within the given
	// radius as potential candidates for matching the waypoint. Specified
	// in meters. Range: [0-200]. Values higher than 200 meters are not
	// supported. Cannot be combined with SnapRadius.
	Radius int `json:"radius,omitempty"`
	// RadiusPenalty is a percentage applied to candidates within Radius
	// based on their air distance to the waypoint, where 100 is just the
	// cost of the air distance and 200 is twice the cost. The penalty
	// must be chosen so that, when multiplied by the radius, the result
	// is less than or equal to 7200; values up to and including 10000 are
	// accepted regardless. Range: [0-10000]. Cannot be combined with
	// SnapRadius.
	//
	// Alpha: This parameter is in development. It may not be stable and
	// is subject to change.
	RadiusPenalty int `json:"radiusPenalty,omitempty"`
	// SnapRadius instructs the router to match the waypoint, within the
	// specified radius, to the most "significant" road, sorting candidates
	// by significance (e.g. highway > national road > city road).
	// Specified in meters. Range: [0-1000000]. Cannot be combined with
	// Radius or RadiusPenalty.
	SnapRadius int `json:"snapRadius,omitempty"`
	// Course is the direction in degrees clockwise from north (0 is
	// north) from which this waypoint should be approached or in which it
	// should be left. Range: [0-359]. Values outside the range are
	// wrapped to the range.
	Course float64 `json:"course,omitempty"`
	// MinCourseDistance instructs the routing service to try to find a
	// route that avoids actions for the indicated distance. For example,
	// if the origin is determined by a moving vehicle, the user might not
	// have time to react to early actions. Specified in meters. Values
	// greater than 2000 meters will be capped at 2000 meters.
	MinCourseDistance int `json:"minCourseDistance,omitempty"`
	// SideOfStreetHint is a point next to the street (e.g. a POI) that
	// indicates which side of the street should be preferred for this
	// waypoint when the street has dividers.
	SideOfStreetHint *SideOfStreetHint `json:"sideOfStreetHint,omitempty"`
}

// SideOfStreetHint is a hint as to which side of the street should be
// preferred for a waypoint. The hint should be a point next to the
// street, e.g. a POI.
type SideOfStreetHint struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lng"`
	// Match determines how the side-of-street hint should be handled.
	// Defaults to SideOfStreetMatchOnlyIfDivided when unset.
	Match SideOfStreetMatch `json:"match,omitempty"`
}

// SideOfStreetMatch determines how a SideOfStreetHint should be handled.
type SideOfStreetMatch string

const (
	// SideOfStreetMatchAlways always uses the side-of-street hint.
	SideOfStreetMatchAlways SideOfStreetMatch = "always"
	// SideOfStreetMatchOnlyIfDivided only uses the side-of-street hint
	// on divided roads. This is HERE's default behavior.
	SideOfStreetMatchOnlyIfDivided SideOfStreetMatch = "onlyIfDivided"
)

// DepartureTimeAny enforces non time-aware routing.
const DepartureTimeAny = "any"

type Profile int

const (
	ProfileUnspecified = iota
	// ProfileCarFast - Car with fast routing mode.
	ProfileCarFast
	// ProfileCarShort - Car with short routing mode.
	ProfileCarShort
	// ProfileTruckFast - Truck with fast routing mode.
	ProfileTruckFast
	// ProfilePedestrian - Pedestrian transport mode.
	ProfilePedestrian
	// ProfileBicycle - Bicycle transport mode.
	ProfileBicycle
)

func (p *Profile) String() string {
	switch *p {
	case ProfileUnspecified:
		return unspecified
	case ProfileCarFast:
		return "carFast"
	case ProfileCarShort:
		return "carShort"
	case ProfileTruckFast:
		return "truckFast"
	case ProfilePedestrian:
		//nolint:goconst
		return "pedestrian"
	case ProfileBicycle:
		//nolint:goconst
		return "bicycle"
	default:
		return invalid
	}
}

func (p *Profile) UnmarshalString(value string) error {
	switch value {
	case "carFast":
		*p = ProfileCarFast
	case "carShort":
		*p = ProfileCarShort
	case "truckFast":
		*p = ProfileTruckFast
	case "pedestrian":
		*p = ProfilePedestrian
	case "bicycle":
		*p = ProfileBicycle
	default:
		return fmt.Errorf("invalid profile")
	}
	return nil
}

func (p *Profile) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(p.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (p *Profile) UnmarshalJSON(b []byte) error {
	value, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return p.UnmarshalString(value)
}

type RegionType int

const (
	RegionTypeUnspecified = iota
	RegionTypeWorld
	RegionTypeCircle
	RegionTypeBoundingBox
	RegionTypePolygon
	RegionTypeAutoCircle
)

func (r *RegionType) String() string {
	switch *r {
	case RegionTypeUnspecified:
		return unspecified
	case RegionTypeWorld:
		return "world"
	case RegionTypeCircle:
		return "circle"
	case RegionTypeBoundingBox:
		return "boundingBox"
	case RegionTypePolygon:
		return "polygon"
	case RegionTypeAutoCircle:
		return "autoCircle"
	default:
		return invalid
	}
}

func (r *RegionType) UnmarshalString(value string) error {
	switch value {
	case "world":
		*r = RegionTypeWorld
	case "circle":
		*r = RegionTypeCircle
	case "boundingBox":
		*r = RegionTypeBoundingBox
	case "polygon":
		*r = RegionTypePolygon
	case "autoCircle":
		*r = RegionTypeAutoCircle
	default:
		return fmt.Errorf("invalid region type")
	}
	return nil
}

func (r RegionType) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(r.String())), nil
}

func (r *RegionType) UnmarshalJSON(b []byte) error {
	value, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return r.UnmarshalString(value)
}

type RegionDefinition struct {
	Type RegionType `json:"type"`
	// Circle
	CircleCenter *GeoWaypoint `json:"center,omitempty"`
	CircleRadius int          `json:"radius,omitempty"`
	// BoundingBox
	BoundingBoxNorth int `json:"north,omitempty"`
	BoundingBoxEast  int `json:"east,omitempty"`
	BoundingBoxSouth int `json:"south,omitempty"`
	BoundingBoxWest  int `json:"west,omitempty"`
	// Polygon
	PolygonOuter []*GeoWaypoint `json:"outer,omitempty"`
	// AutoCircle
	AutoCircleMargin int `json:"margin,omitempty"`
}

type Async bool

func (a Async) String() string {
	if a {
		return "true"
	}
	return "false"
}

type MatrixAttribute int

const (
	MatrixAttributeUnspecified MatrixAttribute = iota
	MatrixAttributeTravelTimes
	MatrixAttributeDistances
)

func (m *MatrixAttribute) String() string {
	switch *m {
	case MatrixAttributeUnspecified:
		return unspecified
	case MatrixAttributeTravelTimes:
		return "travelTimes"
	case MatrixAttributeDistances:
		return "distances"
	default:
		return invalid
	}
}

type MatrixAttributes []MatrixAttribute

func (m *MatrixAttributes) MarshalJSON() ([]byte, error) {
	attributes := make([]string, 0, len(*m))
	for _, attr := range *m {
		attributes = append(attributes, attr.String())
	}
	b, err := json.Marshal(attributes)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type RoutingMode int

const (
	RoutingModeUnspecified RoutingMode = iota
	RoutingModeFast
	RoutingModeShort
)

func (r *RoutingMode) String() string {
	switch *r {
	case RoutingModeUnspecified:
		return unspecified
	case RoutingModeFast:
		return "fast"
	case RoutingModeShort:
		return "short"
	default:
		return invalid
	}
}

func (r *RoutingMode) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(r.String())), nil
}

type TransportMode int

const (
	TransportModeUnspecified TransportMode = iota
	TransportModeCar
	TransportModeTruck
	TransportModePedestrian
	TransportModeBicycle
	TransportModeTaxi
	TransportModeScooter
)

func (t *TransportMode) String() string {
	switch *t {
	case TransportModeUnspecified:
		return unspecified
	case TransportModeCar:
		return "car"
	case TransportModeTruck:
		return "truck"
	case TransportModePedestrian:
		return "pedestrian"
	case TransportModeBicycle:
		return "bicycle"
	case TransportModeTaxi:
		return "taxi"
	case TransportModeScooter:
		return "scooter"
	default:
		return invalid
	}
}

func (t *TransportMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(t.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

type ShippedHazardousGoods int

const (
	ShippedHazardousGoodsUnspecified ShippedHazardousGoods = iota
	ShippedHazardousGoodsExplosive
	ShippedHazardousGoodsGas
	ShippedHazardousGoodsFlammable
	ShippedHazardousGoodsCombustible
	ShippedHazardousGoodsOrganic
	ShippedHazardousGoodsPoison
	ShippedHazardousGoodsRadioactive
	ShippedHazardousGoodsCorrosive
	ShippedHazardousGoodsPoisonousInhalation
	ShippedHazardousGoodsHarmfulToWater
	ShippedHazardousGoodsOther
)

func (s *ShippedHazardousGoods) String() string {
	switch *s {
	case ShippedHazardousGoodsUnspecified:
		return unspecified
	case ShippedHazardousGoodsExplosive:
		return "explosive"
	case ShippedHazardousGoodsGas:
		return "gas"
	case ShippedHazardousGoodsFlammable:
		return "flammable"
	case ShippedHazardousGoodsCombustible:
		return "combustible"
	case ShippedHazardousGoodsOrganic:
		return "organic"
	case ShippedHazardousGoodsPoison:
		return "poison"
	case ShippedHazardousGoodsRadioactive:
		return "radioactive"
	case ShippedHazardousGoodsCorrosive:
		return "corrosive"
	case ShippedHazardousGoodsPoisonousInhalation:
		return "poisonousInhalation"
	case ShippedHazardousGoodsHarmfulToWater:
		return "harmfulToWater"
	case ShippedHazardousGoodsOther:
		return "other"
	default:
		return invalid
	}
}

func (s *ShippedHazardousGoods) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

type ShippedHazardousGoodsList []ShippedHazardousGoods

func (s *ShippedHazardousGoodsList) MarshalJSON() ([]byte, error) {
	goods := make([]string, 0, len(*s))
	for _, g := range *s {
		goods = append(goods, g.String())
	}
	b, err := json.Marshal(goods)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type TunnelCategory int

const (
	TunnelCategoryUnspecified TunnelCategory = iota
	TunnelCategoryB
	TunnelCategoryC
	TunnelCategoryD
	TunnelCategoryE
)

func (t *TunnelCategory) String() string {
	switch *t {
	case TunnelCategoryUnspecified:
		return none
	case TunnelCategoryB:
		return "B"
	case TunnelCategoryC:
		return "C"
	case TunnelCategoryD:
		return "D"
	case TunnelCategoryE:
		return "E"
	default:
		return invalid
	}
}

func (t *TunnelCategory) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(t.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

type Truck struct {
	ShippedHazardousGoods ShippedHazardousGoodsList `json:"shippedHazardousGoods"`
	GrossWeight           int                       `json:"grossWeight"`
	WeightPerAxle         int                       `json:"weightPerAxle"`
	Height                int                       `json:"height"`
	Width                 int                       `json:"width"`
	Length                int                       `json:"length"`
	TunnelCategory        TunnelCategory            `json:"tunnelCategory"`
	AxleCount             int                       `json:"axleCount"`
	TrailerCount          int                       `json:"trailerCount"`
}

type AreaFeature int

const (
	AreaFeatureUnspecified AreaFeature = iota
	AreaFeatureFerry
	AreaFeatureTollRoad
	AreaFeatureTunnel
	AreaFeatureControlledAccessHighway
)

func (t *AreaFeature) String() string {
	switch *t {
	case AreaFeatureUnspecified:
		return unspecified
	case AreaFeatureFerry:
		return "ferry"
	case AreaFeatureTollRoad:
		return "tollRoad"
	case AreaFeatureTunnel:
		return "tunnel"
	case AreaFeatureControlledAccessHighway:
		return "controlledAccessHighway"
	default:
		return invalid
	}
}

func (t *AreaFeature) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(t.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

type SpanAttribute string

// For available span attributes to implementation see:
// https://www.here.com/docs/bundle/routing-api-v8-api-reference/page/index.html#tag/Routing/operation/calculateRoutes
const (
	SpanAttributeNames    SpanAttribute = "names"
	SpanAttributeMaxSpeed SpanAttribute = "maxSpeed"
	SpanAttributeLength   SpanAttribute = "length"
)

func (t *SpanAttribute) String() string {
	switch *t {
	case SpanAttributeNames:
		return string(SpanAttributeNames)
	case SpanAttributeMaxSpeed:
		return string(SpanAttributeMaxSpeed)
	case SpanAttributeLength:
		return string(SpanAttributeLength)
	default:
		return invalid
	}
}

type RouteImportRequest struct {
	TransportMode TransportMode
	// Which attributes to return in the response.
	// If not specified defaults to SummaryReturnAttribute.
	Return []ReturnAttribute
	// The time of departure.
	// If not specified the current time is used.
	// To not take time into account use DepartureTimeAny.
	DepartureTime string
	// Spans define which content attributes that are included in the response spans
	Spans []SpanAttribute
	// An array of GPS coordinates
	Trace []GeoWaypoint
	// Vehicle-specific parameters.
	Vehicle *Vehicle
}

type RouteImportRequestBody struct {
	Trace []GeoWaypoint `json:"trace"`
}

type Vehicle struct {
	// Gross vehicle weight, including trailers and shipped goods when loaded at capacity, specified in kilograms.
	// Needs to be >=0.
	// If unspecified, it will default to currentWeight.
	// If neither parameter has a value specified, it will default to 0.
	// Notes:
	// - Supported in truck, bus, privateBus, car (Beta), taxi (Beta) transport modes.
	// - Maximum weight for a car or taxi without a trailer is 4250 kg.
	// - Maximum weight for a car or taxi with a trailer is 7550 kg.
	GrossWeight int
	// The number of trailers attached to the vehicle. Range: [0-255]. Default: 0.
	// Maximum value when used with transportMode=car or transportMode=taxi is 1.
	// Limitations: Considered for route calculation when transportMode is one of (truck, bus, privateBus).
	// Considered for route calculation for restrictions, but not for speed limits, when transportMode is car or taxi.
	TrailerCount int
	// Specifies the total number of axles the vehicle has, i.e., axles on the base vehicle and any attached trailers.
	// Range: [2-255].
	// Note: Supported in truck, bus, privateBus, car (Beta), taxi (Beta) transport modes.
	AxleCount int
	// Vehicle height, specified in centimeters. Range: [0-5000].
	// Note: Supported in truck, bus, privateBus, car (Beta), taxi (Beta) transport modes.
	Height int
	// Vehicle width, specified in centimeters. Range: [0-5000].
	// Note: Supported in truck, bus, privateBus, car (Beta), taxi (Beta) transport modes.
	Width int
	// Vehicle length, specified in centimeters. Range: [0-30000].
	// Note: Supported in truck, bus, privateBus, car (Beta), taxi (Beta) transport modes.
	Length int
	// Specifies the type of the vehicle. Limitations: only valid for transportMode=truck.
	Type VehicleType
}

type VehicleType string

const (
	// A truck on a single frame with a permanently attached cargo area.
	// Note: default value when truck routing mode is used.
	VehicleTypeStraightTruck VehicleType = "StraightTruck"
	// A towing vehicle that can pull one or more semi-trailers (also known as a semi-truck).
	VehicleTypeTractor VehicleType = "Tractor"
)

func (t *VehicleType) String() string {
	switch *t {
	case VehicleTypeStraightTruck:
		return string(VehicleTypeStraightTruck)
	case VehicleTypeTractor:
		return string(VehicleTypeTractor)
	default:
		return invalid
	}
}
