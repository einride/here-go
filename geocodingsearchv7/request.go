package geocodingsearchv7

type GeoWaypoint struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lng"`
}

type GeocodingRequest struct {
	// Specify the center of the search context expressed as coordinates.
	GeoPosition *GeoWaypoint
	// Free text search query. Example: "Regeringsgatan 65, Stockholm"
	// Either Q, or Address-parameter is required.
	Q *string
	// The address to search for. Works similar to free text query but separates parameters.
	Address *AddressRequest
	// Search within a geographic area.
	// Results will be returned if they are located within the specified area.
	// a country (or multiple countries), provided as comma-separated ISO 3166-1 alpha-3 country codes.
	In *string
}

type AddressRequest struct {
	// The id to recognize this address with, will be returned in result from HERE as recId
	RecID string `json:"recId,omitempty"`
	// The country name or country code in ISO 3166-1-alpha-3 format.
	Country string `json:"country,omitempty"`
	// The state name.
	State string `json:"state,omitempty"`
	// The county name.
	County string `json:"county,omitempty"`
	// The city name.
	City string `json:"city,omitempty"`
	// The district.
	District string `json:"district,omitempty"`
	// The street name.
	Street string `json:"street,omitempty"`
	// The house number.
	HouseNumber string `json:"houseNumber,omitempty"`
	// The postal code.
	PostalCode string `json:"postalCode,omitempty"`
}

type GeoWaypointRequest struct {
	// The id to recognize this address with, will be returned in result from HERE as recId
	RecID string `json:"recId,omitempty"`
	// Coordinates to perform reverse geocoding on to retrieve the address.
	GeoPositions *GeoWaypoint
}

type ReverseGeocodingRequest struct {
	// Specify the coordinates to perform a reverse geocoding. Returns the closest address given the coordinates.
	GeoPosition *GeoWaypoint
	// Search within a geographic area.
	// Results will be returned if they are located within the specified area.
	// a country (or multiple countries), provided as comma-separated ISO 3166-1 alpha-3 country codes.
	In *string
}

type BatchGeocoderUploadRequest struct {
	// List of free text search query, one query per address.
	Queries []*QueryString
	// The addresses to search for. Works similar to free text query but separates parameters.
	Addresses []*AddressRequest
}

type BatchReverseGeocoderUploadRequest struct {
	// Coordinates to perform reverse geocoding on to retrieve the address, and id.
	GeoPositions []*GeoWaypointRequest
}

type BatchGeocoderStatusRequest struct {
	// RequestID for the job to fetch the status on, was obtained from the call to BatchGeocoderUpload.
	RequestID string
}

type BatchGeocoderDownloadRequest struct {
	// RequestID for the job to fetch the result on, was obtained from the call to BatchGeocoderUpload.
	RequestID string
}

type QueryString struct {
	// The id to recognize this address with, will be returned in result from HERE as recId
	RecID string `json:"recId,omitempty"`
	// The free text query.
	Query string
	// The country of where the free query refers to, in ISO 3166-1 alpha-3 format.
	Country string
}
