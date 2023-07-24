package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestDCMITypeConstructor(t *testing.T) {
	v := terms.NewDCMIType()
	if v == nil {
		t.Errorf("terms.NewDCMIType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.DCMIType should validate: %s", err)
	}
}

func TestDCMITypeMarshalUnmarshal(t *testing.T) {
	v := terms.NewDCMIType()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewDCMIType()
	xml.Unmarshal(buf, v2)
}
