package unpack

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	unpackedReg, err := Unpack("ж4b4")
	unpackedNotRight, notRightErr := Unpack("/ж4b4")
	if err != nil {
		t.Errorf("error occured %s, want: no err.", err)
	}
	if unpackedReg != "жжжжbbbb" {
		t.Errorf("string was unpacked incorrect, got: %s, want: %s.", unpackedReg, "жжжжbbbb")
	}
	if notRightErr == nil {
		t.Errorf("error on incorrect string must be")
	}
}
