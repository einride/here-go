HERE Go
=======

[![PkgGoDev](https://pkg.go.dev/badge/go.einride.tech/here)](https://pkg.go.dev/go.einride.tech/here) [![GoReportCard](https://goreportcard.com/badge/go.einride.tech/here)](https://goreportcard.com/report/go.einride.tech/here) [![Codecov](https://codecov.io/gh/einride/here-go/branch/master/graph/badge.svg)](https://codecov.io/gh/einride/here-go)

Go SDK for the HERE Maps API.

Documentation
-------------

API documentation can be found at [developer.here.com](https://developer.here.com).

Installing
----------

```bash
$ go get go.einride.tech/here
```

Authentication
--------------

The package does not directly handle authentication. Instead, when creating a new client, pass an `http.Client` that can handle authentication for you.

Note that when using an authenticated Client, all calls made by the client will include the same authentication data. Therefore, authenticated clients should almost never be shared between different users.

Complete Examples
-----------------

### v7 Routing API

#### Route calculation

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.einride.tech/here/routingv7"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("HERE_API_KEY")
	// Create an authenticated client
	routingClient := routingv7.New(
		routingv7.NewAPIKeyHTTPClient(apiKey, http.DefaultClient.Transport),
	)
	// Einride Gothenburg.
	origin := &routingv7.GeoWaypoint{
		Lat:  57.707752,
		Long: 11.949767,
	}
	// Einride Stockholm.
	destination := &routingv7.GeoWaypoint{
		Lat:  59.337492,
		Long: 18.063672,
	}
	// Call Here Maps API
	response, err := routingClient.Route.CalculateRoute(ctx, &routingv7.CalculateRouteRequest{
		Waypoints: []routingv7.WaypointParameter{origin, destination},
		Mode: routingv7.RoutingMode{
			Type: routingv7.RouteTypeFastest,
		},
	})
	if err != nil {
		panic(err) // TODO: Handle error.
	}
	// Handle result
	for _, route := range response.Routes {
		fmt.Println(route.Summary)
	}
}
```

### v8 Routing API

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.einride.tech/here/routingv8"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("HERE_API_KEY")
	// Create an authenticated client
	routingClient := routingv8.New(
		routingv8.NewAPIKeyHTTPClient(apiKey, http.DefaultClient.Transport),
	)
	// Einride Gothenburg.
	origins := []*routingv8.GeoWaypoint{
		{
			Lat:  57.707752,
			Long: 11.949767,
		},
	}
	// Einride Stockholm.
	destinations := []*routingv8.GeoWaypoint{
		{
			Lat:  59.337492,
			Long: 18.063672,
		},
	}
	// Call Here Maps API
	response, err := routingClient.Matrix.CalculateMatrix(ctx, &routingv8.CalculateMatrixRequest{
		Async: false,
		Body: &routingv8.CalculateMatrixBody{
			Origins:      origins,
			Destinations: destinations,
			RegionDefinition: routingv8.RegionDefinition{
				Type: routingv8.RegionType_World,
			},
			Profile: routingv8.Profile_TruckFast,
			MatrixAttributes: &routingv8.MatrixAttributes{
				routingv8.MatrixAttribute_Distances,
				routingv8.MatrixAttribute_TravelTimes,
			},
		},
	})
	if err != nil {
		panic(err) // TODO: handle error
	}
	// Handle result
	fmt.Printf("matrix ID: %s \n", response.MatrixID)
	for i, distance := range response.Matrix.Distances {
		fmt.Printf("Route %d: %d meters \n", i, distance)
	}
}
```

### v7 Geocoding & Search API

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.einride.tech/here/geocodingsearchv7"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("HERE_API_KEY")
	// Create an authenticated client
	geocodingClient := geocodingsearchv7.NewClient(
		geocodingsearchv7.NewAPIKeyHTTPClient(apiKey, http.DefaultClient.Transport),
	)

	// The query to geocode
	q := "Regeringsgatan 65, Stocholm"
	// Call Here Maps API
	response, err := geocodingClient.Geocoding.Geocoding(ctx, &geocodingsearchv7.GeocodingRequest{
		Q: &q,
	})
	if err != nil {
		panic(err) // TODO: handle error
	}
	// Handle result
	for _, item := range response.Items {
		fmt.Printf("Geocoded location lat/lng: %f %f  \n", item.Position.Lat, item.Position.Long)
	}
}
```
