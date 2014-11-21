package hex

import (

)

// A SpotRegistry is designed to hold a bunch of listeners, each
// affiliated to one spot, and notify
// them when the contents of that spot change.
type SpotRegistry struct {
	// For now, all listeners are delta neurons, but this might be
	// useful to extend.
	// The first key is the spot. The second is the list of listeners.
	listeners [NumTopoSpots][]*DeltaNeuron
}

// Create a new spot registry with no listeners
func NewSpotRegistry() {
}

func (sr *SpotRegistry) Listen(spot TopoSpot, dn *DeltaNeuron) {
	if sr.listeners[spot] == nil {
		sr.listeners[spot] = make([]*DeltaNeuron, 1)
		sr.listeners[spot][0] = dn
	} else {
		sr.listeners[spot] = append(sr.listeners[spot], dn)
	}
}

func (sr *SpotRegistry) Notify(bf BasicFeature) {
	panic("TODO")
}
