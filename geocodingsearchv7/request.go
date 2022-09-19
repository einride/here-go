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
	// The country name.
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

type ReverseGeocodingRequest struct {
	// Specify the coordinates to perform a reverse geocoding. Returns the closest address given the coordinates.
	GeoPosition *GeoWaypoint
	// Search within a geographic area.
	// Results will be returned if they are located within the specified area.
	// a country (or multiple countries), provided as comma-separated ISO 3166-1 alpha-3 country codes.
	In *string
}
