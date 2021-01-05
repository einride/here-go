package routingv7

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// LatLng is a coordinate in the routing v7 API.
type LatLng struct {
	// Latitude of the coordinate.
	Latitude float64
	// Longitude of the coordinate.
	Longitude float64
}

var (
	_ json.Marshaler   = &LatLng{}
	_ json.Unmarshaler = &LatLng{}
)

func (l *LatLng) UnmarshalString(s string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("unmarshal lat/lng '%s': %w", s, err)
		}
	}()
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid format: '%s'", s)
	}
	latitudeStr, longitudeStr := parts[0], parts[1]
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		return fmt.Errorf("invalid latitude: %s", latitudeStr)
	}
	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		return fmt.Errorf("invalid longitude: %s", longitudeStr)
	}
	l.Latitude = latitude
	l.Longitude = longitude
	return nil
}

func (l *LatLng) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return l.UnmarshalString(unquoted)
}

func (l LatLng) MarshalJSON() ([]byte, error) {
	result := make([]byte, 0, 40)
	result = append(result, `"`...)
	result = strconv.AppendFloat(result, l.Latitude, 'f', -1, 64)
	result = append(result, ", "...)
	result = strconv.AppendFloat(result, l.Longitude, 'f', -1, 64)
	result = append(result, `"`...)
	return result, nil
}

func (l LatLng) String() string {
	return fmt.Sprintf("%f,%f", l.Latitude, l.Longitude)
}
