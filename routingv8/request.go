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
}

type GeoWaypoint struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lng"`
}

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
		// nolint:goconst
		return "pedestrian"
	case ProfileBicycle:
		// nolint:goconst
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
