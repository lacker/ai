package hex

import (
)

// A compact representation of a small set of features.
// This should be useful for things like indexing vectors over the
// space of features.
type QFeatureSet uint16

// A QFeatureSet only handles up to 2 features.
// 0: empty feature set
// 1-242: single features (1 + qfeature)
// some weird triangular thing: double features
const EmptyFeatureSet QFeatureSet = 0
const MinSingleFeature QFeatureSet = EmptyFeatureSet + 1
const MaxSingleFeature QFeatureSet = MinSingleFeature +
	QFeatureSet(MaxQFeature)

// TODO: handle double features
