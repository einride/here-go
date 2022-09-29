package geocodingsearchv7

type GeocodingResponse struct {
	Items []GeocodingItem
}

type GeocodingItem struct {
	// A representative string for the result, for instance the name of a place, or a complete address.
	Title string `json:"title,omitempty"`
	// The identifier of the item. Its value can be used to retrieve the very same object using the /lookup endpoint.
	ID string `json:"id,omitempty"`
	// HERE Geocoding and Search supports multiple location object types (place, street, locality, ...).
	ResultType string `json:"resultType,omitempty"`
	// Type of address data (returned only for address results). Is one of PA or interpolated.
	// PA - Point Address, location matches as individual point object.
	// interpolated - location was interpolated from an address range.
	HouseNumberType string `json:"houseNumberType,omitempty"`
	// The result address in its related fields.
	Address Address `json:"address,omitempty"`
	// A representative geo-position (WGS 84) of the result; this is to be used to display the result on a map.
	Position GeoWaypoint `json:"position,omitempty"`
	// The geo-position of the access to the result (for instance the entrance).
	Access []GeoWaypoint `json:"access,omitempty"`
	// Bounding box of the location optimized for display.
	MapView MapView `json:"mapView,omitempty"`
	// indicates for each result how good it matches to the original query.
	// This can be used by the customer application to accept or reject
	// the results depending on how “expensive” is the mistake for their use case.
	Scoring Scoring `json:"scoring,omitempty"`
}

type ReverseGeocodingResponse struct {
	Items []ReverseGeocodingItem
}

type ReverseGeocodingItem struct {
	// A representative string for the result, for instance the name of a place, or a complete address.
	Title string `json:"title,omitempty"`
	// The identifier of the item. Its value can be used to retrieve the very same object using the /lookup endpoint.
	ID string `json:"id,omitempty"`
	// HERE Geocoding and Search supports multiple location object types (place, street, locality, ...).
	ResultType string `json:"resultType,omitempty"`
	// Type of address data (returned only for address results). Is one of PA or interpolated.
	// PA - Point Address, location matches as individual point object.
	// interpolated - location was interpolated from an address range.
	HouseNumberType string `json:"houseNumberType,omitempty"`
	// The result address in its related fields.
	Address Address `json:"address,omitempty"`
	// A representative geo-position (WGS 84) of the result; this is to be used to display the result on a map.
	Position GeoWaypoint `json:"position,omitempty"`
	// The geo-position of the access to the result (for instance the entrance).
	Access []GeoWaypoint `json:"access,omitempty"`
	// Bounding box of the location optimized for display.
	MapView MapView `json:"mapView,omitempty"`
	// The distance in meters to the given spatial context ('at=lat,lon').
	Distance int `json:"distance,omitempty"`
}

type Address struct {
	// A representative string for the result, for instance the name of a place, or a complete address.
	Label string `json:"label,omitempty"`
	// The country code.
	CountryCode string `json:"countryCode,omitempty"`
	// The country name.
	CountryName string `json:"countryName,omitempty"`
	// The state code.
	StateCode string `json:"stateCode,omitempty"`
	// The state name.
	State string `json:"state,omitempty"`
	// The county code.
	CountyCode string `json:"countyCode,omitempty"`
	// The county name.
	CountyName string `json:"countyName,omitempty"`
	// The city name.
	City string `json:"city,omitempty"`
	// The district.
	District string `json:"district,omitempty"`
	// The street name.
	Street string `json:"street,omitempty"`
	// The postal code.
	PostalCode string `json:"postalCode,omitempty"`
	// The house number.
	HouseNumber string `json:"houseNumber,omitempty"`
}

type MapView struct {
	West  float64 `json:"west,omitempty"`
	South float64 `json:"south,omitempty"`
	East  float64 `json:"east,omitempty"`
	North float64 `json:"north,omitempty"`
}

type Scoring struct {
	QueryScore float64    `json:"queryScore,omitempty"`
	FieldScore FieldScore `json:"fieldScore,omitempty"`
}

type FieldScore struct {
	City        float64   `json:"city,omitempty"`
	Streets     []float64 `json:"streets,omitempty"`
	HouseNumber float64   `json:"houseNumber,omitempty"`
}

// HereErrorResponse is returned when an error is returned from the Here Maps API.
type HereErrorResponse struct {
	// Title of the error
	Title string `json:"title" xml:"title"`
	// Http status code
	Status int `json:"status" xml:"status"`
	// Here Maps API error code
	Code string `json:"code" xml:"code"`
	// Cause of the error
	Cause string `json:"cause" xml:"cause"`
	// Action Suggested to fix error
	Action string `json:"action" xml:"action"`
}

type BatchGeocoderResponse struct {
	Response struct {
		MetaInfo struct {
			RequestID string `xml:"RequestId,omitempty"`
		}
		Status         JobStatus `xml:"Status,omitempty"`
		TotalCount     int       `xml:"TotalCount,omitempty"`
		ValidCount     int       `xml:"ValidCount,omitempty"`
		InvalidCount   int       `xml:"InvalidCount,omitempty"`
		ProcessedCount int       `xml:"ProcessedCount,omitempty"`
		PendingCount   int       `xml:"PendingCount,omitempty"`
		SuccessCount   int       `xml:"SuccessCount,omitempty"`
		ErrorCount     int       `xml:"ErrorCount,omitempty"`
		JobStarted     string    `xml:"JobStarted,omitempty"`
		JobFinished    string    `xml:"JobFinished,omitempty"`
	}
}

type JobStatus string

const (
	JobStatusAccepted  = "accepted"
	JobStatusCancelled = "canceled"
	JobStatusCompleted = "completed"
	JobStatusDeleted   = "deleted"
	JobStatusFailed    = "failed"
	JobStatusRunning   = "running"
	JobStatusSubmitted = "submitted"
)

type BatchGeocoderResponseRow struct {
	RecID            string  `csv:"recId"`
	SeqNumber        int     `csv:"SeqNumber"`
	SeqLength        int     `csv:"seqLength"`
	DisplayLatitude  float64 `csv:"displayLatitude"`
	DisplayLongitude float64 `csv:"displayLongitude"`
	LocationLabel    string  `csv:"locationLabel"`
	HouseNumber      string  `csv:"houseNumber"`
	Street           string  `csv:"street"`
	District         string  `csv:"district"`
	City             string  `csv:"city"`
	PostalCode       string  `csv:"postalCode"`
	County           string  `csv:"county"`
	State            string  `csv:"state"`
	Country          string  `csv:"country"`
}
