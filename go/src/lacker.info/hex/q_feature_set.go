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
// 243-29403: double features
//
// The packing logic is basically, with features f1..f242:
// 243: (f1, f2)
// 244, 245: (f1, f3), (f2, f3)
// 246, 247, 248: (f1, f4), (f2, f4), (f3, f4)
// etc.
// There are (242)(242-1)/2 = 29161 double features.

const EmptyFeatureSet QFeatureSet = 0
const MinSingleton QFeatureSet = EmptyFeatureSet + 1
const NumSingletons = QFeatureSet(MaxQFeature)
const MaxSingleton QFeatureSet = MinSingleton + NumSingletons - 1
const MinDoubleton QFeatureSet = MaxSingleton + 1
const NumDoubletons QFeatureSet = NumSingletons * (NumSingletons - 1) / 2
const MaxDoubleton QFeatureSet = MinDoubleton + NumDoubletons - 1

// TODO: handle double features
