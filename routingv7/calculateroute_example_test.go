package routingv7_test

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.einride.tech/here/routingv7"
)

func ExampleRouteService_CalculateRoute() {
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
