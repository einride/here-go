package routingv7

import (
	"time"
)

type Duration float64

func (d Duration) AsDuration() time.Duration {
	return time.Duration(d) * time.Second
}
