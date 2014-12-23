package hex

import (

)

// A SpotRegistry is designed to hold a bunch of listeners, each
// affiliated to one spot, and notify
// them when the contents of that spot change.
type SpotRegistry struct {
	// The first key is the spot. The second is the list of listeners.
	listeners [NumTopoSpots][]Listener
}

// Create a new spot registry with no listeners
func NewSpotRegistry() *SpotRegistry {
	return &SpotRegistry{}
}

// Add a new listener for a spot.
func (sr *SpotRegistry) Listen(spot TopoSpot, x Listener) {
	if sr.listeners[spot] == nil {
		sr.listeners[spot] = make([]Listener, 1)
		sr.listeners[spot][0] = x
	} else {
		sr.listeners[spot] = append(sr.listeners[spot], x)
	}
}

// This doesn't clear the listeners list. It could, though.
func (sr *SpotRegistry) Notify(spot TopoSpot) {
	for _, x := range sr.listeners[spot] {
		x.HandleNotification(spot)
	}
}
