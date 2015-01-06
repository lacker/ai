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
// There are (242)(242-1)/2 = 29161 double features.
// To see how they are packed into the space, it's best just to check
// the code for MakeDoubleton and Features.

const EmptyFeatureSet QFeatureSet = 0
const MinSingleton QFeatureSet = EmptyFeatureSet + 1
const NumSingletons = QFeatureSet(NumFeatures)
const MaxSingleton QFeatureSet = MinSingleton + NumSingletons - 1
const MinDoubleton QFeatureSet = MaxSingleton + 1
const NumDoubletons QFeatureSet = NumSingletons * (NumSingletons - 1) / 2
const MaxDoubleton QFeatureSet = MinDoubleton + NumDoubletons - 1
const NumFeatureSets QFeatureSet = MaxDoubleton + 1
const NotAFeatureSet QFeatureSet = NumFeatureSets

// Map pairs of feature to QFeatureSet and vice versa.
var FeaturePairToFeatureSet [NumFeatures][NumFeatures]QFeatureSet
var FeatureSetToFirstFeature [NumFeatureSets]QFeature
var FeatureSetToSecondFeature [NumFeatureSets]QFeature

func init() {
	for fs := EmptyFeatureSet; fs <= MaxSingleton; fs++ {
		FeatureSetToFirstFeature[fs] = NotAFeature
		FeatureSetToSecondFeature[fs] = NotAFeature
	}
	fs := MinDoubleton
	for f1 := MinFeature; f1 <= MaxFeature; f1++ {
		FeaturePairToFeatureSet[f1][f1] = NotAFeatureSet
		for f2 := f1 + 1; f2 <= MaxFeature; f2++ {
			FeaturePairToFeatureSet[f1][f2] = fs
			FeaturePairToFeatureSet[f2][f1] = fs
			FeatureSetToFirstFeature[fs] = f1
			FeatureSetToSecondFeature[fs] = f2
			fs++
		}
	}
	if fs != MaxDoubleton + 1 {
		panic("bad feature set init code")
	}
}

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

func MakeDoubleton(f1 QFeature, f2 QFeature) QFeatureSet {
	return FeaturePairToFeatureSet[f1][f2]
}

// Returns NotAFeature once we run out
func (fs QFeatureSet) Features() (QFeature, QFeature) {
	if fs.IsEmpty() {
		return NotAFeature, NotAFeature
	}
	if fs.IsSingleton() {
		return fs.SingletonFeature(), NotAFeature
	}
	return FeatureSetToFirstFeature[fs], FeatureSetToSecondFeature[fs]
}
