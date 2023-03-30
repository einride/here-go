package routingv8

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestUnmarshalRoute(t *testing.T) {
	t.Parallel()
	t.Run("route.json", func(t *testing.T) {
		t.Parallel()
		resp := unmarshalRouteResponseFromFile(t, "route.json")
		assert.DeepEqual(t, resp, RoutesResponse{
			Routes: []Route{
				{
					ID: "bfaed7d0-19c7-4e72-81b7-24eeb148b62b",
					Sections: []Section{
						{
							ID:   "85357f8f-00ad-447e-a510-d8c02e0b264f",
							Type: "vehicle",
							Arrival: VehicleDeparture{
								Place: Place{
									Type: "place",
									Location: GeoWaypoint{
										Lat:  52.53232637420297,
										Long: 13.378873988986015,
									},
									OriginalLocation: GeoWaypoint{},
								},
							},
							Departure: VehicleDeparture{
								Place: Place{
									Type: "place",
									Location: GeoWaypoint{
										Lat:  52.53098367713392,
										Long: 13.384566977620125,
									},
									OriginalLocation: GeoWaypoint{},
								},
							},
							Summary: Summary{
								Duration:     123,
								Length:       538,
								BaseDuration: 0,
							},
							Polyline: "",
						},
					},
				},
			},
		})
	})
	t.Run("route-polyline.json", func(t *testing.T) {
		t.Parallel()
		resp := unmarshalRouteResponseFromFile(t, "route-polyline.json")
		assert.DeepEqual(t, resp, RoutesResponse{
			Routes: []Route{
				{
					ID: "81e526c0-5693-4bc0-bdbb-239ecc2857e7",
					Sections: []Section{
						{
							ID:   "3b64cfbf-4a78-487b-ab37-0676b06d2456",
							Type: "vehicle",
							Arrival: VehicleDeparture{
								Place: Place{
									Type: "place",
									Location: GeoWaypoint{
										Lat:  52.53232637420297,
										Long: 13.378873988986015,
									},
									OriginalLocation: GeoWaypoint{},
								},
							},
							Departure: VehicleDeparture{
								Place: Place{
									Type: "place",
									Location: GeoWaypoint{
										Lat:  52.53098367713392,
										Long: 13.384566977620125,
									},
									OriginalLocation: GeoWaypoint{},
								},
							},
							Summary: Summary{},
							Polyline: "BGwynmkDu39wZvBtFAA3InfAAvHrdAAvHvbAAoGzF0FnGoGvHsOvRAA8L3NAAkSnVAAo" +
								"GjIsEzFAAgFvHkDrJAAwHrJoVvb0ezoBAAjInVAA3N_iBAAzJ_Z",
						},
					},
				},
			},
		})
	})
	t.Run("route-restrictions.json", func(t *testing.T) {
		t.Parallel()
		resp := unmarshalRouteResponseFromFile(t, "route-restrictions.json")
		assert.DeepEqual(t, resp, RoutesResponse{
			Routes: []Route{
				{
					ID: "7a63b7b6-7b62-4d80-a7cf-3edd1eafa2f9",
					Sections: []Section{
						{
							ID:   "7cb3dcba-6d43-41e0-95eb-3a3f3af77340",
							Type: "vehicle",
							Arrival: VehicleDeparture{
								Place: Place{
									Type: "place",
									Location: GeoWaypoint{
										Lat:  51.1086699,
										Long: 17.0387979,
									},
									OriginalLocation: GeoWaypoint{
										Lat:  51.1086709,
										Long: 17.0388039,
									},
								},
							},
							Departure: VehicleDeparture{
								Place: Place{
									Type: "place",
									Location: GeoWaypoint{
										Lat:  51.0193731,
										Long: 17.1613281,
									},
									OriginalLocation: GeoWaypoint{
										Lat:  51.019519,
										Long: 17.1615459,
									},
								},
							},
							Summary: Summary{
								Duration:     1624,
								Length:       13451,
								BaseDuration: 1490,
							},
							Polyline: "BG6m_phDgnu3gBif3zBgKvR0K3S4NjXwHvM0FrJoGzK8G7LsJ_O0KvR8Q3c0FrJ0jBj6Bs" +
								"YnpBwR3c8a_sBsJzPoVzjBgUvgBgFjIsE7GoGzKwHjNgKjS0F_J8G7LsJnQ8Q3coG_JoQ7a4N7Vw" +
								"MzUwMnVkN7Vs7B_jDoQ7akI3NkIrO0KjSwHvMsEvHgFjIwHjNkN7VoG_JsEvHoGzK8GzKoG_J0F3" +
								"I0F3IsEnGsEnGsEnGsEzFoGjIwMnQsEnGgF7GoGjIkIzKoLrO0FjI0FjI4I3NoG_J8GnLoLvRoG_" +
								"J0F3IoGrJsJvMwRvWozB79B0PrTsJnL0tBn4BsOjSkNzP4SvWoVnawWvb8kB_sBkNzP8VnasE_Eg" +
								"PjSkI_J4IzK8G3I8GrJ8G_JgFvHgFjIgFjI4IrO0KvRkI3N0KvRkNvWsY_nBwM7VwHjI4DrEgFnG" +
								"kIzKwH_J4D_E4D3DkDvCkDAkDAwCnB8BnBoBnB8B7BwC3DoBvCoB3DUjDU_EAjDTzF8BnG8B_EwC" +
								"zFwHrOgFzKsErJsEnL4DzK4DnGoG_J4DnGwWvlB0FrJkN7V8L_TgFjI8G7LsJzPsEvHsEvHoajrB" +
								"08B7lDsEvHsJzPsEvHoiC7vDwHvM4X_nBsO_YoQnaoLrTsEvH4IrOkcvvBwR3c4IrOsTjhBoGzKw" +
								"WjmBkIrO4IrOgK7Q0FrJ0FrJoGzKwHvM4DnGoV3hBoGrJ0FjIoG3I0FvH8G3IsY_d8G3IwHrJwHr" +
								"J8QnV0KjNoGvH8G3I0F7GsJnLsOvRgKvM4IzKoGvHoGvHwHrJ8LrO8GjIoGvHwH3IoGvH4DrE4Dr" +
								"E0FnG0FnGoGnGgFzFsJzK4IrJgKzKgKzKoL7LwHvH4IrJwHjI0KnLsJ_JsJrJoL7LgK_J8GvHgF_" +
								"EgF_E8GvHkI3IgK_JkN3N4IrJwMvM4IrJsE_EgKzK8LvMwMjN4S3S4IrJ8GvH4IrJsJrJ8LvMkSr" +
								"TgK_JoGnG8GvH0FzFwHjI0UnVgenfsJ_JoG7G0FzFgF_E0FnG0KnL4IrJ8G7GwHvHgZna8G7GoG7" +
								"GoG7GwHjI4I3IosC3uCgevgBgUzU8LvMsnBnpB0P7QoQ7QgP_O4uC3zCssBztBofjhBsYzZ8L7Lo" +
								"L7L0KnLoL7LsJ_J4IrJwHvHkhBriBkS3S8VjX0KnL4IrJ8G7G4I3IkI3I4IrJ8QvRwgBriB4I3Iw" +
								"W3X8VjXgK_JoL7LwMjNsJ_J8LvM8LvMwMvMsOrO4NvMkN7L0P3N4XnV0FzFgF_EsE_EgFnGgF7Gw" +
								"jCz6CoV3ckInLsJvM0Uvb8ankBsJ7L4I_J4IzKwH_J4I7LsJjN4N3S8QvW4I7LwHnL4IjNgKzPge" +
								"rxBoLjSgK7QwHjNsEjIoGzK8LrT8V7kBsEvHgF3I4I_OgK7QgKnQ8Q7akN7VgPzZgPrY4DnGsEvH" +
								"gFjIsEvHsEvH4DvHgF7L4InVsE_JgFrJwHjNoG_JwHzK8LnQ8GrJkDjDwCvCwCjDwCjDgK_OsEnG" +
								"0ZnkB0KzPoQrYoajmBwWzewWnfsE7GkDrEgUrd8V7fwRrdgK7Q0KzPgKrOwHzK4I7LoGjI8GrJ4I" +
								"nLwH_JwHzK8LjS4X_iBkInLoG3I8G_JoQjXgFjI4D7G4DvHsE_JgF7GsEzFsJvMoG3IgPnVoLnQ4" +
								"DzFsEnGkN3SkNjS0Z7kB4IjNwHnL4DzFsOzUsJjNsEnGgFvH8GrJ4I7L4IvMsEnGoG3IgU3cgF7G" +
								"kI7LkI7LwHzKgF7GgFnG4DrEsE_EgF_E0F_EoG_E8G_EoLjI4NzKwM_JkIvH4DjD4DvC4D7B0F7B" +
								"sJ7B8LvCoQjDsJ7B0P3DkI7B8G7BwHvC0FvC0FjD8GrEoG_EoGzFoG7G8GjI8VzZkI3IwHvHgF_E" +
								"0FzFkIjI8V7V0FzF4D3DwMvMkIjI0FzF8anawbjcsOzPsE_EoL3NoL3NoL_OkInL0U3coG3IwHzK" +
								"oG3IwR_YkInLkIzK8L_OkNzP4IzKgK7LwHjI4DrEkIT0FAsEAgFAoGUsEU0FoBoG8B4D8BsEwCwH" +
								"gF4IgFsE8B4InpBoLnkBkDzFrEzFrEzFzK3NokBjzCwHvRgK3XkI_T8G3IoGjN4I3XsE_OkDjSwH" +
								"jrBkIr2BwCrOsEzjBoBvMgF_O8BvH8B3I4D3X4Dna4DjhBkD3I8B_EwCrE4D3DkDjDsJA8GnB8GvCkInD",
						},
					},
				},
			},
		})
	})
}

func unmarshalRouteResponseFromFile(t *testing.T, filename string) RoutesResponse {
	bs, err := os.ReadFile(path.Join("testdata", filename))
	assert.NilError(t, err)
	var resp RoutesResponse
	assert.NilError(t, json.Unmarshal(bs, &resp))
	return resp
}
