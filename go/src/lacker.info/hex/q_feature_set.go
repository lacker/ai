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
const NumFeatureSets QFeatureSet = MaxDoubleton + 1
const NotAFeatureSet QFeatureSet = NumFeatureSets

func (fs QFeatureSet) IsEmpty() bool {
	return fs == EmptyFeatureSet
}

func (fs QFeatureSet) IsSingleton() bool {
	return fs >= MinSingleton && fs <= MaxSingleton
}

func (fs QFeatureSet) IsDoubleton() bool {
	return fs >= MinDoubleton && fs <= MaxDoubleton
}

func (fs QFeatureSet) SingletonFeature() QFeature {
	return QFeature(fs - MinSingleton)
}

func MakeSingleton(f QFeature) QFeatureSet {
	return QFeatureSet(f) + MinSingleton
}

// Returns NotAFeature once we run out
func (fs QFeatureSet) Features() (QFeature, QFeature) {
	if fs.IsEmpty() {
		return NotAFeature, NotAFeature
	}
	if fs.IsSingleton() {
		return fs.SingletonFeature(), NotAFeature
	}
	panic("TODO")
}
