# here-go

here-go is a Go client library for accessing the HERE APIs

API docs can be found [here](https://developer.here.com/documentation/)

## Usage

```go
import "github.com/einride/here-go/pkg/here"
```

Construct a new HERE client, then use the various services on the client to
access different parts of the HERE API. For example:

```go
httpClient := here.NewAPIKeyHTTPClient(".... your api key ...")
client := here.NewClient(httpClient)

// Get routes from Einride Gothenburg to Einride Stockholm
json, err := client.Routing.Get(context.Background(), here.RoutingRequest{
    TransportMode: here.TransportModeCar,
    Origin: here.Waypoint{
        Place: here.LatLong{
            Lat:  57.707746,
            Long: 11.9475834,
        },
    },
    Destination: here.Waypoint{
        Place: here.LatLong{
            Lat:  58.5177323,
            Long: 13.8859,
        },
    },
})
```

### Authentication

The here-go library does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you. The easiest and recommended way to do this is using the [oauth2][]
library, but you can always use any other library that provides an
`http.Client`. If you have an OAuth2 access token (for example, a [personal
API token][]), you can use it with the oauth2 library using:

```go
import "golang.org/x/oauth2"

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := here.NewClient(tc)
}
```

Note that when using an authenticated Client, all calls made by the client will
include the specified OAuth token. Therefore, authenticated clients should
almost never be shared between different users.

See the [oauth2 docs][] for complete instructions on using that library.

If you wish to use an API Key, use the APIKey client

```go
func main() {
	httpClient := here.NewAPIKeyHTTPClient(".... your api key ...")
	client := here.New(httpClient)
}
```
