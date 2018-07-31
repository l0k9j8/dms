package dlna

import (
	"testing"
)

func TestContentFeaturesString(t *testing.T) {
	a := ContentFeatures{
		SupportTimeSeek: true,
	}.String()
	e := "DLNA.ORG_OP=10;DLNA.ORG_CI=1"
	if e != a {
		t.Fatal(a)
	}
}
