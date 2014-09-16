package mime

import (
	"testing"
)

func TestExtraExtensions(t *testing.T) {
	if !IsKnownExtension("jpg") {
		t.Fatalf("jpg should be a known extension")
	}
	if !IsKnownExtension("jpeg") {
		t.Fatalf("jpeg should be a known extension")
	}

	if "jpg" != TypeJPEG.DefaultExtension() {
		t.Fatalf("jpg should be the default extension")
	}
}
