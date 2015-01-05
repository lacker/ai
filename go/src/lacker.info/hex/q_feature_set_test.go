package hex

import (
	"testing"
)

func TestQFeatureSetConversion(t *testing.T) {
	var covered [NumFeatureSets]bool

	for f1 := MinQFeature; f1 <= MaxQFeature; f1++ {
		for f2 := f1 + 1; f2 <= MaxQFeature; f2++ {
			fs := MakeDoubleton(f1, f2)
			fs2 := MakeDoubleton(f2, f1)
			if fs != fs2 {
				t.Fatal("fs != fs2")
			}
			covered[fs] = true
			decoded1, decoded2 := fs.Features()
			if decoded1 > decoded2 {
				decoded1, decoded2 = decoded2, decoded1
			}
			if f1 != decoded1 {
				t.Fatal("f1 != decoded1")
			}
			if f2 != decoded2 {
				t.Fatal("f2 != decoded2")
			}
		}
	}

	for fs := MinDoubleton; fs <= MaxDoubleton; fs++ {
		if !covered[fs] {
			t.Fatalf("fs is not covered")
		}
	}
}
