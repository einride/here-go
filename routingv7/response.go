package routingv7

type RouteResponseMetaInfo struct {
	// Mirrored RequestId value from the request structure. Used to trace requests.
	RequestID string `json:"requestId,omitempty"`
	// Time at which the search was performed.
	Timestamp string `json:"timestamp,omitempty"`
	// Gives the version of the underlying map, upon which the route calculations are based.
	MapVersion string `json:"mapVersion,omitempty"`
	// Gives a list of map versions, upon which the route calculations are based.
	AvailableMapVersion []string `json:"availableMapVersion,omitempty"`
	// Gives the version of the module that performed the route calculations.
	ModuleVersion string `json:"moduleVersion,omitempty"`
	// Required. Gives the version of the schema definition to enable
	InterfaceVersion string `json:"interfaceVersion,omitempty"`
}
