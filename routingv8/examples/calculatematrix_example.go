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
	routingClient := routingv8.NewClient(
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

	response, err := routingClient.Matrix.CalculateMatrix(ctx, &routingv8.CalculateMatrixRequest{
		Async: false,
		Body: &routingv8.CalculateMatrixBody{
			Origins:      origins,
			Destinations: destinations,
			RegionDefinition: routingv8.RegionDefinition{
				Type: routingv8.RegionTypeWorld,
			},
			Profile: routingv8.ProfileTruckFast,
			MatrixAttributes: &routingv8.MatrixAttributes{
				routingv8.MatrixAttributeDistances,
				routingv8.MatrixAttributeTravelTimes,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("matrix ID: %s \n", response.MatrixID)
	for i, distance := range response.Matrix.Distances {
		fmt.Printf("Route %d: %d meters \n", i, distance)
	}
}
