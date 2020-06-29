# here-go

here-go is a Go client library for accessing the HERE APIs

API docs can be found [here](https://developer.here.com/documentation/)

The Client is mostly generated from HERE OpenAPI specifications using the
[OpenAPI generator](https://github.com/OpenAPITools/openapi-generator)

A few non-generated types exist to improve usability of the client

## Usage

```go
import (
    "context"
	"fmt"

	"github.com/antihax/optional"
    routingv8 "github.com/einride/here-go/pkg/openapi/routing/v8"
    v8types "github.com/einride/here-go/pkg/routing/v8"
    
)

func main() {
    // Using API Key
    auth := context.WithValue(context.Background(), routingv8.ContextAPIKey, routingv8.APIKey{
        Key: "....your api key here ....",
    })
    // Construct a new client
    client := routingv8.NewAPIClient(routingv8.NewConfiguration())

    // v8types.Waypoint is not a generated type
    // Einride Gothenburg
	origin := v8types.Waypoint{
		Place: v8types.LatLong{
			Lat:  57.707752,
			Long: 11.949767,
		},
    }
    // Einride Stockholm
	destination := v8types.Waypoint{
		Place: v8types.LatLong{
			Lat:  59.337492,
			Long: 18.063672,
		},
	}
    // Let's request the Travel Summary back in the 
    returnVals := optional.NewInterface([]routingv8.Return{routingv8.RETURN_TRAVEL_SUMMARY})
    
    // Calculate route by Truck
    resp, _, err := client.RoutingApi.CalculateRoutes(auth, routingv8.ROUTERMODE_TRUCK, origin.String(), destination.String(), &routingv8.CalculateRoutesOpts{
        Return_: returnVals,
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("Distance: %d meters\n", resp.Routes[0].Sections[0].TravelSummary.Length)
}
```

### Authentication

Authentication is passed through a couple of helper types which inject the
authentication token in the context. See example on how to use API keys.
If you have an OAuth2 access token (for example, a [personal
API token][]), you can use it with the oauth2 library using:

```go
import "golang.org/x/oauth2"

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	auth := context.WithValue(context.Background(), routingv8.ContextOAuth2, ts)
}
```