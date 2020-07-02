# routingv7

routingv7 is a Go client library for accessing the routing v7 HERE API

API docs can be found [here](https://developer.here.com/documentation/routing/dev_guide/topics/request-constructing.html)

## Usage

```go
import routingv7 "github.com/einride/here-go/pkg/routing/v7"
```

Construct a new HERE client, then use the various services on the client to
access different parts of the HERE API. For example:

```go
httpClient := routingv7.NewAPIKeyHTTPClient(".... your api key ...")
client := routingv7.NewClient(httpClient)

// Get routes from Einride Gothenburg to Einride Stockholm
resp, err := client.Route.CalculateRoute(context.Background(), &routingv7.CalculateRouteRequest{
    Waypoints: []routingv7.WaypointParameter{
        &routingv7.GeoWaypoint{
            Lat:  57.707752,
            Long: 11.949767,
        },
        &routingv7.GeoWaypoint{
            Lat:  59.337492,
            Long: 18.063672,
        },
    },
    Mode: routingv7.RoutingMode{
        Type: routingv7.RouteTypeFastest,
    },
})
```

### Authentication

The package  not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you.

Using an API Key:

```go
func main() {
	httpClient := routingv7.NewAPIKeyHTTPClient(".... your api key ...")
	client := routingv7.New(httpClient)
}
```

Note that when using an authenticated Client, all calls made by the client will
include the same authentication data. Therefore, authenticated clients should
almost never be shared between different users.
