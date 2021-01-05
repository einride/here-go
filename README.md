# HERE Go

[![PkgGoDev][pkg-badge]][pkg]
[![GoReportCard][report-badge]][report]
[![Codecov][codecov-badge]][codecov]

[pkg-badge]: https://pkg.go.dev/badge/go.einride.tech/here
[pkg]: https://pkg.go.dev/go.einride.tech/here
[report-badge]: https://goreportcard.com/badge/go.einride.tech/here
[report]: https://goreportcard.com/report/go.einride.tech/here
[codecov-badge]: https://codecov.io/gh/einride/here-go/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/einride/here-go

Go SDK for the HERE Maps API.

## Documentation

API documentation can be found at [developer.here.com][api-docs].

[api-docs]: https://developer.here.com

## Installing

```bash
$ go get go.einride.tech/here
```

## Examples

### v7 Routing API

Construct a new HERE client, then use the various services on the client
to access different parts of the HERE API.

#### Authentication

The package does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle
authentication for you.

Note that when using an authenticated Client, all calls made by the client
will include the same authentication data. Therefore, authenticated
clients should almost never be shared between different users.

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
	response, err := routingClient.Route.CalculateRoute(ctx, &routingv7.CalculateRouteRequest{
		Waypoints: []routingv7.WaypointParameter{origin, destination},
		Mode: routingv7.RoutingMode{
			Type: routingv7.RouteTypeFastest,
		},
	})
	if err != nil {
		panic(err) // TODO: Handle error.
	}
	for _, route := range response.Routes {
		fmt.Println(route.Summary)
	}
}
```
