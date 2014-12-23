package hex

import (

)

type Listener interface {
	HandleNotification(spot TopoSpot)
}
